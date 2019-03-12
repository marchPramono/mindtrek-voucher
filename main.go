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

// Voucher for mind_voucher
type Voucher struct {
	VoucherCode   string `json:"voucher_code"`
	ProductID     string `json:"product_id"`
	Nominal       string `json:"nominal"`
	CreatedAt     string `json:"created_at"`
	DurationMonth string `json:"duration_month"`
	ExpiredAt     string `json:"expired_at"`
	ActivatedAt   string `json:"activated_at"`
}

// Product for mind_product
type Product struct {
	ProductID          string `json:"product_id"`
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
}

// Partner for mind_partner
type Partner struct {
	PartnerID   string `json:"partner_id"`
	PartnerName string `json:"partner_name"`
	City        string `json:"city"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
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
	//e.GET("/mind_voucher/:id", getVoucher)
	//e.GET("/mind_voucher", getAllVouchers)

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

/*
func getVoucher(c echo.Context) error {

	id := c.Param("id")

	sqlStatement := `SELECT FROM mind_voucher WHERE id=$1`

	res, err := db.Query(sqlStatement, id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusOK, "Selected")
	}
	return c.String(http.StatusOK, id+"Selected")
}
*/

/*
func getAllVouchers(c echo.Context) error {
	sqlStatement := ``
	rows, err := db.Query(sqlStatement)
	if err != nil {
	fmt.Println(err)
	}
	defer rows.Close()
	result := Vouchers{}

			for rows.next() {
				voucher := Vouchers{}
				err 2 := rows.Scan()
				if err2 != nil {
					return err2
				}
				result.Vouchers = append(result.Vouchers, voucher)
			}
			result c.JSON(http.StatusCreated, result)

	}
*/
