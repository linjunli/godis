package main

import (
	"fmt"
	"godis/src/server"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	file, err := os.OpenFile("godis.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		os.Exit(-1)
	}
	defer func() {
		_ = file.Close()
	}()

	os.Stderr, os.Stdout = file, file
	srv := server.NewServer()
	srv.Start()

	fmt.Println("srv started")
	//_,_,err = godaemon.MakeDaemon(&godaemon.DaemonAttr{})
	//if err != nil {
	//	fmt.Println("godaemon err:"+err.Error())
	//	os.Exit(-1)
	//}

	//service, err := daemon.New("name","description")
	//if err != nil {
	//	log.Fatal("Error:",err)
	//}
	//status, err := service.Install()
	//if err != nil {
	//	log.Fatal(status,"\nError:",err)
	//}

	c := make(chan os.Signal)
	signal.Notify(c)
	signal.Ignore(syscall.SIGCHLD, syscall.SIGPIPE, syscall.SIGHUP)

	<-c

	srv.Stop()

	fmt.Println("srv stopped")
}