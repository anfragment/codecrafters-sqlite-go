// Page related methods and utilities
package page

import (
	"github/com/codecrafters-io/sqlite-starter-go/app/utils"
)

type Page []byte

var (
	InteriorIndex = 2
	InteriorTable = 5
	LeafIndex     = 10
	LeafTable     = 13
)

type PageMetadata struct {
	Type            uint8  // offset 0, one byte
	FreeblockOffset uint16 // offset 1, two bytes
	Cells           uint16 // offset 3, two bytes
	CellOffset      uint16 // offset 5, two bytes
	FreeBytes       uint8  // offset 7, one byte
}

func (p *Page) Metadata() (PageMetadata, error) {
	var metadata PageMetadata
	if err := utils.ReadUint((*p)[0:1], &metadata.Type); err != nil {
		return PageMetadata{}, err
	}
	if err := utils.ReadUint((*p)[1:3], &metadata.FreeblockOffset); err != nil {
		return PageMetadata{}, err
	}
	if err := utils.ReadUint((*p)[3:5], &metadata.Cells); err != nil {
		return PageMetadata{}, err
	}
	if err := utils.ReadUint((*p)[5:7], &metadata.CellOffset); err != nil {
		return PageMetadata{}, err
	}
	if err := utils.ReadUint((*p)[7:8], &metadata.FreeBytes); err != nil {
		return PageMetadata{}, err
	}
	return metadata, nil
}

func (p *Page) CellPointerIndex() ([]uint16, error) {
	metadata, err := p.Metadata()
	if err != nil {
		return nil, err
	}
	index := make([]uint16, 0, metadata.Cells)
	for i := 0; i < int(metadata.Cells); i++ {
		var pointer uint16
		if err := utils.ReadUint((*p)[8+i*2:10+i*2], &pointer); err != nil {
			return nil, err
		}
		index = append(index, pointer)
	}
	return index, nil
}
