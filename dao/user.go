package dao

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetDb() {
	url := "http://admin:123456@127.0.0.1:5984/_all_dbs"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func DeleteDb(name string) {
	url := fmt.Sprintf("http://admin:123456@127.0.0.1:5984/%s", name)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
