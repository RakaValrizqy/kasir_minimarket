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
	var listProd tabProd
	// var listTrans tabTrans
	var pilih, nProd int

	nProd = 0

	for pilih != 5 {
		menu()
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			tambahProd(&listProd,&nProd)
		case 2:
			tampilProd(listProd,nProd)
		}
	}

}

func menu(){
	fmt.Println("MAIN MENU")
	fmt.Println("1. Tambah produk\n2. Lihat produk\n3. Tambah transaksi\n4. Omze hari ini\n5. Exit")
	fmt.Print("Pilih: ")
}

func tambahProd(A *tabProd, n *int){
	/*	I.S terdefinisi n adalah banyak data dalam array
		F.S Array A berisi data produk sebanyak n	*/
	var nama string
	// var i int

	// i = *n
	
	fmt.Scan(&nama)
	for nama != "done" {
		A[*n].nama = nama
		fmt.Scan(&A[*n].harga)
		*n++
		fmt.Scan(&nama)
	}
	
}

func tampilProd(A tabProd, n int){
	/*	I.S terdefinisi array A dan n
		F.S menampilkan seluruh data dalam array A	*/
	var i int

	i=0
	fmt.Println("\nLIST PRODUK")
	for i<n {
		fmt.Println(i+1,".",A[i].nama,A[i].harga)
		i++
	}
	fmt.Println(" ")
}