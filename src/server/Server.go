package server

import "sync"

type Server struct {
	Port int
	Ip string
	Mux sync.Mutex
}
