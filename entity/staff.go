package entity

type Staff struct {
	ID    int
	Name  string
	Email string
	Role
}

type Role struct {
	ID   int
	Name string
}
