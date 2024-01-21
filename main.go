package main

import "github.com/gin-gonic/gin"



func main() {
	router := gin.Default()
	router.GET("/members", getMember)
	router.POST("/addmember", addMember)
	router.GET("getmemberbyid/:id", getMemberByID)
	router.PUT("/updatememberbyid/:id", updateMemberByID)
	router.GET("/add/:num1/:num2", add)
	router.GET("/div/:number1/:number2" , div)
	router.GET("/check/:userinput" , check)
	router.GET("/pass/:inp" , rej)
	router.Run("localhost:8080")
}
