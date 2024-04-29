package types

import (
	"fmt"
	"time"
)

type User struct {
	Id        int
	Name      string `form:"name"`
	Email     string `form:"email"`
    Birthdate time.Time 
    IsAdmin   bool `form:"isAdmin"`
	CreateAt  time.Time
}

func (u User) Permalink(action string) string {
	return fmt.Sprintf("/users/%d/%s/", u.Id, action)
}

func (u User) Validate() {
}
