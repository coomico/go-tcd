package main

import (
	"github.com/coomico/go-tcd"
)

func main() {
	raw := tcd.New().FetchData()
	raw.GetFileBulk()
}