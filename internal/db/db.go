package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Queryer interface {
}

var _ Queryer = (*Client)(nil)

var client *Client

type Client struct {
	*sqlx.DB
}

func Init(connectionString string) {
	db := sqlx.MustConnect("postgres", connectionString)
	client = &Client{DB: db}
}

func GetDB() *Client {
	return client
}
