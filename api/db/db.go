package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq" //import postgres
)

//BaseDbEntity ...
type BaseDbEntity struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

//IContext interface for mocking context
type IContext interface {
	GetDB() *sql.DB
	Close()
}

//Context struct that implements IContext
type Context struct {
	db *sql.DB
}

//GetDBContext inits
func GetDBContext() IContext {
	var c Context
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PW"), os.Getenv("DB_NAME"))
	c.initConnection(connectionString)
	return c
}

//GetDB ...
func (c Context) GetDB() *sql.DB {
	return c.db
}

//Close closes db connection
func (c Context) Close() {
	c.db.Close()
}
func (Context) connectDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Successfully connected to database!")
	return db, nil
}

func (c *Context) initConnection(dbName string) {

	var err error
	c.db, err = c.connectDB(dbName)

	if err != nil || c.db == nil {
		log.Fatal(err)
		panic("db connection cannot be established")
	}

}
