package day

import (
	"encoding/json"
	"errors"
	"net/http"
)

const url = "http://v.juhe.cn/toutiao/index?type=top&key=8a26f55d1e10c193259c4606d27a00ac"

type news struct {
	title string `json:"title"`

}
//News get news everyday
func News() ([]string,error) {
	n := []string{}
	res,err := http.Get(url)
	if (err != nil) {
		return n, errors.New("请求失败")
	}
	news := res.Body
	json.Unmarshal()
	return []string{}
}
