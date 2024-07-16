package controllers

import (
	"UjianKetrampilan/db"
	"UjianKetrampilan/dto"
	"UjianKetrampilan/models"
	"UjianKetrampilan/pkg/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePhoto(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	userID := int(userId.(uint))

	if !ok {
		helpers.GenerateErrResponse(
			ctx,
			http.StatusInternalServerError,
			"something went wrong",
		)
		return
	}

	var payload dto.PhotoCreateDTO

	if err := ctx.ShouldBindBodyWithJSON(&payload); err != nil {
		helpers.GenerateErrResponse(
			ctx,
			http.StatusUnprocessableEntity,
			"invalid request body",
		)
		return
	}

	newPhoto := models.Photo{
		Title:    payload.Title,
		Caption:  payload.Caption,
		PhotoURL: payload.PhotoURL,
		UserID:   uint(userID),
	}

	pg := db.GetDB()

	if err := pg.Create(&newPhoto).Error; err != nil {
		helpers.GenerateErrResponse(
			ctx,
			http.StatusInternalServerError,
			"something went wrong",
		)
		return
	}

	result := dto.PhotoResponseCreateDTO{
		ID:        newPhoto.ID,
		Title:     newPhoto.Title,
		PhotoURL:  newPhoto.PhotoURL,
		Caption:   newPhoto.Caption,
		UserID:    newPhoto.UserID,
		CreatedAt: newPhoto.CreatedAt,
	}

	ctx.JSON(http.StatusCreated, result)
}

func UpdatePhoto(ctx *gin.Context) {
	userID, ok := ctx.Get("userId")
	userId := int(userID.(uint))
	if !ok {
		helpers.GenerateErrResponse(
			ctx,
			http.StatusInternalServerError,
			"something went wrong",
		)
		return
	}

	var payload dto.PhotoUpdateDTO

	if err := ctx.ShouldBindBodyWithJSON(&payload); err != nil {
		helpers.GenerateErrResponse(
			ctx,
			http.StatusUnprocessableEntity,
			"invalid request body",
		)
		return
	}

	photoIdStr := ctx.Param("photoID")
	photoId, err := strconv.Atoi(photoIdStr)
	if err != nil {
		helpers.GenerateErrResponse(ctx, http.StatusBadRequest, "invalid photo ID")
		return
	}
	var photo models.Photo
	pg := db.GetDB()
	if err := pg.Where("id = ? AND user_id = ?", photoId, uint(userId)).First(&photo).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	photo.Title = payload.Title
	photo.Caption = payload.Caption
	photo.PhotoURL = payload.PhotoURL

	if err := pg.Save(&photo).Error; err != nil {
		helpers.GenerateErrResponse(ctx, http.StatusInternalServerError, "failed to update photo")
		return
	}
	result := dto.PhotoResponseUpdateDTO{
		ID:        photo.ID,
		Title:     photo.Title,
		PhotoURL:  photo.PhotoURL,
		Caption:   photo.Caption,
		UserID:    photo.UserID,
		UpdatedAt: photo.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, result)

}

func GetPhotos(c *gin.Context) {
	var photos []models.Photo

	pg := db.GetDB()
	if err := pg.Find(&photos).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photos not found"})
		return
	}

	var result []dto.PhotoResponseDTO
	for _, photo := range photos {
		result = append(result, dto.PhotoResponseDTO{
			ID:        photo.ID,
			Title:     photo.Title,
			PhotoURL:  photo.PhotoURL,
			Caption:   photo.Caption,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User:      dto.UserPhotoDTO{Email: photo.User.Email, Username: photo.User.Username},
		})
	}

	c.JSON(http.StatusOK, result)
}

func DeletePhoto(ctx *gin.Context) {

	userID, ok := ctx.Get("userId")
	userId := int(userID.(uint))
	if !ok {
		helpers.GenerateErrResponse(
			ctx,
			http.StatusInternalServerError,
			"something went wrong",
		)
		return
	}

	pg := db.GetDB()
	id := ctx.Param("photoID")
	if err := pg.Where("id = ? AND user_id = ?", id, userId).Delete(&models.Photo{}).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Photo deleted"})
}
