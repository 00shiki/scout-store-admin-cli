package handler

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"scout-store-admin-cli/entity"
	"scout-store-admin-cli/utils"
)

func (h *Handler) AddProduct() {
	var name string
	var price float64
	var stock int
	fmt.Print("Masukkan nama produk: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name = scanner.Text()
	fmt.Print("Masukkan harga produk: ")
	_, err := fmt.Scan(&price)
	if err != nil {
		log.Fatalf("error scanning price: %v", err)
		return
	}
	fmt.Print("Masukkan stok awal produk: ")
	_, err = fmt.Scan(&stock)
	if err != nil {
		log.Fatalf("error scanning stock: %v", err)
		return
	}
	product := entity.Product{
		Name:  name,
		Price: price,
		Stock: stock,
	}
	err = h.ProductRepo.AddProduct(product)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Produk berhasil ditambahkan")
	fmt.Println("Tekan tombol ENTER untuk melanjutkan ke menu...")
	fmt.Scanf("\n")
}

func (h *Handler) ShowProducts() {
	products, err := h.ProductRepo.FetchProducts()
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, product := range products {
		fmt.Printf("%d. %s Rp %s (%d)\n", product.ID, product.Name, utils.PriceFormat(product.Price), product.Stock)
	}
}
