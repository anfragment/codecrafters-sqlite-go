package utils

import (
	"bytes"
	"encoding/binary"
)

func ReadUint(b []byte, n interface{}) error {
	if err := binary.Read(bytes.NewReader(b), binary.BigEndian, n); err != nil {
		return err
	}
	return nil
}

const msbMask = byte(uint8(1) << 7)
const removeMSBmask = byte(^(msbMask))

// https://github.com/sqlite/sqlite/blob/master/ext/lsm1/lsm_varint.c#L26
func ReadVarInt(br *bytes.Reader) (uint64, error) {
	var x uint64
	for {
		z, err := br.ReadByte()
		if err != nil {
			return 0, err
		}
		x = x<<6 + uint64(z&removeMSBmask)
		if z&msbMask == byte(0) {
			break
		}
	}
	return x, nil
}
