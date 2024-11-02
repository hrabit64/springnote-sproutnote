package data

var data = map[string]string{}

func SetData(key string, value string) {
	data[key] = value
}

func GetData(key string) string {
	return data[key]
}

func ClearData() {
	data = map[string]string{}
}
