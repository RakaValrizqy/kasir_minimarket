package main
import "fmt"

const nMax int = 1000

type produk struct {
	nama string
	harga int
}

type transaksi struct {
	prod produk
	banyak int
	subtotal int
}

type tabProd [nMax]produk
type tabTrans [nMax]transaksi

func main(){
	// var listProd tabProd
	// var listTrans tabTrans
	var pilih int

	for pilih != 5 {
		menu()
		fmt.Scan(&pilih)
	}

}

func menu(){
	fmt.Println("MAIN MENU")
	fmt.Println("1. Tambah produk\n2. Lihat produk\n3. Tambah transaksi\n4. Omze hari ini\n5. Exit")
	fmt.Print("Pilih: ")
}