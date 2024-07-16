package middleware

import (
	"UjianKetrampilan/db"
	"UjianKetrampilan/models"
	"UjianKetrampilan/pkg/helpers"
	"UjianKetrampilan/pkg/internal_jwt"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Authentication(ctx *gin.Context) {
	jwtToken := ctx.Request.Header.Get("Authorization")

	claim, err := internal_jwt.ValidateToken(jwtToken)

	if err != nil {
		helpers.GenerateErrResponse(
			ctx,
			http.StatusUnauthorized,
			"unauthorized",
		)
		return
	}
	userID := uint(claim["id"].(float64))
	ctx.Set("userId", userID)
	if userID, ok := claim["id"].(float64); ok {
		fmt.Printf("UserID from contextatttttttt: %d\n", uint(userID))
	} else {
		fmt.Println("Failed to retrieve userID from claim")
	}
	ctx.Next()
}

func Authorization(ctx *gin.Context) {
	userID, ok := ctx.MustGet("userId").(float64)
	fmt.Printf("UserID from contextattttt2: %f\n", userID)
	if !ok {
		helpers.GenerateErrResponse(
			ctx,
			http.StatusForbidden,
			"forbidden access1",
		)
		return
	}
	// Retrieve and convert photoID from URL parameters
	photoIDStr := ctx.Param("photoID")
	photoID, err := strconv.Atoi(photoIDStr)
	if err != nil {
		helpers.GenerateErrResponse(
			ctx,
			http.StatusBadRequest,
			"ID has to be a number",
		)
		return
	}

	var photo models.Photo

	// Initialize the database connection
	pg := db.GetDB()

	// Retrieve the photo by ID
	err = pg.Debug().First(&photo, photoID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.GenerateErrResponse(
				ctx,
				http.StatusNotFound,
				"Photo not found",
			)
		} else {
			helpers.GenerateErrResponse(
				ctx,
				http.StatusInternalServerError,
				"Something went wrong",
			)
		}
		return
	}

	// Check if the user has permission to access the photo
	if uint(userID) != photo.UserID {
		helpers.GenerateErrResponse(
			ctx,
			http.StatusForbidden,
			"forbidden access2",
		)
		return
	}

	// Proceed to the next middleware or handler
	ctx.Next()
}
