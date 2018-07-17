package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize Init db connectio to SQL db
func (a *App) Initialize(user, password, dbname string) {}

// Run run our server by addr
func (a *App) Run(addr string) {}
