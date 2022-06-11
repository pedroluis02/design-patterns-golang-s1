package singleton

import "fmt"

const (
	Connecting string = "CONNECTING"
	Success    string = "SUCCESS"
	Failed     string = "FAILED"
	Closed     string = "CLOSED"
)

var connection *database_connection

func GetDatabaseConnection() *database_connection {
	if connection == nil {
		credentials := &database_credentials{
			host:     "localhost",
			database: "db",
			username: "root",
			password: "123456",
		}

		connection = &database_connection{
			status: Closed,
		}
		connection.Connect(credentials)
	}

	return connection
}

type database_connection struct {
	status string
}

func (c *database_connection) Connect(credentials *database_credentials) {
	c.status = Success
	fmt.Println("DB connection: OK")
}

func (c *database_connection) Disconnect() {
	c.status = Closed
	fmt.Println("DB connection: closed")
}

func (c *database_connection) GetStatus() string {
	return c.status
}

func (c *database_connection) Execute(query string) {
	fmt.Println(fmt.Sprintf("SQL: %s", query))
}
