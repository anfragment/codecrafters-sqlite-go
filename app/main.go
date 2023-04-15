package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"

	// Available if you need it!
	// "github.com/pingcap/parser"
	// "github.com/pingcap/parser/ast"

	"github/com/codecrafters-io/sqlite-starter-go/app/page"
)

// Usage: your_sqlite3.sh sample.db .dbinfo
func main() {
	databaseFilePath := os.Args[1]
	command := os.Args[2]

	switch command {
	case ".dbinfo":
		databaseFile, err := os.Open(databaseFilePath)
		if err != nil {
			log.Fatal(err)
		}

		header := make([]byte, 100)

		_, err = databaseFile.Read(header)
		if err != nil {
			log.Fatal(err)
		}

		var pageSize uint16
		if err := binary.Read(bytes.NewReader(header[16:18]), binary.BigEndian, &pageSize); err != nil {
			fmt.Println("Failed to read integer:", err)
			return
		}

		schemaPage := page.Page(make([]byte, pageSize))
		_, err = databaseFile.Read(schemaPage)
		if err != nil {
			log.Fatal(err)
		}
		metadata, err := schemaPage.Metadata()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("database page size: %v\n", pageSize)
		fmt.Printf("number of tables: %d\n", metadata.Cells)
	default:
		fmt.Println("Unknown command", command)
		os.Exit(1)
	}
}
