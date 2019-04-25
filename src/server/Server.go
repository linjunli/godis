package server

import (
	"godis/src/server/handler"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Server struct {
	Port int
	Ip string
	server *http.Server
	data map[string]interface{}
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server)Start()  {
	s.Ip = "127.0.0.1"
	s.Port = 8888
	s.data = make(map[string]interface{},100)

	mux := http.NewServeMux()
	mux.Handle("/server/ping", &handler.Ping{
		Data: s.data,
	})
	mux.Handle("/server/string", &handler.String{
		Data: s.data,
	})
	mux.Handle("/server/list", &handler.List{
		Data: s.data,
	})

	s.server = &http.Server{
		Addr        : s.Ip + ":" + strconv.Itoa(s.Port),
		ReadTimeout : 60 * time.Second,
		WriteTimeout: 60 * time.Second,
		Handler		: mux,
	}

	go func() {
		err := s.server.ListenAndServe()
		if err != nil {
			os.Exit(-1)
		}
	}()
}

func (s *Server)Stop()  {
	_ = s.server.Close()
}