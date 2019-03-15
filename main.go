package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func init() {

	connStr := "host=localhost port=5432 user=postgres password=nav123 dbname=voucher_db sslmode=disable"
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB connected...")
	}

}

func main() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server connected!")
	})

	e.POST("/mind_voucher", addVoucher)
	e.POST("/mind_product", addProduct)
	e.POST("/mind_partner", addPartner)
	e.POST("/mind_invoice_item", addInvoiceItem)
	e.GET("/mind_voucher/:code", getVoucher)
	e.GET("/mind_voucher", getAllVouchers)

	e.Logger.Fatal(e.Start(":2005"))
}
