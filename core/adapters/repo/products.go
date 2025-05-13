package repo

import (
	"CaliYa/core/domain/models"
	"CaliYa/core/domain/ports"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type productsRepo struct {
	db *bun.DB
}

func NewProductsRepo(db *bun.DB) ports.ProductsRepo {
	return &productsRepo{db}
}

// func newConnection() *bun.DB {
// 	dsn := "postgres://test:test@localhost:5432/caliya?sslmode=disable"
// 	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
// 	db := bun.NewDB(sqldb, pgdialect.New())
// 	db.AddQueryHook(bundebug.NewQueryHook(
// 		bundebug.WithVerbose(true), // Imprime todas las consultas
// 	))

// 	return db
// }

func (p *productsRepo) GetProducts(ctx context.Context) (*models.ProductsShops, error) {

	products := new(models.ProductsShops)

	if err := p.db.NewSelect().
		Model(products).
		Where("opened = ?", true).
		Relation("Categories").
		Relation("Categories.Items", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.OrderExpr("price ASC")
		}).
		Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No rows found")
			return products, echo.NewHTTPError(http.StatusNotFound, "not found")
		}
	}

	return products, nil
}
