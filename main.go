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
        case "print_table":
            print_table(db, "users")
        }
    }

    defer db.Close()
}

func print_help() {
    // TODO: print a list of available commands
}

func print_table(db *sql.DB, table string) {
    rows, err := db.Query(fmt.Sprintf("SELECT * FROM %q", table))
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    fmt.Println("=== users ===")
    for rows.Next() {
        var id int
        var name string
        var email string

        err := rows.Scan(&id, &name, &email)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("id %d, name %q, email %q\n", id, name, email)
    }
}

