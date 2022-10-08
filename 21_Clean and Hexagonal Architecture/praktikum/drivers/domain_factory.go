package drivers

import (
	userDomain "praktikum/businesses/users"
	userDB "praktikum/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}
