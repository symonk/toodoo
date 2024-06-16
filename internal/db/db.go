package db

var client *Client

type Client struct {
}

func Init(connectionString string) {
	client = New(connectionString)

}

func New(connectionString string) *Client {
	return &Client{}

}

func GetDB() *Client {
	return client
}
