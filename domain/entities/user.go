package entities

import "errors"

type User struct {
	ID       int
	Username string
	Income   float64
}

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("o usuário deve ter um nome")
	}

	if u.Income <= 0 {
		return errors.New("a renda do usuário deve ser positiva")
	}

	return nil
}
