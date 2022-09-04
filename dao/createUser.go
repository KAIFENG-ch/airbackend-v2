package dao

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"test/serialize"
)

const (
	userId = "org.couchdb.user"
	TYPE   = "user"
	ROLE   = "boss"
)

type UserInfo struct {
	Id       string `json:"_id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Roles    string `json:"roles"`
	Password string `json:"password"`
}

func CreateUser(enterprise string, docId string, username string, password string) serialize.Response {
	target := url + "_users/" + userId + ":" + username
	newUser := &UserInfo{
		Name:     username,
		Type:     TYPE,
		Roles:    ROLE,
		Password: password, // todo 密码
	}
	info, err := json.Marshal(newUser)
	body := ioutil.NopCloser(bytes.NewBuffer(info))
	req, err := http.NewRequest("PUT", target, body)
	req.Header.Add("Content-Type", "application-json")
	if err != nil {
		log.Println(err)
		return serialize.Response{
			Err: err,
		}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return serialize.Response{
			Err: err,
		}
	}
	if resp.Status != "200 OK" {
		return serialize.Response{
			Err: errors.New("no response"),
		}
	}
	return serialize.Response{
		Status: 200,
		Msg:    "ok",
		Err:    nil,
	}
}
