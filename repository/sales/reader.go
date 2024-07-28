package sales

import "scout-store-admin-cli/entity"

func (r *Repository) FetchSales(startDate, endDate string) ([]entity.Sale, error) {
	// TODO: add query
	query := `
			SELECT s.SaleID,
				   s.Date,
				   p.ProductName,
				   p.Price,
				   s.Quantity,
				   s.TotalAmount,
				   st.StaffName
			FROM Sales s
					 JOIN Products p
						  ON s.ProductID = p.ProductID
					 JOIN Staff st
						  ON s.StaffID = st.StaffID
			WHERE s.Date BETWEEN ? AND ?`
	rows, err := r.DB.Query(query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var sales []entity.Sale
	for rows.Next() {
		var sale entity.Sale
		rows.Scan(
			&sale.ID,
			&sale.Date,
			&sale.Product.Name,
			&sale.Product.Price,
			&sale.Quantity,
			&sale.TotalAmount,
			&sale.Staff.Name,
		)
		sales = append(sales, sale)
	}
	return sales, nil
}
