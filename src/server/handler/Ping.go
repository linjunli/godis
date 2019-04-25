package handler

import "net/http"

type Ping struct {
	Data map[string]interface{}
}

func (p *Ping)ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	_,_ = w.Write([]byte("dsad"))
}