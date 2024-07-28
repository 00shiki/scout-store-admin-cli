package utils

import (
	"github.com/jedib0t/go-pretty/table"
	"os"
	"scout-store-admin-cli/entity"
)

func SalesTable(sales []entity.Sale) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(
		table.Row{
			"#",
			"Date",
			"Product Name",
			"Price",
			"Quantity",
			"Total Amount",
			"Staff",
		},
	)
	for i, sale := range sales {
		t.AppendRow(
			table.Row{
				i + 1,
				sale.Date,
				sale.Product.Name,
				"Rp " + PriceFormat(sale.Product.Price),
				sale.Quantity,
				"Rp " + PriceFormat(sale.TotalAmount),
				sale.Staff.Name,
			},
		)
	}
	t.Render()
}
