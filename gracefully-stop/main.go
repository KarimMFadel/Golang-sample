package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// referance: https://kpbird.medium.com/golang-gracefully-stop-application-23c2390bb212

func main() {

	args := os.Args
	if args != nil {
		fmt.Printf("Type of Args = %T\n", args)
		for _, arg := range args {
			fmt.Println(arg)
		}
	}


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		   fmt.Fprint(w,"Server is running")
	})
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		   sig := <-gracefulStop
		   fmt.Printf("caught sig: %+v", sig)
		   fmt.Println("Wait for 2 second to finish processing")
		   time.Sleep(2*time.Second)
		   os.Exit(0) 
	}()
	http.ListenAndServe(":8080",nil)
}