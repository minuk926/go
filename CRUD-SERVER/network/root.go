package network

import (
	"github.com/gin-gonic/gin"
)

type Network struct {
	engin *gin.Engine
}

// mux, echo, gin
func NewNetwork() *Network {
	var n = &Network{
		//engin: gin.Default(),
		engin: gin.New(),
	}

	NewUserRouter(n)

	return n
}

func (n *Network) ServerStart(port string) error {
	return n.engin.Run(port)
}

func (n *Network) create(path string, h ...gin.HandlerFunc) {
	n.engin.POST(path, h...)
}

func (n *Network) get(path string, h ...gin.HandlerFunc) {
	n.engin.GET(path, h...)
}

func (n *Network) update(path string, h ...gin.HandlerFunc) {
	n.engin.PUT(path, h...)
}

func (n *Network) delete(path string, h ...gin.HandlerFunc) {
	n.engin.DELETE(path, h...)
}
