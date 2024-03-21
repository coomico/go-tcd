package main

import (
	"github.com/coomico/go-tcd"
)

func main() {
	tcd.UserAgent = "your user-agent. WHERE?? put here!"
	q := tcd.New()
	q.PageSize = "3"

	raw := q.FetchData()
	raw.GetFileBulk()
}