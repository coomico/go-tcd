package main

import (
	"slices"

	"github.com/coomico/go-tcd"
)

func main() {
	q := tcd.New()
	q.Page = 1
	q.PageSize = 30
	q.Sort = tcd.AbsrisAsc
	q.FilterByAbsris("bea")

	raw, err := q.FetchData()
	if err != nil {
		panic(err)
	}

	// removing data that does not have a final decision file
	raw.Data = slices.DeleteFunc(raw.Data, func(d tcd.Data) bool {
		return d.Flris == ""
	})

	err = raw.GetFileBulk()
	if err != nil {
		panic(err)
	}
}
