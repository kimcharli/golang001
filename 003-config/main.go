package main

import (
	"log"
	"net/http"
	"os"
	"syscall"

	"github.com/kimcharli/go101/config"
)

// implementing https://medium.com/@ramseyjiang_22278/golang-hot-reload-and-shutdown-gracefully-2024b0016c1f

var conf = config.Config{Name: "default"}

func multiSigHandler(signal os.Signal) {
	switch signal {
	case syscall.SIGHUP:
		log.Println("Signal:", signal.String())
		log.Println("After hot reload")
		conf.Name = "Hot reload has been finished."
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(conf.Name))
	})

	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("Starting up with pid %d...", os.Getpid())

	conf.LoadYaml("config.yaml")
	conf.DumpYaml("config-out.yaml")

	// router()
	// sigCh := make(chan os.Signal, 1)
	// signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	// for {
	// 	multiSigHandler(<-sigCh)
	// }

}
