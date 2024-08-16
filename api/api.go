package api

import (
	"PassportScanner/api/handler"
	"PassportScanner/scannersdk"
	"PassportScanner/store"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type API struct {
	port         int
	app          *gin.Engine
	queries      *store.Queries
	contactStore handler.ContactStore
	scanner      scannersdk.ScannerSDK
}

func New() *API {
	connStr := "postgres://postgres:secret@localhost:5432/sqlctest?sslmode=disable"

	dbpool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbpool.Close()

	return &API{
		port:         4000,
		app:          gin.Default(),
		queries:      store.New(dbpool),
		scanner:      scannersdk.NewScannerSDK(),
		contactStore: handler.NewContactStore(),
	}
}

func (a *API) StartServer() {
	a.app.Static("/public", "./public")
	a.registerMiddlewares()
	a.registerRoutes()

	port := fmt.Sprintf(":%d", a.port)
	a.app.Run(port)
}

func (a *API) registerMiddlewares() {
}
