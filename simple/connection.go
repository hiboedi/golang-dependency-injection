package simple

import "fmt"

type Connection struct {
	*File
}

func (c *Connection) Close() {
	fmt.Println("Close connection", c.File.Name)
}

func NewConnection(file *File) (*Connection, func()) {
	connect := &Connection{File: file}
	return connect, func() {
		connect.Close()
	}
}
