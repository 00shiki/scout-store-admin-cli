package entity

type Staff struct {
	ID   int
	Name string
	Role
}

type Role struct {
	ID   int
	Name string
}
