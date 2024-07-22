package main

import (
	"database/sql"
	"fmt"
	"log"
    "os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Please specify a database")
    }

    db_path := os.Args[1]

	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		log.Fatal(err)
	}
    defer db.Close()

    var input string
    InputLoop:
    for {
        fmt.Print("> ")

        _, err := fmt.Scanf("%s", &input)
        if err != nil {
            fmt.Println("Error reading input:", err)
            continue
        }

        switch input {
        case "exit", "quit", "q":
            break InputLoop
        case "help":
            print_help()
        case "init":
            init_db(db)
        case "print_invoices":
            print_invoices(db)
        default:
            fmt.Println("Invalid command. Run `help` to see a list of available commands.")
            continue
        }
    }
}

func print_help() {
    // TODO: print a list of available commands
}

func init_db(db *sql.DB) {
    fmt.Println("Are you sure you want to initialize this database? (This will overwrite any existing data!) [Y/n]")
    fmt.Print("> ")
    
    var input string
    _, errr := fmt.Scanf("%s", &input)
    if errr != nil {
        log.Fatal(errr)
    }
    if input == "N" || input == "n" {
        fmt.Println("Cancelled init.")
        return
    }

    // Create invoices table
    _, err := db.Exec(`CREATE TABLE invoices (
        invoice_no STRING PRIMARY KEY,
        date DATE,
        description STRING,
        customer_name STRING,
        customer_no INT,
        status STRING,
        price REAL,
        vat_percentage REAL,
        vat REAL,
        total REAL,
        order_id
    )`)
    if err != nil {
        log.Fatal(err)
    }

    // TODO: Create transactions table
}

func print_invoices(db *sql.DB) {
    rows, err := db.Query(fmt.Sprintf("SELECT * FROM %q", "invoices"))
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    fmt.Println("=== Invoices ===")
    fmt.Println("INV #, DATE, DESCRIPTION, CUSTOMER, STATUS, TOTAL (€)")
    for rows.Next() {
        var invoice_no string
        var date string
        var description string
        var customer_name string
        var customer_no int
        var status string
        var price float32
        var vat_percentage float32
        var vat float32
        var total float32
        var order_id int

        err := rows.Scan(&invoice_no, &date, &description, &customer_name, &customer_no, &status, &price, &vat_percentage, &vat, &total, &order_id)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("%v, %v, %v, %v (%v), %v, %v\n", invoice_no, date, description, customer_name, customer_no, status, total)
    }
}

/*
func add_invoice(db *sql.DB) {
    _, err := db.Exec(`INSERT`)
}
*/

