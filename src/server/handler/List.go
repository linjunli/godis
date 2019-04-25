package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	LPUSH = "LPUSH"
	LPOP  = "LPOP"
)

type ListReq struct {
	Operate string `json:"operate"`
	Key string `json:"key"`
	Values []string `json:"values"`
}

type ListRsp struct {
	Ret int `json:"ret"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

type List struct {
	Data map[string]interface{}
	w http.ResponseWriter
	r *http.Request
}

func (l *List)ServeHTTP(w http.ResponseWriter,r *http.Request) {
	l.w = w
	l.r = r

	if err := r.ParseForm(); err != nil {
		log.Printf("request parse form err:%s",err.Error())
		l.listResult(-1000, err.Error(),nil)
		return
	}

	if strings.ToUpper(r.Method) != "POST" {
		l.listResult(-2000,"request method must be post",nil)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		l.listResult(-3000,"read request body err:"+err.Error(),nil)
		return
	}
	
	req := &ListReq{}
	
	err = json.Unmarshal(b,req)
	if err != nil {
		l.listResult(-4000,"request json parse err:"+err.Error(),nil)
		return
	}

	if req.Key == "" {
		l.listResult(-5000,"request miss key param",nil)
		return
	}

	switch req.Key {
	case LPOP:
		break
	case LPUSH:
		break
	default:
		l.listResult(-7000,"operation not match",req.Operate)
		break
	}


}

func (l *List)listResult(ret int, msg string, data interface{})  {
	rsp := &StringRsp{
		Ret:ret,
		Msg:msg,
		Data:data,
	}

	b,_ := json.Marshal(rsp)

	_,_ = l.w.Write(b)
}
