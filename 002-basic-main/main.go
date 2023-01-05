package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// implementing https://medium.com/@ramseyjiang_22278/golang-hot-reload-and-shutdown-gracefully-2024b0016c1f

type Config struct {
	Message string
}

var conf = &Config{Message: "Before hot reload"}

func multiSigHandler(signal os.Signal) {
	switch signal {
	case syscall.SIGHUP:
		log.Println("Signal:", signal.String())
		log.Println("After hot reload")
		conf.Message = "Hot reload has been finished."
	case syscall.SIGINT:
		log.Println("Signal:", signal.String())
		log.Println("Interrupted by Ctrl+C")
		os.Exit(0)
	case syscall.SIGTERM:
		log.Println("Signal:", signal.String())
		log.Println("Process is killed.")
		os.Exit(0)
	default:
		log.Println("Unhandled/unkonwn signal")
	}
}

func router() {
	log.Printf("Starting up with pid %d...", os.Getpid())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(conf.Message))
	})

	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
}

func main() {

	router()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	for {
		multiSigHandler(<-sigCh)
	}
}
