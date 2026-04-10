package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DatabaseClient interface {
	Ready() bool
}

type Client struct {
	Db *gorm.DB
}

func NewDatabaseClient()(DatabaseClient, error){
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", "localhost", "root", "Lam1234567890", "gomicro",3306, "disable")
	db , err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy:schema.NamingStrategy{
			TablePrefix:"micro",
		},
		NowFunc: func () time.Time  {
			return time.Now().Truncate()
		},
		QueryFields: true,
	})
	if err != nil {
		return nil, err
	}
	client := Client{
		Db: db,
	}
	return client, nil
}

func (c Client) Ready() bool{
	var ready strring
	tx := c.Db.Raw("SELECT 1 as ready").Scan($ready)
	if tx.error != nil {
		return false
	}
	if ready == "1"{
		return true
	}
	return false
}