package util

import (
	"bytes"
	"encoding/binary"
	"google.golang.org/protobuf/proto"
	"os"
	"path/filepath"
	"reflect"
)

func GetTableData(gameDataDir string, gameServer string, payload proto.Message) error {
	tableType := reflect.TypeOf(payload)
	tablePath := filepath.Join(gameDataDir, "Table", gameServer, tableType.Elem().Name()+".bytes")

	encryptedTableData, err := os.ReadFile(tablePath)
	if err != nil {
		return err
	}

	var head uint32
	err = binary.Read(bytes.NewReader(encryptedTableData[:4]), binary.LittleEndian, &head)
	if err != nil {
		return err
	}

	tableData := encryptedTableData[head+4:]

	err = proto.Unmarshal(tableData, payload)
	if err != nil {
		return err
	}

	return nil
}
