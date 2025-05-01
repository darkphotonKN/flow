package booking

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookingHandler struct {
	Service *BookingService
}

func NewBookingHandler(service *BookingService) *BookingHandler {
	return &BookingHandler{
		Service: service,
	}
}

func (h *BookingHandler) CreateBookingHandler(c *gin.Context) {
	userIdParam := c.Param("user_id")

	// parse and check userId is a valid uuid
	userId, err := uuid.Parse(userIdParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode:": http.StatusBadRequest, "message": fmt.Sprintf("Error with id %d, not a valid uuid.", userId)})
		return
	}

	var booking CreateBookingRequest

	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode:": http.StatusBadRequest, "message": fmt.Sprintf("Error with parsing payload as JSON.")})
		return
	}

	err = h.Service.CreateBookingService(userId, booking)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"statusCode:": http.StatusInternalServerError, "message": fmt.Sprintf("Error when attempting to create booking: %s", err.Error())})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"statusCode:": http.StatusCreated, "message": "Successfully created booking."})
}

func (h *BookingHandler) GetBookingByIdHandler(c *gin.Context) {
	// get id from param
	idParam := c.Param("id")

	// check that its a valid uuid
	id, err := uuid.Parse(idParam)

	// get user_id from query param
	userIdQuery := c.Query("user_id")

	userId, userIdErr := uuid.Parse(userIdQuery)

	if err != nil || userIdErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode:": http.StatusBadRequest, "message": fmt.Sprintf("Error with ids pass in: \n%d\n%d, are not a valid uuids.", id, userId)})
		// return to stop flow of function after error response
		return
	}

	booking, err := h.Service.GetBookingByIdService(userId, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode:": http.StatusBadRequest, "message": fmt.Sprintf("Error when attempting to get user with id %d %s", id, err.Error())})

		return
	}

	c.JSON(http.StatusOK, gin.H{"statusCode:": http.StatusOK, "message": "Successfully retrived booking.",
		// de-reference to return the user struct, not pointer
		"result": *booking})

}
