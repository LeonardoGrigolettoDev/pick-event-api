package utils

import (
	"github.com/gin-gonic/gin"
)

// UtilsInterface defines common utility functions
type UtilsInterface interface {
	EncodeImageToBase64(imageBytes []byte) (string, error)
	RespondWithError(c *gin.Context, code int, message string)
}

// Utils implements UtilsInterface
type Utils struct{}

// NewUtils creates a new Utils instance
func NewUtils() UtilsInterface {
	return &Utils{}
}

// EncodeImageToBase64 converts image bytes to base64 string
func (u *Utils) EncodeImageToBase64(imageBytes []byte) (string, error) {
	// Keep your existing implementation here
	return EncodeImageToBase64(imageBytes)
}

// RespondWithError sends an error response
func (u *Utils) RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
}
