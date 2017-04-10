// This file was generated by github.com/nelsam/hel.  Do not
// edit this code by hand unless you *really* know what you're
// doing.  Expect any changes made manually to be overwritten
// the next time hel regenerates this file.

package handlers_test

import "users"

type mockUserService struct {
	RegisterUserCalled chan bool
	RegisterUserInput  struct {
		NewUser chan *users.User
	}
	RegisterUserOutput struct {
		Ret0 chan error
	}
	GetUserCalled chan bool
	GetUserInput  struct {
		EmailAddress chan string
	}
	GetUserOutput struct {
		Ret0 chan *users.User
		Ret1 chan error
	}
}

func newMockUserService() *mockUserService {
	m := &mockUserService{}
	m.RegisterUserCalled = make(chan bool, 100)
	m.RegisterUserInput.NewUser = make(chan *users.User, 100)
	m.RegisterUserOutput.Ret0 = make(chan error, 100)
	m.GetUserCalled = make(chan bool, 100)
	m.GetUserInput.EmailAddress = make(chan string, 100)
	m.GetUserOutput.Ret0 = make(chan *users.User, 100)
	m.GetUserOutput.Ret1 = make(chan error, 100)
	return m
}
func (m *mockUserService) RegisterUser(newUser *users.User) error {
	m.RegisterUserCalled <- true
	m.RegisterUserInput.NewUser <- newUser
	return <-m.RegisterUserOutput.Ret0
}
func (m *mockUserService) GetUser(emailAddress string) (*users.User, error) {
	m.GetUserCalled <- true
	m.GetUserInput.EmailAddress <- emailAddress
	return <-m.GetUserOutput.Ret0, <-m.GetUserOutput.Ret1
}
