package jd_union_go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type JDUnion struct {
	ID string
}

type App struct {
	ID     string
	Name   string
	Key    string
	Secret string

	debug bool
}

type JdUnionErrResp struct {
	ErrorResponse ErrorResponse `json:"error_response"`
}

type ErrorResponse struct {
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	ZhDesc    string `json:"zh_desc"`
	EnDesc    string `json:"en_desc"`
	RequestId string `json:"request_id"`
}

const RouterURL = "https://api.jd.com/routerjson"
const RequestMethod = "POST"

func (app *App) Debug() {
	app.debug = true
}

func (app *App) Request(method string, paramJSON map[string]interface{}) ([]byte, error) {
	// common params
	params := map[string]interface{}{}
	params["method"] = method
	params["app_key"] = app.Key
	params["format"] = "json"
	params["v"] = "1.0"
	params["sign_method"] = "md5"
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")

	// api params
	paramJSONStr, _ := json.Marshal(paramJSON)
	params["360buy_param_json"] = string(paramJSONStr)
	params["sign"] = GetSign(app.Secret, params)
	if app.debug {
		log.Printf("Request: %s, %s, %v", RouterURL, method, params)
	}
	resp, err := http.PostForm(RouterURL, app.Values(params))
	if app.debug {
		log.Printf("Responce:%v %v", resp, err)
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if app.debug {
		log.Printf("Responce Body:%v ", string(body))
	}
	if err != nil {
		return nil, err
	}

	jdErr := JdUnionErrResp{}
	if err := json.Unmarshal(body, &jdErr); err != nil {
		return nil, err
	}

	if jdErr.ErrorResponse.Code != "" {
		return nil, fmt.Errorf("%s:%s:%s:%s", jdErr.ErrorResponse.Code, jdErr.ErrorResponse.Msg, jdErr.ErrorResponse.ZhDesc, jdErr.ErrorResponse.RequestId)
	}
	return body, nil
}

func (app *App) Values(params map[string]interface{}) url.Values {
	vals := url.Values{}
	for key, val := range params {
		vals.Add(key, GetString(val))
	}
	return vals
}
