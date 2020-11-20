package routes

import (
	"github.com/gin-gonic/gin"

	serv "nano-world-backend-golang/services"
	"net/http"
)

type headerBind struct {
	ContentType string `header:"Content-Type"`
}

type errorMessage struct {
	s string
}

func (e *errorMessage) Error() string {
	return e.s
}

// GetBalance endpoint
// @Summary fetches all locations by ip
// @Description return a list of locations
// @ID /getBalance/:address
// @Accept  json
// @Produce  json
// @Param   address      path   string     true  "Address of the wallet"
// @Success 200 {string} string	"ok"
// @Failure 415 {string} string "Invalid media type, must be application/json"
// @Failure 404 {string} string "Page not found"
// @Router /getBalance/{address} [get]
func GetBalance() gin.HandlerFunc {

	return func(c *gin.Context) {

		if msg, err := checkHeader(c); err != nil {
			c.JSON(http.StatusUnsupportedMediaType, gin.H{"message": msg.s})
			return
		}

		address := c.Param("address")
		resp, err := serv.GetFunds(address)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"Internal error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"balance": resp.Balance,
			"pending": resp.Pending,
		})
	}

}

// GetPeers endpoint
// @Summary fetches all peers
// @Description return a list of peers
// @ID /getPeers
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Failure 415 {string} string "Invalid media type, must be application/json"
// @Failure 404 {string} string "Page not found"
// @Router /getPeers [get]
func GetPeers() gin.HandlerFunc {

	return func(c *gin.Context) {

		if msg, err := checkHeader(c); err != nil {
			c.JSON(http.StatusUnsupportedMediaType, gin.H{"message": msg.s})
			return
		}

		resp, err := serv.GetPeers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		c.Header("Content-Type", "application/json") // Define header manually
		c.String(http.StatusOK, string(resp))
	}

}

// GetGeoLocations returns all city locations
// @Summary fetches all locations by ip
// @Description return a list of locations
// @ID /getGeoLocations
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Failure 415 {string} string "Invalid media type, must be application/json"
// @Failure 404 {string} string "Page not found"
// @Router /getGeoLocations [get]
func GetGeoLocations() gin.HandlerFunc {

	return func(c *gin.Context) {

		if msg, err := checkHeader(c); err != nil {
			c.JSON(http.StatusUnsupportedMediaType, gin.H{"message": msg.s})
			return
		}

		resp, err := serv.GetGeoLocations()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		c.JSONP(http.StatusOK, gin.H{"cityResponse": resp})
	}

}

func checkHeader(c *gin.Context) (errorMessage, error) {

	header := headerBind{}
	logger := gin.DefaultWriter

	if err := c.ShouldBindHeader(&header); err != nil {

		logger.Write([]byte(err.Error()))
		return errorMessage{}, err
	}

	if header.ContentType != "application/json" {

		msgError := errorMessage{s: "Invalid media type, must be application/json"}
		logger.Write([]byte(msgError.Error()))
		return msgError, &msgError
	}

	return errorMessage{}, nil

}
