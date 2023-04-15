// Table related methods and utilities
package table

import (
	"bytes"
	"fmt"
	"github/com/codecrafters-io/sqlite-starter-go/app/page"
	"github/com/codecrafters-io/sqlite-starter-go/app/row"
	"github/com/codecrafters-io/sqlite-starter-go/app/utils"
)

const (
	TableColumnString = iota
	TableColumnInteger
)

type TableColumn struct {
	Title string
	Type  int
}

type Table struct {
	Title   string
	Columns []TableColumn
}

func (t Table) ReadRow(page *page.Page, index int) (row.Row, error) {
	var row row.Row
	reader := bytes.NewReader((*page)[index:])
	length, err := utils.ReadVarInt(reader)
	fmt.Println(length)
	if err != nil {
		return nil, err
	}
	reader = bytes.NewReader((*page)[index : index+int(length)+1])

	for _, c := range t.Columns {
		if c.Type == TableColumnString {
			length, err := utils.ReadVarInt(reader)
			fmt.Println(length)
			if err != nil {
				return nil, err
			}
			str := make([]byte, length)
			_, err = reader.Read(str)
			if err != nil {
				return nil, err
			}
			fmt.Println(string(str))
		}
	}
	return row, nil
}
