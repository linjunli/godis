package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	GET = "GET"
	SET = "SET"
)

type String struct {
	Data map[string]interface{}
	w http.ResponseWriter
	r *http.Request
}

type StringReq struct {
	Operate string `json:"operate"`
	Key string `json:"key"`
	Value string `json:"value"`
}

type StringRsp struct {
	Ret int `json:"ret"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func (s *String)ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	s.w = w
	s.r = r

	if err := r.ParseForm(); err != nil {
		log.Printf("request parse form err:%s",err.Error())
		s.stringResult(-1000, err.Error(),nil)
		return
	}

	if strings.ToUpper(r.Method) != "POST" {
		s.stringResult(-2000,"request method must be post",nil)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		s.stringResult(-3000,"read request body err:"+err.Error(),nil)
		return
	}

	req := &StringReq{}

	err = json.Unmarshal(b,req)
	if err != nil {
		s.stringResult(-4000,"request json parse err:"+err.Error(),nil)
		return
	}

	if req.Key == "" {
		s.stringResult(-5000,"request miss key param",nil)
		return
	}

	switch req.Operate {
	case GET:
		hasFound := false
		for k,v := range s.Data {
			if k == req.Key {
				s.stringResult(0,"ok", v)
				hasFound = true
			}
		}
		if !hasFound {
			s.stringResult(-5500,"get key:"+req.Key+" value not found",nil)
		}
		break
	case SET:
		if req.Value == "" {
			s.stringResult(-6000,"request set operation miss value param",nil)
			return
		}
		hasCheck := false
		for k,_ := range s.Data {
			if k == req.Key {
				s.Data[k] = req.Value
				hasCheck = true
				break
			}
		}
		if !hasCheck {
			s.Data[req.Key] = req.Value
		}
		s.stringResult(0,"set ok",req.Key)
		break
	default:
		s.stringResult(-7000,"operation not match",req.Operate)
		break
	}

}

func (s *String)stringResult(ret int, msg string, data interface{})  {
	rsp := &StringRsp{
		Ret:ret,
		Msg:msg,
		Data:data,
	}

	b,_ := json.Marshal(rsp)

	_,_ = s.w.Write(b)
}
