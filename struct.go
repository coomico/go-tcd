package tcd

type Raw struct {
	Data []Data
}

type Data struct {
	ID int `json:"id"`
	Noput string `json:"noput"`
	Absris string `json:"absris"`
	Tahun string `json:"tahun"`
	Author string `json:"author"`
	Dateris string `json:"dateris"`
	DatePublish string `json:"datepublish"`
	Flris string `json:"flris"`
	Tampil bool `json:"tampil"`
	IsPutusan bool `json:"isPutusan"`
}

type Q struct {
	Sort string
	Page string
	PageSize string
	Group string
	Filter string
}