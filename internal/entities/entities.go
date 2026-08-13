package entities

type User struct {
	Name           string
	Age            int
	CurrentBalance float64
}

func NewUser(name string, age int, currentBalance float64) *User {
	return &User{Name: name, Age: age, CurrentBalance: currentBalance}
}
