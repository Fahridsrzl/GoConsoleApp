package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "001213"
	dbname   = "enigma_laundry"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. View Customer")
		fmt.Println("2. Insert Customer")
		fmt.Println("3. Update Customer")
		fmt.Println("4. Delete Customer")
		fmt.Println("5. View Transaksi")
		fmt.Println("6. Insert Transaksi")
		fmt.Println("7. Update Transaksi")
		fmt.Println("8. Delete Transaksi")
		fmt.Println("9. View Detail Transaksi")
		fmt.Println("10. Insert Detail Transaksi")
		fmt.Println("11. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			viewCustomer()
		case 2:
			insertCustomer()
		case 3:
			updateCustomer()
		case 4:
			deleteCustomer()
		case 5:
			viewTransaksi()
		case 6:
			insertTransaksi()
		case 7:
			updateTransaksi()
		case 8:
			deleteTransaksi()
		case 9:
			viewDetailTransaksi()
		case 10:
			insertDetailTransaksi()
		case 11:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func viewCustomer() {
	rows, err := db.Query("SELECT * FROM Customer")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("\nCustomer Table:")
	fmt.Println("ID\tName\tPhone Number")

	for rows.Next() {
		var id int
		var name, phone string
		err := rows.Scan(&id, &name, &phone)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d\t%s\t%s\n", id, name, phone)
	}
}

func insertCustomer() {
	var name, phone string

	fmt.Print("Enter Customer Name: ")
	fmt.Scan(&name)

	if strings.TrimSpace(name) == "" {
		fmt.Println("Name cannot be empty.")
		return
	}

	fmt.Print("Enter Phone Number: ")
	fmt.Scan(&phone)

	_, err := strconv.ParseInt(phone, 12, 64)
	if err != nil {
		fmt.Println("Invalid phone number. Please enter a numeric value.")
		return
	}

	if len(phone) != 12 {
		fmt.Println("Invalid phone number length. It should have 12 digits.")
		return
	}

	_, err = db.Exec("INSERT INTO Customer (customer_name, phone_number) VALUES ($1, $2)", name, phone)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Customer inserted successfully.")
}

func updateCustomer() {
	var id int
	var newPhone string

	fmt.Print("Enter Customer ID to update: ")
	fmt.Scan(&id)

	if id <= 0 {
		fmt.Println("Invalid ID.")
		return
	}

	fmt.Print("Enter New Phone Number: ")
	fmt.Scan(&newPhone)

	_, err := strconv.ParseInt(newPhone, 12, 64)
	if err != nil {
		fmt.Println("Invalid new phone number. Please enter a numeric value.")
		return
	}

	if len(newPhone) != 12 {
		fmt.Println("Invalid new phone number length. It should have 12 digits.")
		return
	}

	result, err := db.Exec("UPDATE Customer SET phone_number = $1 WHERE customer_id = $2", newPhone, id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected == 0 {
		fmt.Println("Customer not found.")
	} else {
		fmt.Println("Customer updated successfully.")
	}
}

func deleteCustomer() {
	var id int

	fmt.Print("Enter Customer ID to delete: ")
	fmt.Scan(&id)

	if id <= 0 {
		fmt.Println("Invalid ID.")
		return
	}

	result, err := db.Exec("DELETE FROM Customer WHERE customer_id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected == 0 {
		fmt.Println("Customer not found.")
	} else {
		fmt.Println("Customer deleted successfully.")
	}
}

func viewTransaksi() {
	rows, err := db.Query("SELECT * FROM Transaksi")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("\nTransaksi Table:")
	fmt.Println("ID\tEntry Date\tCompletion Date\tReceived By\tCustomer ID")

	for rows.Next() {
		var id, customerID int
		var entryDate, completionDate, receivedBy string
		err := rows.Scan(&id, &entryDate, &completionDate, &receivedBy, &customerID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d\t%s\t%s\t%s\t%d\n", id, entryDate, completionDate, receivedBy, customerID)
	}
}

func insertTransaksi() {
	var entryDate, completionDate, receivedBy string
	var customerID int

	fmt.Print("Enter Entry Date (YYYY-MM-DD): ")
	fmt.Scan(&entryDate)

	if strings.TrimSpace(entryDate) == "" {
		fmt.Println("Entry date cannot be empty.")
		return
	}

	fmt.Print("Enter Completion Date (YYYY-MM-DD): ")
	fmt.Scan(&completionDate)

	if strings.TrimSpace(completionDate) == "" {
		fmt.Println("Completion date cannot be empty.")
		return
	}

	fmt.Print("Enter Received By: ")
	fmt.Scan(&receivedBy)

	if strings.TrimSpace(receivedBy) == "" {
		fmt.Println("Received by cannot be empty.")
		return
	}

	fmt.Print("Enter Customer ID: ")
	fmt.Scan(&customerID)

	if customerID <= 0 {
		fmt.Println("Invalid Customer ID.")
		return
	}

	_, err := db.Exec("INSERT INTO Transaksi (entry_date, completion_date, received_by, customer_id) VALUES ($1, $2, $3, $4)", entryDate, completionDate, receivedBy, customerID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaksi inserted successfully.")
}

func updateTransaksi() {
	var id, customerID int
	var entryDate, completionDate, receivedBy string

	fmt.Print("Enter Transaksi ID to update: ")
	fmt.Scan(&id)

	if id <= 0 {
		fmt.Println("Invalid ID.")
		return
	}

	fmt.Print("Enter Entry Date (YYYY-MM-DD): ")
	fmt.Scan(&entryDate)

	if strings.TrimSpace(entryDate) == "" {
		fmt.Println("Entry date cannot be empty.")
		return
	}

	fmt.Print("Enter Completion Date (YYYY-MM-DD): ")
	fmt.Scan(&completionDate)

	if strings.TrimSpace(completionDate) == "" {
		fmt.Println("Completion date cannot be empty.")
		return
	}

	fmt.Print("Enter Received By: ")
	fmt.Scan(&receivedBy)

	if strings.TrimSpace(receivedBy) == "" {
		fmt.Println("Received by cannot be empty.")
		return
	}

	fmt.Print("Enter Customer ID: ")
	fmt.Scan(&customerID)

	if customerID <= 0 {
		fmt.Println("Invalid Customer ID.")
		return
	}

	result, err := db.Exec("UPDATE Transaksi SET entry_date = $1, completion_date = $2, received_by = $3, customer_id = $4 WHERE transaction_id = $5",
		entryDate, completionDate, receivedBy, customerID, id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected == 0 {
		fmt.Println("Transaksi not found.")
	} else {
		fmt.Println("Transaksi updated successfully.")
	}
}

func deleteTransaksi() {
	var id int

	fmt.Print("Enter Transaksi ID to delete: ")
	fmt.Scan(&id)

	if id <= 0 {
		fmt.Println("Invalid ID.")
		return
	}

	result, err := db.Exec("DELETE FROM Transaksi WHERE transaction_id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected == 0 {
		fmt.Println("Transaksi not found.")
	} else {
		fmt.Println("Transaksi deleted successfully.")
	}
}

func viewDetailTransaksi() {
	rows, err := db.Query("SELECT * FROM detail_transaksi")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("\nDetail Transaksi Table:")
	fmt.Println("ID\tTransaction ID\tService Name\tQuantity\tUnit\tPrice\tTotal")

	for rows.Next() {
		var id, transactionID, quantity int
		var serviceName, unit string
		var price, total float64
		err := rows.Scan(&id, &transactionID, &serviceName, &quantity, &unit, &price, &total)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d\t%d\t%s\t%d\t%s\t%.2f\t%.2f\n", id, transactionID, serviceName, quantity, unit, price, total)
	}
}

func insertDetailTransaksi() {
	var serviceName, unit string
	var quantity, transactionID int
	var price, total float64

	fmt.Print("Enter Service Name: ")
	fmt.Scan(&serviceName)

	if strings.TrimSpace(serviceName) == "" {
		fmt.Println("Service Name cannot be empty.")
		return
	}

	fmt.Print("Enter Quantity: ")
	fmt.Scan(&quantity)

	if quantity <= 0 {
		fmt.Println("Invalid Quantity.")
		return
	}

	fmt.Print("Enter Unit: ")
	fmt.Scan(&unit)

	if strings.TrimSpace(unit) == "" {
		fmt.Println("Unit cannot be empty.")
		return
	}

	fmt.Print("Enter Price: ")
	fmt.Scan(&price)

	if price <= 0 {
		fmt.Println("Invalid Price.")
		return
	}

	fmt.Print("Enter Total: ")
	fmt.Scan(&total)

	if total <= 0 {
		fmt.Println("Invalid Total.")
		return
	}

	fmt.Print("Enter Transaction ID: ")
	fmt.Scan(&transactionID)

	if transactionID <= 0 {
		fmt.Println("Invalid Transaction ID.")
		return
	}

	_, err := db.Exec("INSERT INTO detail_transaksi (transaction_id, service_name, quantity, unit, price, total) VALUES ($1, $2, $3, $4, $5, $6)",
		transactionID, serviceName, quantity, unit, price, total)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Detail Transaksi inserted successfully.")
}
