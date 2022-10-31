package dao

import (
	"cs5234/dao"
	"path/filepath"
	"testing"
)

func TestReplaceNullValues(t *testing.T) {
	err := filepath.Walk("D:\\GitHub\\CS5424_YCQL\\project_files\\data_files", dao.Visit)
	if err != nil {
		panic(err)
	}
}
