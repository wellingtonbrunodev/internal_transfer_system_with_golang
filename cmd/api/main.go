package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/go-chi/chi/v5"

	"github.com/wellingtonbrunodev/internal_transfer_system_with_golang/internal/database"
	"github.com/wellingtonbrunodev/internal_transfer_system_with_golang/internal/handler"
	"github.com/wellingtonbrunodev/internal_transfer_system_with_golang/internal/repository"
	"github.com/wellingtonbrunodev/internal_transfer_system_with_golang/internal/service"

)

func main() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbname,
	)

	if err := database.RunMigrations(dbURL); err != nil {
		log.Fatalf("migration error: %v", err)
	}

	db, err := database.NewConnection()
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	defer db.Close()

	accountRepo := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepo)
	accountHandler := handler.NewAccountHandler(accountService)

	transactionRepo := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	r := chi.NewRouter()

	r.Post("/accounts", accountHandler.CreateAccount)
	r.Get("/accounts/{accountID}", accountHandler.GetAccount)
	r.Post("/transactions", transactionHandler.Transfer)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Println("starting server on :8080")
	log.Fatal(server.ListenAndServe())

}
