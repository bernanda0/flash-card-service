package main

import (
	"br/simple-service/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	l := log.New(os.Stdout, "SERVER-", log.LstdFlags)

	// reference to the handler
	hello_handler := handlers.NewHello(l)

	// handle multiplexer
	mux := http.NewServeMux()
	mux.Handle("/hello", hello_handler)

	server := &http.Server{
		Addr:        "localhost:4444",
		Handler:     mux,
		IdleTimeout: 30 * time.Second,
		ReadTimeout: time.Second,
	}

	go startServer(server, l)

}

func startServer(s *http.Server, l *log.Logger) {
	l.Println("Server is running on", s.Addr)
	err := s.ListenAndServe()
	if err != nil {
		l.Println("Failed to start the server due to ", err)
	}
}
