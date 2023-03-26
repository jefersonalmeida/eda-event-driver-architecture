package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jefersonalmeida/go-wallet/internal/database"
	"github.com/jefersonalmeida/go-wallet/internal/event"
	"github.com/jefersonalmeida/go-wallet/internal/usecase/create_account"
	"github.com/jefersonalmeida/go-wallet/internal/usecase/create_client"
	"github.com/jefersonalmeida/go-wallet/internal/usecase/create_transaction"
	"github.com/jefersonalmeida/go-wallet/internal/web"
	"github.com/jefersonalmeida/go-wallet/internal/web/webserver"
	"github.com/jefersonalmeida/go-wallet/pkg/events"
	uow2 "github.com/jefersonalmeida/go-wallet/pkg/uow"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root", "root", "localhost", "3306", "wallet_core",
	))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	//eventDispatcher.Register("TransactionCreated", handler)

	clientDB := database.NewClientDB(db)
	accountDB := database.NewAccountDB(db)

	ctx := context.Background()
	uow := uow2.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})
	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createClientUseCase := create_client.NewCreateClientUseCase(clientDB)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDB, clientDB)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent)

	server := webserver.NewWebServer(":3000")

	clientHandler := web.NewClientHandler(*createClientUseCase)
	accountHandler := web.NewAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewTransactionHandler(*createTransactionUseCase)

	server.AddHandler("/clients", clientHandler.CreateClient)
	server.AddHandler("/accounts", accountHandler.CreateAccount)
	server.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Server is running")
	server.Start()
}
