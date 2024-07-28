package sales

import (
	"fmt"
	"scout-store-admin-cli/entity"
)

func (r *Repository) AddSale(sale entity.Sale) error {
	query := `INSERT INTO Sales (Date, ProductID, Quantity, TotalAmount, StaffID) VALUE (?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, sale.Date, sale.Product.ID, sale.Quantity, sale.TotalAmount, sale.Staff.ID)
	if err != nil {
		return fmt.Errorf("could not insert sale: %v", err)
	}
	return nil
}
