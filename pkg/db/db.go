package db

import "os"

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
