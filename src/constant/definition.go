package constant

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	Cid string
	Pid int
	Cname string
	Pname string
	Price float32
	Energy int
	Img string
	Flag string
	Tag string
	Note string
}

type ResJson struct {
	Header struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	} `json:"header"`
	Body [] interface{} `json:"body"`
}