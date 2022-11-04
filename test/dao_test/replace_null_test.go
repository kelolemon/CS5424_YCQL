package dao

import (
	"cs5234/dao"
	"path/filepath"
	"testing"
)

func TestReplaceNullValues(t *testing.T) {
	err := filepath.Walk("/Users/yiyangliu/Desktop/NUS Master/CS5424/CS5424_YCQL/data_files", dao.ReplaceNullCarrier)
	if err != nil {
		panic(err)
	}
}
