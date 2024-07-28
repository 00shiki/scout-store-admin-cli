package product

import (
	"fmt"
	"scout-store-admin-cli/entity"
)

func (r *Repository) AddProduct(product entity.Product) error {
	query := `INSERT INTO Products (ProductName, Price, Stock) VALUES (?, ?, ?)`
	_, err := r.DB.Query(query, product.Name, product.Price, product.Stock)
	if err != nil {
		return fmt.Errorf("could not insert product: %v", err)
	}
	return nil
}

func (r *Repository) UpdateStockByID(productID, newStock int) error {
	query := `UPDATE Products SET Stock = ? WHERE ProductID = ?`
	_, err := r.DB.Exec(query, newStock, productID)
	if err != nil {
		return fmt.Errorf("error updating product: %v", err)
	}
	return nil
}
