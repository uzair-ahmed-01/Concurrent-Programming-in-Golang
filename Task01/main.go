package main

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to MySQL
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/task")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// Set the RequiredAcks and IdempotentProducer configuration options
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(0)
	db.Exec("SET SESSION binlog_format = 'ROW'")

	// Create a slice of int
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// Use a wait group to wait for all goroutines to complete
	var wg sync.WaitGroup
	wg.Add(len(numbers))

	// Iterate the slice and concurrently insert each number into the "numbers" table
	for _, number := range numbers {
		go func(number int) {
			defer wg.Done()

			// Use recover to catch panics
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Recovered from panic: %v\n", r)
				}
			}()

			_, err := db.Exec("INSERT INTO numbers (number) VALUES (?)", number)
			if err != nil {
				log.Fatalln(err)
			}
		}(number)
	}

	// Wait for all goroutines to complete
	wg.Wait()
}
