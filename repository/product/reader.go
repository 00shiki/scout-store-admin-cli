package product

import (
	"fmt"
	"scout-store-admin-cli/entity"
)

func (r *Repository) FetchProducts() ([]entity.Product, error) {
	query := `SELECT ProductID, ProductName, Price, Stock FROM Products`
	var products []entity.Product
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying products: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var p entity.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
		if err != nil {
			return nil, fmt.Errorf("error scanning product: %v", err)
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *Repository) FetchProductByID(id int) (entity.Product, error) {
	query := `SELECT ProductID, ProductName, Price, Stock FROM Products WHERE ProductID = ?`
	var p entity.Product
	rows, err := r.DB.Query(query, id)
	if err != nil {
		return entity.Product{}, fmt.Errorf("error querying product: %v", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return entity.Product{}, fmt.Errorf("product not found")
	}
	if err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock); err != nil {
		return entity.Product{}, fmt.Errorf("error scanning product: %v", err)
	}
	return p, nil
}
