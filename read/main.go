package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"npm-repo/helpers"
	"npm-repo/routes"
)

func main() {
	var rt routes.Route
	r := rt.Init()

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v", sig)
		fmt.Println("Wait for 2 second to finish processing")
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()

	port := helpers.GetEnv("PORT", "8080")
	fmt.Println("Server served at port " + port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Unable to start service: " + err.Error())
	}

}
