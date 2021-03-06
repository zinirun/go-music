package rest

import "github.com/gin-gonic/gin"

func RunAPI(address string) error {
	r := gin.Default()
	h, _ := NewHandler()

	//get products
	r.GET("/products", h.GetProducts)
	//get promos
	r.GET("/promos", h.GetPromos)
	/*
		//post user sign in
		r.POST("/user/signin", h.SignIn)
		//post user sign out
		r.POST("/user/:id/signout", h.SignOut)
		//get user orders
		r.GET("/user/:id/orders", h.GetOrders)
		//post purchase charge
		r.POST("/user/charge", h.Charge)
	*/

	//create group in same route
	userGroup := r.Group("/user")
	{
		userGroup.POST("/:id/signout", h.SignOut)
		userGroup.GET("/:id/orders", h.GetOrders)
	}

	usersGroup := r.Group("/users")
	{
		usersGroup.POST("/charge", h.Charge)
		usersGroup.POST("/signin", h.SignIn)
		usersGroup.POST("", h.AddUser)
	}

	return r.Run(address)
}
