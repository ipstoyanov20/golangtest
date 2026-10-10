package routes

import (
	"net/http"

	"example.com/restapi/models"
	"example.com/restapi/utils"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the data"})
		return

	}
	err = user.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save the user"})
		return

	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created sucessfully"})
}

func login(ctx *gin.Context){
	var user models.User
	err:= ctx.ShouldBindJSON(&user)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return 

	}
	err = user.ValideCredentials()

	if err != nil{
		ctx.JSON(http.StatusUnauthorized, gin.H{"message":"Could not authentiate"})
		return
	}

	token, err:=utils.GenerateToken(user.Email, user.ID)


	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"message":err.Error()})
		return
	}




	ctx.JSON(http.StatusOK, gin.H{"message":"Login successfull", "token": token})

}
