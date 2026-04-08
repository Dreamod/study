package db

import (
	"go/study/configs"
	"os"
)
import "gorm.io/gorm"
import "gorm.io/driver/postgres"

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(conf.Db.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Db{db}
}

type JsonDb struct {
	fileName string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		fileName: name,
	}
}

// читает файл
func (db *JsonDb) Read() ([]byte, error) {
	data, error := os.ReadFile(db.fileName)
	if error != nil {
		return nil, error
	}
	return data, nil
}

// пишет файл
func (db *JsonDb) Write(content []byte) (int, error) {
	file, error := os.Create(db.fileName)
	if error != nil {
		return 0, error
	}
	defer file.Close()
	return file.Write(content)
}
