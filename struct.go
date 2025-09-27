package tcd

type Raw struct {
	Total int `json:"Total"`
	Data  []Data
}

type Data struct {
	ID          int    `json:"id"`
	Noput       string `json:"noput"`
	Absris      string `json:"absris"`
	Tahun       string `json:"tahun"`
	Author      string `json:"author"`
	Dateris     string `json:"dateris"`
	DatePublish string `json:"datepublish"`
	Flris       string `json:"flris"`
	Tampil      bool   `json:"tampil"`
	IsPutusan   bool   `json:"isPutusan"`
}

type Query struct {
	Sort     sort
	Page     int
	PageSize int
	Group    string
	filter   string
}
