package mocks

import "snippetbox.marouane.com/internal/models"

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "user1@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return &models.User{Email: "user1@example.com", Name: "user1"}, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == "user1@example.com" && password == "pa$$word" {
		return 1, nil
	}
	return 0, models.ErrInvalidCredentials
}
func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

func (m *UserModel) PasswordUpdate(id int, currentPassword, newPassword string) error {
	if currentPassword != "password" {
		return models.ErrInvalidCredentials
	}
	return nil
}
