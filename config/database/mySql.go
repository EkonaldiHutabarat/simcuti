package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

func InitMySQL() error {
	// Konfigurasi koneksi MySQL
	db, err := sql.Open("mysql", "root:manasge@tcp(localhost)/mat_db")
	if err != nil {
		return fmt.Errorf("gagal membuka koneksi ke MySQL: %v", err)
	}

	// Menguji koneksi ke database
	err = db.Ping()
	if err != nil {
		db.Close()
		return fmt.Errorf("koneksi ke MySQL gagal: %v", err)
	}

	fmt.Println("Koneksi ke MySQL sukses!")

	DB = db
	return nil
}
