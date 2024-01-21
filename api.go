package main

import (
	"net/http"
	"regexp"
	"strconv"
	"github.com/gin-gonic/gin"
)

func getMember(c *gin.Context) {
	c.IndentedJSON(200, Members)
}


func addMember(c *gin.Context) {
	var newMember member
	if err := c.BindJSON(&newMember); err != nil {
		return
	}
	Members = append(Members, newMember)
	c.IndentedJSON(http.StatusOK, newMember)
}


func getMemberByID(c *gin.Context) {
	id := c.Param("id")

	for _, m := range Members {
		if m.ID == id {
			c.IndentedJSON(http.StatusOK, m)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "member not found"})
}


func updateMemberByID(c *gin.Context) {
	id := c.Param("id")
	var updatemember member

	if err := c.BindJSON(&updatemember); err != nil {
		return
	}

	for i, m := range Members {
		if m.ID == id {
			Members[i] = updatemember
			c.IndentedJSON(200, Members[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "member not found"})
}


func add(c *gin.Context) {
	num1 := c.Param("num1")
	num2 := c.Param("num2")

	intNum1 , err := strconv.Atoi(num1)
	if err!= nil{
		c.JSON(400, gin.H{"msg" : "Invalid value for num1"})
		return
	}

	intNum2 , err := strconv.Atoi(num2)
	if err!= nil{
		c.JSON(400, gin.H{"msg" : "Invalid value for num2"})
		return
	}

	result := intNum1 + intNum2
	c.JSON(200, gin.H{"result" : result})
}


func div(c *gin.Context){
	number1:= c.Param("number1")
	number2:= c.Param("number2")

	intnumber1 , err := strconv.Atoi(number1)
	if err != nil{
		c.JSON(400, gin.H{"msg" : "Invalid value for num1"})
		return
	}

	intnumber2, err := strconv.Atoi(number2)
	if err != nil{
		c.JSON(400, gin.H{"msg" : "Invalid value for num1"})
		return
	}

	result2:= intnumber1 / intnumber2
	c.JSON(200, gin.H{"result" : result2})
}


func check(c *gin.Context){
	userinput:= c.Param("userinput")

	if userinput == "dodo"{
		c.JSON(200, gin.H{"msg" : "u r right", "sucess" : "true"})
	}else if userinput == "omer"{
		c.JSON(200, gin.H{"msg" : "u r not right" , "sucess" : "false"})
	}else{
		c.JSON(200, gin.H{"msg" : "try again" , "sucess" : "false"})
	}
}


func passwordvalidation(password string) bool {
	passwordPattern := regexp.MustCompile(`^[A-Za-z\d@$!%*?&]*[a-z][A-Za-z\d@$!%*?&]*[A-Z][A-Za-z\d@$!%*?&]*\d[A-Za-z\d@$!%*?&]*[@$!%*?&][A-Za-z\d@$!%*?&]*$`)
	return passwordPattern.MatchString(password)
}


func rej(c *gin.Context){
	inp := c.Param("inp")
	isvalid:= passwordvalidation(inp)
	c.JSON(200 , gin.H{"password" : inp , "isvalid" : isvalid})
}