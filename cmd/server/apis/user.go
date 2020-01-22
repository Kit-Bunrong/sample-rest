package apis

import (
	"github.com/Kit-Bunrong/sample-rest/cmd/server/daos"
	"github.com/Kit-Bunrong/sample-rest/cmd/server/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetUser godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
// @Security ApiKeyAuth

// This is a function used to serve an API endpoint, by first we create service with supplied user
// DAO. Next we parse ID which we expect to be in URL(something like /users/{id}), then we use
// service to get us user data from database and finally, if the data is found we return in JSON
// format with 200 status code.
func GetUser(c *gin.Context) {
	s := services.NewUserService(daos.NewUserDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if user, err := s.Get(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
