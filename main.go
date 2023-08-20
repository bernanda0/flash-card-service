package main

import (
	"br/simple-service/db"
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
	ctx := context.Background()

	// CRUD
	db, queries := db.Instantiate(l)
	if db == nil || queries == nil {
		l.Println("Exiting due to database connection error")
		return
	}
	defer db.Close()

	// reference to the handler
	hello_handler := handlers.NewHello(l)
	account_handler := handlers.NewAccountHandler(l, queries)

	// handle multiplexer
	mux := http.NewServeMux()
	mux.Handle("/hello", hello_handler)

	// account crud
	mux.HandleFunc("/account/get", account_handler.GetAccount)
	mux.HandleFunc("/account/all", account_handler.ListAccounts)
	mux.HandleFunc("/account/create", account_handler.CreateAccount)
	mux.HandleFunc("/account/delete", account_handler.DeleteAccount)

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

	timeout_ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	stopServer(server, l, &timeout_ctx, &cancel)

}

func startServer(s *http.Server, l *log.Logger) {
	l.Println("🔥 Server is running on", s.Addr)
	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		l.Fatalln("Server is failed due to", err)
	}
}

func stopServer(s *http.Server, l *log.Logger, ctx *context.Context, cancel *context.CancelFunc) {
	l.Println("💅 Shutting down the server")
	s.Shutdown(*ctx)
	c := *cancel
	c()
}
