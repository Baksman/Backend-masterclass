package model

type User struct {
	name     string
	email    string
	password string
}

func (user *User) createUser() error {
	// controller.CreateUser(
	// 	user.name,
	// 	user.email,
	// 	user.password,
	// )

	return nil

}
