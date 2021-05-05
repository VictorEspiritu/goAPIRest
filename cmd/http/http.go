package http

import (
   "github.com/gin-gonic/gin"
)

func SetupRouter()  *gin.Engine {
   route:= gin.Default()
   route.GET("/ping", func(c *gin.Context){
      c.JSON(200, gin.H{
         "message": "pong",
      })
   })

   return route
}


