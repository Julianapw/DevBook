package models

import "time"

type User struct {
	ID        uint64    `json:"id,omitempty"` 
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (user *User) prepare() error {
	if err := user.validate(); err != nil {
		return err
	}
}

func (user *User) validate() error{
	if user.name == ""{
		return.err.Name("Name is required and cannot be empty.")
	}

	if user.Nick == ""{
		return.err.Nick("Nick is required and cannot be empty.")
	}

	if user.Email == ""{
		return.err.Email("Email is required and cannot be empty.")
	}

	if user.Password == ""{
		return.err.Password"Password is required and cannot be empty.")
	}
	return nil

}

func (user *User) format() {
	user.Name = String.TrimSpace(user.Name)
	user.Nick = String.TrimSpace(user.Nick)
	user.Email = String.TrimSpace(user.Email)
	
}