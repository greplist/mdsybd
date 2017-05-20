package models

import "database/sql"

// Client - main database app client
type Client struct {
	oracle *sql.DB
}

// NewClient - constructor for Client struct
func NewClient() (client *Client, err error) {
	oracle, err := sql.Open("ora", "anton/anton@127.0.0.1:1521/")
	if err != nil {
		return nil, err
	}
	return &Client{
		oracle: oracle,
	}, nil
}

// Close db connections
func (c *Client) Close() error {
	return c.oracle.Close()
}
