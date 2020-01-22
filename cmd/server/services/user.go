package services

import "github.com/Kit-Bunrong/sample-rest/cmd/server/models"

// Define interface that groups all previously created DAO functions,
// in this case just Get(id uint) from previous section
type userDAO interface {
	Get(id uint) (*models.User, error)
}

// Define UserService which containes our DAO and a function that create
// it by using DAO supplied as parameter.
type UserService struct {
	dao userDAO
}

// NewUserService create a new UserService with the given user DAO
func NewUserService(dao userDAO) *UserService {
	return &UserService{dao}
}

// Get just retrieves user using UserDAO, here can be additional logic for processing data by DAOs
func (s *UserService) Get(id uint) (*models.User, error) {
	return s.dao.Get(id)
}
