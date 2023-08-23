package main

import (
	"br/simple-service/db"
	"br/simple-service/db/sqlc"
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
	l := log.New(os.Stdout, "BR-SERVER-", log.LstdFlags)
	ctx := context.Background()

	// CRUD
	db, queries := db.Instantiate(l)
	if db == nil || queries == nil {
		l.Println("Exiting due to database connection error")
		return
	}
	defer db.Close()

	server := &http.Server{
		Addr:        ":" + "4444",
		Handler:     defineMultiplexer(l, queries),
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
	defer cancel()

	stopServer(server, l, &timeout_ctx, &cancel)

}

func startServer(s *http.Server, l *log.Logger) {
	l.Println("🔥 Server is starting on", s.Addr)

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

func defineMultiplexer(l *log.Logger, q *sqlc.Queries) *http.ServeMux {
	// reference to the handler
	hello_handler := handlers.NewHello(l)
	account_handler := handlers.NewAccountHandler(l, q)
	deck_handler := handlers.NewDeckHandler(l, q)
	card_handler := handlers.NewCardHandler(l, q)

	// handle multiplexer
	mux := http.NewServeMux()
	mux.Handle("/hello", hello_handler)

	// account crud
	mux.HandleFunc("/account/get", account_handler.GetAccountH)
	mux.HandleFunc("/account/getAll", account_handler.ListAccountsH)
	mux.HandleFunc("/account/create", account_handler.CreateAccountH)
	mux.HandleFunc("/account/delete", account_handler.DeleteAccountH)

	// deck crud
	mux.HandleFunc("/deck/get", deck_handler.GetDeckH)
	mux.HandleFunc("/deck/create", deck_handler.CreateDeckH)
	mux.HandleFunc("/deck/delete", deck_handler.DeleteDeckH)
	mux.HandleFunc("/deck/update", deck_handler.UpdateDeckTitleH)
	mux.HandleFunc("/deck/getAll", deck_handler.ListDecksByAccountH)

	// card crud
	mux.HandleFunc("/card/get", card_handler.GetCardH)
	mux.HandleFunc("/card/create", card_handler.CreateCardH)
	mux.HandleFunc("/card/delete", card_handler.DeleteCardH)
	mux.HandleFunc("/card/update", card_handler.UpdateCardH)
	mux.HandleFunc("/card/getAll", card_handler.ListCardsH)

	return mux
}
