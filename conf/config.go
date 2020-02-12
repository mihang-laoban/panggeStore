package conf

type ConfigStruct struct {
	Debug string
	Base struct {
		Env string
		Db map[string]interface{}
		Server map[string]interface{}
	}
}