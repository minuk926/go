package network

import (
	"github.com/gin-gonic/gin"
)

type Network struct {
	engin *gin.Engine
}

// mux, echo, gin
func NewNetwork() *Network {
	var r = &Network{
		//engin: gin.Default(),
		engin: gin.New(),
	}

	return r
}

func (n *Network) ServerStart(port string) error {
	return n.engin.Run(port)
}
