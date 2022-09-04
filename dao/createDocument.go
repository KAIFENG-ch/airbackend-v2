package dao

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"test/serialize"
)

const (
	url = "http://localhost:5984/"
)

type EnterpriseInfo struct {
	_id      string
	db       string
	staffer  Staffer
	workshop WorkShop
}

type Staffer struct {
	role     string
	username string
}

type WorkShop struct {
}

func Head(enterpriseName string, docId string) bool {
	headUrl := url + enterpriseName + "/" + docId
	resp, err := http.Head(headUrl)
	if err != nil {
		return false
	}
	if resp.Status != "200 OK" {
		return false
	}
	return true
}

func GetDoc(enterpriseName string, docId string) map[string]interface{} {
	target := url + enterpriseName + "/" + docId
	resp, err := http.Get(target)
	if err != nil {
		return nil
	}
	data := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&data)
	fmt.Println(data)
	if err != nil {
		return nil
	}
	return data
}

func CreateDoc(enterpriseName string, dbName string, username string, password string) bool {
	target := fmt.Sprintf("http://%s:%s@127.0.0.1:5984/enterprises/%s", username, password, dbName)
	client := &EnterpriseInfo{
		_id: enterpriseName,
		db:  dbName,
	}
	msg, err := json.Marshal(&client)
	if err != nil {
		log.Println(err)
		return false
	}
	body := bytes.NewBuffer(msg)
	req, err := http.NewRequest("PUT", target, body)
	if err != nil {
		log.Println(err)
		return false
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}
	if resp.Status != "201 Created" {
		return false
	}
	return true
}

func CreateDb(name string, dbName string, username string, password string) serialize.Response {
	url := fmt.Sprintf("http://%s:%s@127.0.0.1:5984/%s", username, password, name)
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application-json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return serialize.Response{
			Err: err,
		}
	}
	if resp.StatusCode != http.StatusCreated {
		return serialize.Response{Err: errors.New("database create failed")}
	}
	exist := Head(name, dbName)
	if exist {
		return serialize.Response{Err: err}
	}
	ok := CreateDoc(name, dbName, username, password)
	if !ok {
		return serialize.Response{}
	}
	return serialize.Response{
		Status: 200,
		Msg:    "OK",
		Err:    nil,
	}
}
