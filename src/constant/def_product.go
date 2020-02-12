package constant

type ItemJson struct {
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

type ProductJson struct {
	classification string
	tid int
	items ItemJson
}

type ResJson struct {
	Header struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	} `json:"header"`
	Body [] interface{} `json:"body"`
}

type Ids struct {
	Cid int
	Pid int
}