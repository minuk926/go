package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	userRouterInit     sync.Once
	userRouterInstance *userRouter
)

type userRouter struct {
	network *Network
}

func NewUserRouter(network *Network) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{network: network}

		network.create("/user", userRouterInstance.create)
		network.get("/user", userRouterInstance.get)
		network.update("/user", userRouterInstance.update)
		network.delete("/user", userRouterInstance.delete)
		//userRouterInstance.network.engin.POST("/user", userRouterInstance.create)
		//userRouterInstance.network.engin.GET("/user", userRouterInstance.get)
		//userRouterInstance.network.engin.PUT("/user", userRouterInstance.update)
		//userRouterInstance.network.engin.DELETE("/user", userRouterInstance.delete)
	})
	return userRouterInstance
}

func (u *userRouter) create(c *gin.Context) {
	fmt.Println("create user")

}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get user")

}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update user")

}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete user")
}
