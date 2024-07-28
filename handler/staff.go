package handler

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"scout-store-admin-cli/entity"
)

func (h *Handler) RegisterStaff() {
	var name string
	var email string
	var roleID int
	fmt.Print("Masukkan nama staff: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name = scanner.Text()
	fmt.Print("Masukkan email staff: ")
	_, err := fmt.Scan(&email)
	if err != nil {
		log.Fatalf("error scanning email: %v", err)
		return
	}
	fmt.Println("Pilihan Role:")
	fmt.Println("1. Admin")
	fmt.Println("2. Kasir")
	fmt.Print("Masukkan role (1/2): ")
	_, err = fmt.Scan(&roleID)
	if err != nil {
		log.Fatalf("error scanning role: %v", err)
		return
	}
	staff := entity.Staff{
		Name:  name,
		Email: email,
		Role: entity.Role{
			ID: roleID,
		},
	}
	err = h.StaffRepo.AddStaff(staff)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Staff berhasil ditambahkan")
	fmt.Println("Tekan tombol ENTER untuk melanjutkan ke menu...")
	fmt.Scanf("\n")
}
