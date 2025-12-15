package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	// In the future, you will add your Database Service here.
	// DB *database.Service
}

// NewUserHandler initializes the handler
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// GetProfile responds with the authenticated user's details
func (h *UserHandler) GetProfile(c echo.Context) error {
	// 1. Retrieve the User ID from the context
	// We set this in the middleware using c.Set("user_id", ...)
	userID, ok := c.Get("user_id").(string)
	
	// Safety check: This should never happen if the middleware is running,
	// but it's good defensive programming.
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve user ID",
		})
	}

	// 2. (Future) Fetch user details from your database using userID
	// user := h.DB.FindUser(userID)

	// 3. Return the JSON response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":      userID,
		"message": "Successfully fetched profile data",
		"role":    "admin", // Dummy data
	})
}