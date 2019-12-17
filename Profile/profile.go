package profile

import "fmt"

type User struct {
	Firstname string
	Lastname  string
	Address   string
}

func (u *User) GetProfile() {
	//u.Firstname = "Agung"
	//u.Lastname = "Wicaksono"
	//u.Address = "Jakarta"

	fmt.Println(*u)
}

func (u *User) SetProfile(fname string, lname string, addname string) {
	u.Firstname = fname
	u.Lastname = lname
	u.Address = addname

	fmt.Println(*u)
}
