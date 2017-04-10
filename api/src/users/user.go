package users

import (
	"encoding/json"

	redis "gopkg.in/redis.v5"
)

type User struct {
	FirstName    string
	LastName     string
	EmailAddress string
	Password     string
}

type UserManager struct {
	client *redis.Client
}

func NewUserManager(url, password string, db int) *UserManager {
	local := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       db,
	})

	if local == nil {
		return nil
	}

	return &UserManager{
		client: local,
	}
}

func (u *UserManager) RegisterUser(newUser *User) error {
	jsonIn, err := json.Marshal(newUser)
	if err != nil {
		return err
	}

	return u.client.Set(newUser.EmailAddress, string(jsonIn), 0).Err()
}

func (u *UserManager) GetUser(emailAddress string) (*User, error) {
	encString, err := u.client.Get(emailAddress).Result()
	if err != nil {
		return nil, err
	}

	temp := &User{}
	err = json.Unmarshal([]byte(encString), temp)
	if err != nil {
		return nil, err
	}

	return temp, nil
}
