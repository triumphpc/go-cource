package user

import "errors"

type User struct {
	Name       string
	SecondName string
	Age        int
}

// Set name for user struct
func (u *User) SetName(n string) error {
	if len(n) < 2 {
		return errors.New("invalid name")
	}
	u.Name = n
	return nil
}

// Set Second Name for user struct
func (u *User) SetSecondName(n string) error {
	if len(n) < 2 {
		return errors.New("invalid name")
	}
	u.SecondName = n
	return nil
}

// Set Age for user struct
func (u *User) SetAge(a int) error {
	if a < 1 || a > 110 {
		return errors.New("invalid age")
	}
	u.Age = a
	return nil
}
