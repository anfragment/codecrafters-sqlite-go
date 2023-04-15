package utils_test

import (
	"bytes"
	"fmt"
	"testing"

	"github/com/codecrafters-io/sqlite-starter-go/app/utils"
)

func TestReadVarInt(t *testing.T) {
	ba := []byte{0b10000100, 0b00101100}
	reader := bytes.NewReader(ba)
	fmt.Println(utils.ReadVarInt(reader))
}
