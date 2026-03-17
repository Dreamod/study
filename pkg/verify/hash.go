package verify

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"go/study/configs"
	"go/study/pkg/db"
	"log"
	"os"
)

type DataDb struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

func MakeHash(s string) string {
	data := []byte(s)
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

func SaveHash(email string, hash string) error {
	conf := configs.LoadConfig()
	db := db.NewJsonDb(conf.Db.File)
	data, err := json.Marshal(DataDb{
		Email: email,
		Hash:  hash,
	})
	if err != nil {
		return err
	}
	_, err = db.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func VerifyHash(hash string) bool {
	dbData, err := getDbData()
	if err != nil {
		log.Fatal(err)
	}
	return dbData.Hash == hash
}

func getDbData() (DataDb, error) {
	conf := configs.LoadConfig()
	db := db.NewJsonDb(conf.Db.File)
	data, err := db.Read()
	var dbData DataDb
	if err != nil {
		return dbData, err
	}
	err = json.Unmarshal([]byte(data), &dbData)
	if err != nil {
		return dbData, err
	}
	return dbData, nil
}

func RemoveDbData() {
	conf := configs.LoadConfig()
	err := os.Remove(conf.Db.File)
	if err != nil {
		log.Fatal(err)
	}
}
