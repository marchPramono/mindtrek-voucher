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

// Voucher for mind_voucher table
type Voucher struct {
	VoucherCode   string `json:"voucher_code"`
	ProductID     string `json:"product_id"`
	Nominal       string `json:"nominal"`
	CreatedAt     string `json:"created_at"`
	DurationMonth string `json:"duration_month"`
	ExpiredAt     string `json:"expired_at"`
	ActivatedAt   string `json:"activated_at"`
}

// Vouchers struct
type Vouchers struct {
	Vouchers []Voucher `json:"vouchers"`
}

// Product for mind_product table
type Product struct {
	ProductID          string `json:"product_id"`
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
}

// Products struct
type Products struct {
	Products []Product `json:"products"`
}

// Partner for mind_partner table
type Partner struct {
	PartnerID   string `json:"partner_id"`
	PartnerName string `json:"partner_name"`
	City        string `json:"city"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// Partners struct
type Partners struct {
	Partners []Partner `jason:"partners"`
}

// Invoice for mind_invoice table
type Invoice struct {
	InvoiceID       string `json:"invoice_id"`
	PartnerID       string `json:"partner_id"`
	PaymentMethodID string `json:"paymentmethod_id"`
	CreatedAt       string `json:"created_at"`
}

// Invoices struct
type Invoices struct {
	Invoices []Invoice `json:"invoices"`
}

// PartnerVoucher for mind_partner_voucher table
type PartnerVoucher struct {
	InvoiceID     string `json:"invoice_id"`
	PartnerID     string `json:"partner_id"`
	VoucherCode   string `json:"voucher_code"`
	PurchaseValue string `json:"purchase_value"`
}

// Generate Vouchers for partner purchase
type PartnerVouchers struct {
	PartnerVouchers []PartnerVoucher `json:"partner_vouchers"`
}

// InvoiceItem for invoice_item table
type InvoiceItem struct {
	InvoiceItemID string `json:"invoice_item_id"`
	ItemID        string `json:"item_id"`
	ItemAmount    int    `json:"item_amount"`
	ItemPrice     string `json:"item_price"`
	ItemDiscount  string `json:"item_discount"`
}

// Generate Invoices per Item
type InvoiceItems struct {
	InvoiceItems []InvoiceItem `json:"invoice_items"`
}

// InvoiceItemRelationship for invoice_item_relationship table
type InvoiceItemRelationship struct {
	InvoiceID     string `json:"invoice_id"`
	InvoiceItemID string `json:"invoice_item_id"`
}

type InvoiceItemRelationships struct {
	InvoiceItemRelationships []InvoiceItemRelationship `json:"invoice_item_relationship"`
}

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

func addVoucher(c echo.Context) error {

	u := new(Voucher)
	if err := c.Bind(u); err != nil {
		return err
	}

	sqlStatement := `INSERT INTO mind_voucher(product_id, voucher_code, nominal, duration_month, expired_at)
		VALUES($1, random_string(8), $2 , $3, now() + '1 month'::interval * 120)
		RETURNING voucher_code`

	res, err := db.Query(sqlStatement, u.ProductID, u.Nominal, u.DurationMonth)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusCreated, u)
	}
	return c.String(http.StatusOK, "ok")

}

func getVoucher(c echo.Context) error {

	voucher := Voucher{}
	id := c.Param("voucher_code")

	sqlStatement := `SELECT FROM mind_voucher WHERE voucher_code=$1`

	res, err := db.Query(sqlStatement, voucher)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusOK, "Selected")
	}
	return c.String(http.StatusOK, id+"Selected")
}

func addProduct(c echo.Context) error {

	u := new(Product)
	if err := c.Bind(u); err != nil {
		return err
	}

	sqlStatement := `INSERT INTO mind_product(product_name, product_description)
	VALUES($1, $2)`

	res, err := db.Query(sqlStatement, u.ProductName, u.ProductDescription)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusCreated, u)
	}
	return c.String(http.StatusOK, "ok")
}

func addPartner(c echo.Context) error {

	u := new(Partner)
	if err := c.Bind(u); err != nil {
		return err
	}

	sqlStatement := `INSERT INTO mind_partner(partner_name, city)
	VALUES($1, $2)`

	res, err := db.Query(sqlStatement, u.PartnerName, u.City)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusCreated, u)
	}
	return c.String(http.StatusOK, "ok")
}

func addInvoiceItem(c echo.Context) error {

	u := new(InvoiceItem)
	if err := c.Bind(u); err != nil {
		return err
	}

	itemsID := u.ItemID
	itemsAmount := u.ItemAmount

	for item, amount := range itemsID {
		for i := 0; i < u.ItemAmount; i++ {
			return itemsAmount
		}
		return itemsID
	}

	sqlStatement := `INSERT INTO mind_invoice_item(item_id, item_amount, item_price, item_discount)
	VALUES($1, $2, $3, #4)`

	res, err := db.Query(sqlStatement, u.ItemID, u.ItemAmount, u.ItemPrice, u.ItemDiscount)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusCreated, u)
	}
	return c.String(http.StatusOK, "ok")
}

func getAllVouchers(c echo.Context) error {
	sqlStatement := `SELECT invoice_id, voucher_code, partner_id, purchase_value 
	FROM mind_partner_voucher ORDER BY invoice_id`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	result := PartnerVouchers{}

	for rows.Next() {

		voucher := PartnerVoucher{}

		err2 := rows.Scan(&voucher.InvoiceID, &voucher.PartnerID, &voucher.VoucherCode, &voucher.PurchaseValue)
		if err2 != nil {
			return err2
		}
		result.PartnerVouchers = append(result.PartnerVouchers, voucher)
	}
	return c.JSON(http.StatusAccepted, result)
}
