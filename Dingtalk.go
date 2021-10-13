package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const webhook_url = "https://oapi.dingtalk.com/robot/send?access_token=27d648673d9a92c5ca19de7d92d432011ef34dd48670062bd4051805ae5cc457"

func DingToInfo(s string) string {
	content, data := make(map[string]string), make(map[string]interface{})
	content["content"] = s
	data["msgtype"] = "text"
	data["text"] = content

	b, _ := json.Marshal(data)
	resp, err := http.Post(webhook_url,
		"application/json",
		bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	return string(body)
}
