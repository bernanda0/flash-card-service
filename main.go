package main

import (
	"br/simple-service/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	// now the startServer is run by a routine
	go startServer(server, l)

	// inorder to block the routine, we might use a channel (we can use wait group also)
	shut := make(chan os.Signal, 1)
	signal.Notify(shut, syscall.SIGINT, syscall.SIGTERM)

	<-shut // Block until a signal is received

	timeout_ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	stopServer(server, l, &timeout_ctx)

}

func startServer(s *http.Server, l *log.Logger) {
	l.Println("🔥 Server is running on", s.Addr)
	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		l.Fatalln("Server is failed due to", err)
	}
}

func stopServer(s *http.Server, l *log.Logger, ctx *context.Context) {
	l.Println("💅 Shutting down the server")
	s.Shutdown(*ctx)
}
