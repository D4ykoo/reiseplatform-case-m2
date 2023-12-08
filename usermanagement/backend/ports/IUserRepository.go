package ports

import "model"

type IUserRepository interface {
	createUser(user model.User)
	updateUser(id int, user model.User)
	deleteUser(id int)
	getUser(id int)
}
