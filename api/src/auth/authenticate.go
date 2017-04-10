package auth

import "users"

type authenticator struct {
	userManager *user.UserManager
}

func GetAuthenticator(userManager *users.UserManager) *authenticator {
	return &authenicator{
		userManager: userManager,
	}
}

func (a authenticator) Authenticate(emailAddress, password string) (bool, error) {
	user, err := a.userManager.GetUser(emailAddress)
	if err != nil {
		return false, error
	}

	return user.Password == password, nil
}
