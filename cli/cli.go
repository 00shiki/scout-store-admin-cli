package cli

import (
	"fmt"
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
loop:
	for {
		fmt.Println("1. Tambahkan Produk")
		fmt.Println("2. Restock Produk")
		fmt.Println("3. Tambahkan Staff")
		fmt.Println("4. Laporan Penjualan")
		fmt.Println("5. Melakukan Pembelian Produk")
		fmt.Println("6. Keluar")
		fmt.Println()

		fmt.Print("Masukkan pilihan fitur (1/2/3/4/5/6): ")
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
			c.Handler.RestockProduct()
		case 3:
			// TODO: register staff
		case 4:
			// TODO: sales report
		case 5:
			// TODO: product purchase
		case 6:
			fmt.Println("\nTerima Kasih sudah menggunakan aplikasi Toko Aksesoris Pramuka. Sampai Jumpa!")
			break loop
		default:
			fmt.Println("pilihan tidak tersedia mohon untuk pilih (1/2/3/4/5/6)")
		}
	}
}
