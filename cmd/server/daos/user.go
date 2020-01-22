package daos

import (
	"github.com/Kit-Bunrong/sample-rest/cmd/server/config"
	"github.com/Kit-Bunrong/sample-rest/cmd/server/models"
)

// UserDAO persists user data in database
type UserDAO struct{}

// NewUserDAO creates a new UserDAO
func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

// Get does the acutal query to database, if user with specified id is not fund error is returned
func (dao *UserDAO) Get(id uint) (*models.User, error) {
	var user models.User

	// Query Database here...

	// user = models.User{
	// 	Model : models.Model{ID: 1}
	// 	FirstName: "Gary",
	// 	LastName: "Swartz",
	// 	Address: "Not gonna tell you",
	// 	Email: "garyswart@gmail.com"
	// }

	// if using Gorm:
	err := config.Config.DB.Where("id = ?", id).First(&user).Error
	
	return &user, err
}
