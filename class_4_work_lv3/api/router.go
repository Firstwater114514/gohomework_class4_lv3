package api

import (
	"github.com/gin-gonic/gin"
	"lanshan_homework/go1.19.2/go_homework/class_4_work_lv3/api/middleware"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())
	r.POST("/register", register)
	r.POST("/login", login)
	r.POST("/quit", quit)
	r.POST("/unsubscribe", unsubscribe)
	r.POST("/add friend", addFriend)
	r.POST("/delete friend", deleteFriend)
	r.POST("/scan friends", scanFriends)
	r.POST("/change group", changeGroup)
	r.POST("/search friend", searchFriend)
	r.POST("/scan group", scanGroup)
	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.GET("/get", getUsernameFromToken)
	}
	r.Run(":11451")
}
