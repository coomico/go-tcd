package tcd

import (
	"os"
	"io"
	"fmt"
	"log"
	"sync"
	"strconv"
	"net/http"
	"encoding/json"
	"path/filepath"
)

var (
	Path = _path_files()
	UserAgent = "put_yours_here"
)

const (
	url_data = "https://setpp.kemenkeu.go.id/risalah/Putusan_Read?"
	url_file = "https://setpp.kemenkeu.go.id/risalah/ambilFileDariDisk/"
)

func New() *Q {
	q := new(Q)
	q.Sort = IDDesc
	q.Page = "1"
	q.PageSize = "5"
	q.Group = ""
	q.Filter = ""

	return q
}

func (q *Q) FetchData() *Raw {
	form := "?sort=%s&page=%s&pageSize=%s&group=%s&filter=%s"
	formFormated := fmt.Sprintf(
		form,
		q.Sort,
		q.Page,
		q.PageSize,
		q.Group,
		q.Filter,
	)
	url := url_data+formFormated

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "39")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("User-Agent", UserAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		log.Fatal("not ok!")
	}

	defer resp.Body.Close()

	var result Raw
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		panic(err)
	}
	return &result
}

func GetFile(id int) {
	url := url_file+strconv.Itoa(id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", UserAgent)
	req.Close = true

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal(resp.StatusCode)
	}

	log.Printf("getting file %d...", id)
	filename := Path+strconv.Itoa(id)+".pdf"
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}
	
	return
}

//error EOF if pagesize >5
func (r *Raw) GetFileBulk() {
	var wg sync.WaitGroup

	done := make(chan int, len(r.Data))

	for i := 0; i < len(r.Data); i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			GetFile(id)

			done <- id
		}(r.Data[i].ID)
	}

	for i := 0; i < len(r.Data); i++ {
		log.Printf("done %d!", <-done)
	}

	wg.Wait()
	return
}

func _path_files() string {
	path, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}

	path = path + `\files\`

	_, err = os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		if err = os.Mkdir("files", 0755); err != nil {
			panic(err)
		}
	}

	return path
}