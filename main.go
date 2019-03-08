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
	ProductID   string `json:"product_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// Partner for mind_partner
type Partner struct {
	PartnerID string `json:"partner_id"`
	Name      string `json:"name"`
	City      string `json:"city"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func init() {

	connStr := "host=localhost port=5432 user=postgres password=FR4uT dbname=voucher_db sslmode=disable"
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
	e.GET("/mind_voucher/:id", getVoucher)
	//e.GET("/mind_voucher", getAllVouchers)

	e.Logger.Fatal(e.Start(":2005"))
}

func addVoucher(c echo.Context) error {

	u := new(Voucher)
	if err := c.Bind(u); err != nil {
		return err
	}

	sqlStatement := `INSERT INTO mind_voucher(voucher_code, product_id, nominal, created_at, duration_month)
	VALUES($1, $2, $3, $4, $5)`

	res, err := db.Query(sqlStatement, u.VoucherCode, u.ProductID, u.Nominal, u.CreatedAt, u.DurationMonth)
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

	sqlStatement := `INSERT INTO mind_product(product_id, name, description, created_at, updated_at)
	VALUES($1, $2, $3, $4, $5)`

	res, err := db.Query(sqlStatement, u.ProductID, u.Name, u.Description, u.CreatedAt, u.UpdatedAt)
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

	sqlStatement := `INSERT INTO mind_partner(partner_id, name, city, created_at, updated_at)
	VALUES($1, $2, $3, $4, $5)`

	res, err := db.Query(sqlStatement, u.PartnerID, u.Name, u.City, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusCreated, u)
	}
	return c.String(http.StatusOK, "ok")
}

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
