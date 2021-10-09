package main

import (
	"fmt"
	"github.com/bdemirpolat/test/pkg/email"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
)



func main() {
	// New email service
	svc := email.NewService()

	// Make email endpoints
	eps := email.MakeEndpoints(svc,nil)

	// router
	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/send").Handler(email.GetSendHandler(eps.Send,nil))

	// http server
	svr := http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	go func() {
		err := svr.ListenAndServe()
		if err != nil {
			log.Fatalf("Listening on port %s failed, err: %s\n",svr.Addr,err.Error())
		}
	}()

	interruptChan := make(chan os.Signal,1)
	signal.Notify(interruptChan,os.Interrupt)
	<-interruptChan
	fmt.Println("interrupting...")
}


