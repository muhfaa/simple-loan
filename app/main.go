package main

import (
	"context"

	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/jmoiron/sqlx"
	echo "github.com/labstack/echo/v4"

	"loanapp/modules/mysql"

	loanController "loanapp/api/v1/loan"
	loanService "loanapp/business/loan"
	config "loanapp/config"
	loanModule "loanapp/modules/loan"

	loanApp "loanapp/api/v1"

	"github.com/labstack/gommon/log"
)

func newDatabaseConnection() *sqlx.DB {
	uri := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		config.GetConfig().Mysql.User,
		config.GetConfig().Mysql.Password,
		config.GetConfig().Mysql.Host,
		config.GetConfig().Mysql.Port,
		config.GetConfig().Mysql.Name)

	db := mysql.NewDatabaseConnection(uri)

	return db
}

func newLoanService(dbsql *sqlx.DB) loanService.Service {
	loanRepoSQL := loanModule.NewMySQLRepository(dbsql)

	return loanService.NewService(loanRepoSQL)
}

func main() {

	// database init
	dbsql := newDatabaseConnection()

	// service init
	loanService := newLoanService(dbsql)

	// controller init
	loanControllerV1 := loanController.NewController(loanService)

	e := echo.New()

	//register API path and handler
	loanApp.RegisterPath(
		e,
		*loanControllerV1)

	// run server
	go func() {
		address := fmt.Sprintf("0.0.0.0:%d", config.GetConfig().BackendPort)

		if err := e.Start(address); err != nil {
			log.Error("failed to start server", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// Wait for interrupt signal to gracefully shutdown the server with
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Error("failed to grafefully shutting down server", err)
	}

}
