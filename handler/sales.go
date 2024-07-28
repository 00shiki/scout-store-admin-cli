package handler

import (
	"fmt"
	"log"
	"scout-store-admin-cli/utils"
)

func (h *Handler) SalesReport() {
	fmt.Print("Masukkan periode penjualan (YYYY-MM-DD): ")
	var startDate, endDate string
	_, err := fmt.Scan(&startDate, &endDate)
	if err != nil {
		log.Fatalf("error scanning date: %v", err)
		return
	}
	if startDate == "" || endDate == "" {
		log.Fatalf("error scanning date: startDate or endDate must not empty")
	}
	if err := utils.ValidateDate(startDate); err != nil {
		log.Fatalf("error validating start date: %v", err)
	}
	if err := utils.ValidateDate(endDate); err != nil {
		log.Fatalf("error validating end date: %v", err)
	}
	sales, err := h.SalesRepo.FetchSales(startDate, endDate)
	if err != nil {
		log.Fatalf("error fetching sales: %v", err)
	}
	utils.SalesTable(sales)
	fmt.Println("\nTekan tombol ENTER untuk melanjutkan ke menu...")
	fmt.Scanf("\n")
}
