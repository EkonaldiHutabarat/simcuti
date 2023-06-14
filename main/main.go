package main

import (
	"simcuti/config/database"
)

func main() {
	// Inisialisasi koneksi MySQL
	database.InitMySQL()
	defer database.DB.Close()

	// Melakukan operasi database lainnya di sini...
}
