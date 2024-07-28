package cli

import (
	"fmt"
	"log"
	"scout-store-admin-cli/handler"
)

type Cli struct {
	Handler *handler.Handler
}

func New(handler *handler.Handler) *Cli {
	return &Cli{
		Handler: handler,
	}
}

func (c *Cli) Init() {
	fmt.Println("Selamat Datang di Toko Aksesoris Pramuka!\n")

	for {
		fmt.Println("1. Tambahkan Produk")
		fmt.Println("2. Restock Produk")
		fmt.Println("3. Tambahkan Staff")
		fmt.Println("4. Laporan Penjualan")
		fmt.Println("5. Melakukan Pembelian Produk")
		fmt.Println()

		log.Print("Masukkan pilihan fitur (1/2/3/4/5): ")
		var option int
		_, err := fmt.Scanf("%d", &option)
		if err != nil {
			fmt.Println("Error:", err)
		}
		switch option {
		case 1:
			c.Handler.AddProduct()
		case 2:
			c.Handler.ShowProducts()
		case 3:
			// TODO: register staff
		case 4:
			// TODO: sales report
		case 5:
			// TODO: product purchase
		default:
			fmt.Println("pilihan tidak tersedia mohon untuk pilih (1/2/3/4/5)")
		}
	}
}
