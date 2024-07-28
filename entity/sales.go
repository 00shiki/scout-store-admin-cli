package entity

type Sale struct {
	ID          int
	Date        string
	Quantity    int
	TotalAmount float64
	Product
	Staff
}
