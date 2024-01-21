package domain

import model "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"

type IUserHttpService interface {
	RegisterUser(user model.User) (*uint, error)
	LoginUser(username string, password string) (*uint, error)
	ResetPassword(username string, newPassword string) (*uint, error)
}
