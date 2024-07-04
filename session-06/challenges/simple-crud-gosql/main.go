package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

const (
	user     = "root"
	password = ""
	dbName   = "db_mc6"
)

type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Variant struct {
	ID          int
	VariantName string
	Quantity    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func main() {
	// add the parseTime=true parameter
	mysqlInfo := fmt.Sprintf("%s:%s@/%s?parseTime=true", user, password, dbName)
	db, err = sql.Open("mysql", mysqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping() // wajib dipanggil
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	// handler atau controller
	// createProduct()
    // updateProduct(5, "Mouse Fantech")
    // getProductById(40)
    createVariant(4, "RGB Blue Switch", 30)
    // updateVariantById(1, 54)
    // deleteVariantById(1)
    getProductWithVariant(4)

}

func createProduct() {
	var product = Product{}

	sqlStatement := `INSERT INTO products (name, created_at, updated_at) VALUES (?, NOW(), NOW())`

	// execute raw query
	result, err := db.Exec(sqlStatement, "Mechanical Keyboard Keychron")
	if err != nil {
		panic(err)
	}

	// check last inserted id
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	// retreive data
	sqlRetrieve := `SELECT * FROM products WHERE id = ?`
	err = db.QueryRow(sqlRetrieve, lastInsertId).Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		panic(err)
	}

	fmt.Printf("New product inserted: %+v \n", product)
}

func updateProduct(id int, newName string) {
	sqlStatement := `UPDATE products SET name = ?, updated_at = NOW() WHERE id = ?`
	_, err := db.Exec(sqlStatement, newName, id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Product with ID %d successfully updated: %s \n", id, newName)
}

func getProductById(id int) {
	var product Product

	sqlStatement := `SELECT * FROM products WHERE id = ?`

	err := db.QueryRow(sqlStatement, id).Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("No product found with id %+v.", id)
		} else {
			panic(err)
		}
		return
	}

	fmt.Printf("Product retreived: %+v\n", product)
}

func createVariant(productId int, variantName string, quantity int) {
	sqlStatement := `INSERT INTO variants (product_id, variant_name, quantity, created_at, updated_at) VALUES (?, ?, ?, NOW(), NOW())`
		_, err := db.Exec(sqlStatement, productId, variantName, quantity)
		if err != nil {
			panic(err)
		}

		fmt.Printf("New variant created for product ID %d: %s, Quantity: %d\n", productId, variantName, quantity)
}

func updateVariantById(id int, newQuantity int) {
	sqlStatement := `UPDATE variants SET quantity = ?, updated_at = NOW() WHERE id = ?`
	_, err := db.Exec(sqlStatement, newQuantity, id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Variant with ID %d updated to new quantity: %d\n", id, newQuantity)
}

func deleteVariantById(id int) {
	sqlStatement := `DELETE FROM variants WHERE id = ?`
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Variant with ID %d deleted\n", id)
}

func getProductWithVariant(productId int) {
	var product Product
	sqlStatement := `SELECT id, name, created_at, updated_at FROM products WHERE id = ?`
	err := db.QueryRow(sqlStatement, productId).Scan(&product.ID, &product.Name, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No product found with the given ID")
		} else {
			panic(err)
		}
		return
	}

	fmt.Printf("Product retrieved: %+v\n", product)

	var variants []Variant
	sqlStatement = `SELECT id, variant_name, quantity, created_at, updated_at FROM variants WHERE product_id = ?`
	rows, err := db.Query(sqlStatement, productId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var variant Variant
		err := rows.Scan(&variant.ID, &variant.VariantName, &variant.Quantity, &variant.CreatedAt, &variant.UpdatedAt)
		if err != nil {
			panic(err)
		}
		variants = append(variants, variant)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Variants retrieved: %+v\n", variants)
}
