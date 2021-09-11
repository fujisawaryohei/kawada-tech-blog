package gateways

import (
	"github.com/fujisawaryohei/echo-app/database/dao"
	"github.com/fujisawaryohei/echo-app/web/dto"
	"gorm.io/gorm"
)

type UserRepository struct {
	dbConn *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		dbConn: db,
	}
}

func (repo *UserRepository) FindById(id string) (dao.User, error) {
	var userDAO dao.User
	result := repo.dbConn.First(userDAO, "id = ?", id)
	if err := result.Error; err != nil {
		return userDAO, err
	}
	return userDAO, nil
}

func (repo *UserRepository) SaveUser(user *dto.UserDTO) error {
	var userDAO dao.User
	if err := repo.dbConn.Create(userDAO.ConvertToDAO(user)).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) Delete(id string) error {
	var userDAO dao.User
	if err := repo.dbConn.Delete(userDAO, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
