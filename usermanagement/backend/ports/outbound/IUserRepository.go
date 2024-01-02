package outbound

import model "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"

type IUserRepository interface {
	createUser(user model.User)
	updateUser(id int, user model.User)
	deleteUser(id int)
	getUser(id int)
}
