package controllers

import (
	"UjianKetrampilan/db"
	"UjianKetrampilan/dto"
	"UjianKetrampilan/models"
	"UjianKetrampilan/pkg/helpers"
	"UjianKetrampilan/pkg/internal_jwt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(ctx *gin.Context) {
	var payload dto.UserCreateDTO

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		helpers.GenerateErrResponse(
			ctx,
			http.StatusUnprocessableEntity,
			"invalid request",
		)
		return
	}

	pg := db.GetDB()
	var user models.User
	if err := pg.Where("username = ?", payload.Username).First(&user).Error; err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username already taken"})
		return
	}
	if err := pg.Where("email = ?", payload.Email).First(&user).Error; err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email already taken"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(payload.Password),
		8,
	)

	if err != nil {
		helpers.GenerateErrResponse(ctx, http.StatusInternalServerError, "something went wrong")
		return
	}

	newUser := models.User{
		Email:    payload.Email,
		Username: payload.Username,
		Age:      payload.Age,
		Password: string(hashedPassword),
	}

	if err := pg.Create(&newUser).Error; err != nil {
		helpers.GenerateErrResponse(ctx, http.StatusInternalServerError, "something went wrong")
		return
	}

	result := dto.UserResponseDTO{
		ID:       newUser.ID,
		Username: newUser.Username,
		Email:    newUser.Email,
		Age:      newUser.Age,
	}

	ctx.JSON(http.StatusCreated, result)
}

func Login(ctx *gin.Context) {
	var payload dto.UserLoginDTO

	if err := ctx.ShouldBindBodyWithJSON(&payload); err != nil {
		helpers.GenerateErrResponse(
			ctx,
			http.StatusUnprocessableEntity,
			"invalid request body",
		)
		return
	}

	pg := db.GetDB()

	user := models.User{}

	if err := pg.First(&user, "email = ?", payload.Email).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.GenerateErrResponse(ctx,
				http.StatusUnauthorized,
				"invalid email/password",
			)
			return
		}

		helpers.GenerateErrResponse(
			ctx,
			http.StatusInternalServerError,
			"something went wrong",
		)
		return
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(payload.Password),
	); err != nil {
		helpers.GenerateErrResponse(ctx,
			http.StatusUnauthorized,
			"invalid email/password",
		)
		return
	}

	claim := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
	}

	result := dto.LoginResponseDTO{
		Status: "success",
		Data:   internal_jwt.GenerateToken(claim),
	}

	ctx.JSON(http.StatusOK, result)
}

func UpdateUser(c *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	userID, _ := c.Get("userID")

	if err := c.ShouldBindJSON(&userUpdateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pg := db.GetDB()
	var user models.User
	if err := pg.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Username = userUpdateDTO.Username
	user.Email = userUpdateDTO.Email

	if err := pg.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.UserResponseDTO{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Age:       user.Age,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func DeleteUser(c *gin.Context) {
	userID, _ := c.Get("userID")

	pg := db.GetDB()
	var user models.User
	if err := pg.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := pg.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
