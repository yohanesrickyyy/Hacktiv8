package cli

import (
	"PairProjectPhase1/config"
	"PairProjectPhase1/entity"
	"PairProjectPhase1/handler"
	"bufio"
	"fmt"
	"net/mail"
	"os"
	"regexp"
	"strconv"
	"time"
)

func RunProgram() {
	scanner := bufio.NewScanner(os.Stdin)
	err := config.ConnectDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer config.DB.Close()
	fmt.Println("WELCOME TO SY STORE!")
	for {
		fmt.Println("1. Login\n2. Register\n3. Exit")
		input := userInput("Masukkan input berdasarkan angka di menu (1/2/3): ", scanner)
		switch input {
		case "1":
			fmt.Println("LOGIN")
			email := userInput("Masukkan Email Anda: ", scanner)
			password := userInput("Masukkan Password Anda: ", scanner)
			fmt.Println("Mohon menunggu sesaat")
			user, valid := handler.Login(email, password)
			if valid {
				fmt.Println("Login berhasil! Selamat datang", user.Name)
				if user.Role == "Admin" {
					for {
						fmt.Println("1. Menampilkan user report\n2. Menampilkan order report\n3. Menampilkan stock report\n4. Menambahkan barang\n5. Update stock barang\n6. Exit")
						inputAdmin := userInput("Masukkan input berdasarkan angka di menu (1-7): ", scanner)
						switch inputAdmin {
						case "1":
							handler.ReportTotal()
							scanEnter()
						case "2":
							handler.OrderReport()
							scanEnter()
						case "3":
							handler.StockReport()
							scanEnter()
						case "4":
							fmt.Println("MENAMBAHKAN PRODUCT!")
							inputProduct := userInput("Masukkan nama product yang ingin ditambahkan: ", scanner)
							fmt.Println("1. S\n2. M\n3. L\n4. XL")
							inputSize := userInput("Masukkan Size (1/2/3/4): ", scanner)
							inputSizeInt, err := strconv.Atoi(inputSize)
							if err != nil {
								fmt.Println("Input tidak valid!")
								continue
							}
							inputPrice := userInput("Masukkan stock: ", scanner)
							inputPriceFloat, err := strconv.ParseFloat(inputPrice, 64)
							if err != nil {
								fmt.Println("Input tidak valid!")
								continue
							}
							inputStock := userInput("Masukkan stock: ", scanner)
							inputStockInt, err := strconv.Atoi(inputStock)
							if err != nil {
								fmt.Println("Input tidak valid!")
								continue
							}
							product := entity.Product{
								ProductName: inputProduct,
								SizeName:    inputSizeInt,
								Price:       inputPriceFloat,
								Stock:       inputStockInt,
							}
							err = handler.AddProduct(product)
							if err != nil {
								fmt.Println(err)
							} else {
								fmt.Println("Barang berhasil ditambahkan")
							}
						case "5":
							handler.ShowProduct()
							inputProduct := userInput("Pilih product berdasarkan angka di menu: ", scanner)
							inputProductInt, err := strconv.Atoi(inputProduct)
							if err != nil {
								fmt.Println("Input tidak valid!")
								continue
							}
							inputStock := userInput("Masukkan quantity: ", scanner)
							inputStockInt, err := strconv.Atoi(inputStock)
							if err != nil {
								fmt.Println("Input tidak valid!")
								continue
							}
							product := entity.Product{
								ProductDetailId: inputProductInt,
								Stock:           inputStockInt,
							}
							err = handler.UpdateStock(product)
							if err != nil {
								fmt.Println(err)
							} else {
								fmt.Println("Berhasil menambahkan stock")
								time.Sleep(1 * time.Second)
							}
						case "6":
							return
						default:
							fmt.Println("Input yang dimasukkan tidak valid")
						}
					}
				} else {
					for {
						fmt.Println("1. Menampilkan Barang\n2. Menambahkan barang ke cart\n3. Menghapus barang dari cart\n4. Menambahkan Wishlist\n5. Menghapus Wishlist\n6. Menampilkan Wishlist\n7. History Order\n8. TopUp saldo\n9. Pembayaran\n10. Exit")
						inputCustomer := userInput("Masukkan input berdasarkan angka di menu (1-10): ", scanner)
						switch inputCustomer {
						case "1":
							handler.ShowProduct()
							scanEnter()
						case "2":
							handler.ShowProduct()
							inputCart := userInput("Masukkan ke Cart berdasarkan angka di menu: ", scanner)
							inputCartInt, err := strconv.Atoi(inputCart)
							if err != nil {
								fmt.Println("Input tidak valid!")
								continue
							}
							inputQuantity := userInput("Masukkan quantity: ", scanner)
							inputQuantityInt, err := strconv.Atoi(inputQuantity)
							if err != nil {
								fmt.Println("Input tidak valid!")
								continue
							}
							product := entity.Product{
								ProductDetailId: inputCartInt,
							}
							quantity := entity.Cart{
								Quantity: inputQuantityInt,
							}
							err = handler.AddToCart(user, product, quantity)
							if err != nil {
								fmt.Println(err)
							} else {
								fmt.Println("Berhasil Add to Cart")
								time.Sleep(1 * time.Second)
							}
						case "3":
							handler.ShowCart(user)
							inputCart := userInput("Remove dari Cart berdasarkan angka di menu: ", scanner)
							inputCartInt, err := strconv.Atoi(inputCart)
							if err != nil {
								fmt.Println("Input tidak valid!")
								continue
							}
							product := entity.Product{
								ProductDetailId: inputCartInt,
							}
							err = handler.RemoveCart(user, product)
							if err != nil {
								fmt.Println(err)
							} else {
								fmt.Println("Berhasil Remove dari Cart")
								time.Sleep(1 * time.Second)
							}
						case "4":
							handler.ShowProduct()
							inputWishlist := userInput("Masukkan Wishlist berdasarkan angka di menu: ", scanner)
							inputWishlistInt, err := strconv.Atoi(inputWishlist)
							if err != nil {
								fmt.Println("Input tidak valid!")
								continue
							}
							wishlist := entity.Product{
								ProductDetailId: inputWishlistInt,
							}
							err = handler.AddWishlist(user, wishlist)
							if err != nil {
								fmt.Println(err)
							} else {
								fmt.Println("Wishlist berhasil ditambahkan")
								time.Sleep(1 * time.Second)
							}
						case "5":
							handler.ShowWishlist(user)
							inputWishlist := userInput("Remove dari Cart berdasarkan angka di menu: ", scanner)
							inputWishlistInt, err := strconv.Atoi(inputWishlist)
							if err != nil {
								fmt.Println("Input tidak valid!")
								continue
							}
							product := entity.Product{
								ProductDetailId: inputWishlistInt,
							}
							err = handler.RemoveWishlist(user, product)
							if err != nil {
								fmt.Println(err)
							} else {
								fmt.Println("Berhasil Remove dari Wishlist")
								time.Sleep(1 * time.Second)
							}
						case "6":
							handler.ShowWishlist(user)
							scanEnter()
						case "7":
							handler.HistoryOrder(user)
							scanEnter()
						case "8":
							inputTopUp := userInput("Masukkan nominal untuk TopUp: Rp", scanner)
							inputTopUpInt, err := strconv.ParseFloat(inputTopUp, 64)
							if err != nil {
								fmt.Println("Input tidak valid!")
								continue
							}
							user.Balance = inputTopUpInt
							err = handler.AddBalance(user)
							if err != nil {
								fmt.Println(err)
							} else {
								fmt.Println("TopUp telah berhasil")
								time.Sleep(1 * time.Second)
							}
						case "9":
							var cart entity.Cart
							var product entity.Product
							var order entity.Order
							err = handler.Checkout(user, product, cart, order)
							if err != nil {
								fmt.Println(err)
							}

						case "10":
							return
						default:
							fmt.Println("Input yang dimasukkan tidak valid")
						}
					}
				}
			} else {
				fmt.Println("Login Gagal!")
			}
		case "2":
			fmt.Println("REGISTER")
			email := userInput("Masukkan Email: ", scanner)
			if len(email) < 1 {
				fmt.Println("Input tidak boleh kosong!")
				continue
			}
			_, err = mail.ParseAddress(email)
			if err != nil {
				fmt.Println("Format Email tidak valid!")
				continue
			}
			password := userInput("Masukkan Password: ", scanner)
			if len(password) < 1 {
				fmt.Println("Input tidak boleh kosong!")
				continue
			}
			name := userInput("Masukkan Nama: ", scanner)
			if len(name) < 1 {
				fmt.Println("Input tidak boleh kosong!")
				continue
			}
			address := userInput("Masukkan Alamat: ", scanner)
			if len(address) < 1 {
				fmt.Println("Input tidak boleh kosong!")
				continue
			}
			phoneNumber := userInput("Masukkan Nomor HP: ", scanner)
			if len(phoneNumber) < 1 {
				fmt.Println("Input tidak boleh kosong!")
			}
			regexPattern := `^(0|\+62)\d{9,12}$`
			re := regexp.MustCompile(regexPattern)
			phoneNumberValid := re.MatchString(phoneNumber)
			if !phoneNumberValid {
				fmt.Println("Nomor telepon valid")
				continue
			}
			newUser := entity.User{
				Email:       email,
				Password:    password,
				Name:        name,
				Address:     address,
				PhoneNumber: phoneNumber,
			}
			err = handler.Register(newUser)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Registrasi Berhasil!")
			}
		case "3":
			return
		default:
			fmt.Println("Input yang dimasukkan tidak valid")
		}
	}
}

func userInput(text string, scanner *bufio.Scanner) string {
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
}

func scanEnter() {
	fmt.Print("\nTekan enter untuk kembali ke menu ")
	fmt.Scanln()
}
