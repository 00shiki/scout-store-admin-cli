package handler

import (
	"fmt"
	"log"
	"scout-store-admin-cli/entity"
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
		return
	}
	if err := utils.ValidateDate(startDate); err != nil {
		log.Fatalf("error validating start date: %v", err)
		return
	}
	if err := utils.ValidateDate(endDate); err != nil {
		log.Fatalf("error validating end date: %v", err)
		return
	}
	sales, err := h.SalesRepo.FetchSales(startDate, endDate)
	if err != nil {
		log.Fatalf("error fetching sales: %v", err)
		return
	}
	utils.SalesTable(sales)
	fmt.Println("\nTekan tombol ENTER untuk melanjutkan ke menu...")
	fmt.Scanf("\n")
}

func (h *Handler) PurchaseProduct() {
	fmt.Print("Masukkan id produk yang ingin dibeli: ")
	var productID int
	_, err := fmt.Scan(&productID)
	if err != nil {
		log.Fatal("error input product id:", err)
		return
	}
	product, err := h.ProductRepo.FetchProductByID(productID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Print("Masukkan kuantitas pembelian: ")
	var quantity int
	_, err = fmt.Scan(&quantity)
	if err != nil {
		log.Fatal("error input quantity:", err)
		return
	}
	if quantity > product.Stock {
		log.Fatal("quantity must be smaller then stock")
		return
	}
	h.ShowStaff()
	fmt.Print("Masukkan id staff yang bertugas: ")
	var staffID int
	_, err = fmt.Scan(&staffID)
	if err != nil {
		log.Fatal("error input staff id:", err)
	}
	staff, err := h.StaffRepo.FetchStaffByID(staffID)
	if err != nil {
		log.Fatal(err)
		return
	}
	sale := entity.Sale{
		Date:        utils.GetCurrentDate(),
		Quantity:    quantity,
		TotalAmount: float64(quantity) * product.Price,
		Product:     product,
		Staff:       staff,
	}
	err = h.SalesRepo.AddSale(sale)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = h.ProductRepo.UpdateStockByID(productID, product.Stock-quantity)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Pembelian berhasil dibuat")
	fmt.Println("Tekan tombol ENTER untuk melanjutkan ke menu...")
	fmt.Scanf("\n")
}
