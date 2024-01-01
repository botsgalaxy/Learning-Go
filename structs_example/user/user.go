package user

import ( 
	"fmt"
	"time"
	"errors"
)

type User struct {
	firstName  string
	lastName   string
	birthdate  string
	created_at time.Time
}

type Admin struct {
	email string
	password string
	User
}

func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)

}

func (u *User) ClearUserName(){ 
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email,password string) Admin { 
	return Admin{email: email,password: password, 
	User: User{ 
		firstName: "admin",
		lastName: "admin",
		birthdate: "___",
		created_at: time.Now(),
	},}
}

func New(firstName,lastName,birthdate string) (*User, error){ 
	if firstName == "" || lastName == "" || birthdate== "" { 
		return nil, errors.New("firstname, lastname and birthdate are required")
	}
	return &User{ 
		firstName: firstName,
		lastName: lastName,
		birthdate: birthdate,
		created_at: time.Now(),
	},nil

}
