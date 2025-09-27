package tcd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var (
	DirName   = "files"
	UserAgent = "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36"
)

const (
	baseUrl    = "https://setpp.kemenkeu.go.id"
	risalahUrl = baseUrl + "/risalah"
)

func New() *Query {
	return &Query{
		Sort:     IDDesc,
		Page:     1,
		PageSize: 5,
	}
}

func (q *Query) FetchData() (*Raw, error) {
	form := url.Values{}
	form.Add("sort", string(q.Sort))
	form.Add("page", strconv.Itoa(q.Page))
	form.Add("pageSize", strconv.Itoa(q.PageSize))
	form.Add("group", q.Group)
	form.Add("filter", q.filter)

	formEncoded := form.Encode()
	payload := bytes.NewBufferString(formEncoded)

	req, err := http.NewRequest("POST", risalahUrl+"/Putusan_Read", payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Origin", baseUrl)
	req.Header.Set("Referer", risalahUrl)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("User-Agent", UserAgent)
	req.Close = true

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status response %s", resp.Status)
	}

	var result Raw
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if result.Total == 0 {
		return nil, errors.New("empty data results")
	}

	return &result, nil
}

func GetFile(id int) error {
	var attempt int
	var filename string

	dirPath := pathFiles()

retry:
	if attempt > 3 {
		return fmt.Errorf("failed getting file id=%d filename=%s", id, filename)
	}

	url := risalahUrl + "/ambilFileDariDisk/" + strconv.Itoa(id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Origin", baseUrl)
	req.Header.Set("Referer", risalahUrl)
	req.Header.Set("User-Agent", UserAgent)
	req.Close = true

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("status response %s", resp.Status)
	}

	disp := resp.Header.Get("Content-Disposition")
	filename = strings.Split(disp, "=")[1]
	filepath := dirPath + filename
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	switch err {
	case nil:
		log.Printf("done file id=%d filename=%s", id, filename)
		return nil
	case io.ErrUnexpectedEOF:
		attempt++
		goto retry
	default:
		return err
	}
}

func (r *Raw) GetFileBulk() error {
	amount := len(r.Data)
	if amount == 0 {
		return errors.New("empty data list")
	}

	pathFiles()

	errCh := make(chan error, 1)
	var wg sync.WaitGroup

	for i := 0; i < amount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			errCh <- GetFile(id)
		}(r.Data[i].ID)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	errs := []error{}
	for err := range errCh {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func pathFiles() string {
	path, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}

	path += "/" + DirName + "/"

	_, err = os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		if err = os.Mkdir(DirName, 0o755); err != nil {
			panic(err)
		}
	}

	return path
}
