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
	var listTrans tabTrans
	var pilih, nProd, nTrans int
	var cari string

	nProd = 0
	nTrans = 0

	for pilih != 8 {
		menu()
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			tambahProd(&listProd,&nProd)
		case 2:
			tampilProd(listProd,nProd)
		case 3:
			tambahTrans(&listTrans,listProd,&nTrans,nProd)
		case 4:
			tampilTrans(listTrans,nTrans)
		case 5:
			fmt.Scan(&cari)
			cariProd(listProd,nProd,cari)
		case 6:
			ubahProd(&listProd,nProd)
		case 7:
			fmt.Scan(&cari)
			hapusProd(&listProd,&nProd,cari)
		}
	}

}

func menu(){
	fmt.Println("MAIN MENU")
	fmt.Println("1. Tambah produk\n2. Lihat produk\n3. Tambah transaksi\n4. Tampil transaksi\n5. Cari produk\n6. Ubah produk\n7. Hapus produk\n8. Exit")
	fmt.Print("Pilih: ")
}

func tambahProd(A *tabProd, n *int){
	/*	I.S terdefinisi n adalah banyak data dalam array
		F.S Array A berisi data produk sebanyak n	*/
	var nama string
	
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

func cariIndex(A tabProd, n int, x string)int{
	//mengembalikan index dari array A jika terdapat data dalam array A yang sama dengan x
	var i, index int

	index = -1
	i=0
	for i<n && index == -1 {
		if A[i].nama == x {
			index = i
		}
		i++
	}
	return index
}

func cariProd(A tabProd, n int, x string){
	/*	I.S terdefinisi array A, bilangan bulat n dan string x
		F.S menampilkan data dalam array A jika nama produk sama dengan x	*/
	var index int

	index = cariIndex(A,n,x)
	if index == -1 {
		fmt.Println("Produk tidak ditemukan")
	} else {
		fmt.Println("Index produk:",index,"\nNama produk:",A[index].nama,"\nHarga produk:",A[index].harga)
	}
	fmt.Println(" ")
}

func ubahProd(A *tabProd, n int){
	/*	I.S terdefinisi array A dan bilangan bulat n
		F.S mengubah data dalam array A		*/
	var nom int

	tampilProd(*A,n)
	fmt.Print("Nomor produk yang ingin diubah: ")
	fmt.Scan(&nom)

	fmt.Println("Data produk baru: (nama) (harga)")
	fmt.Scan(&A[nom-1].nama,&A[nom-1].harga)
}

func hapusProd(A *tabProd, n *int, x string){
	/*	I.S terdefinisi array A, bilangan bulat n dan string x
		F.S menghapus data dalam array A jika nama produk sama dengan x	*/
	var index int

	index = cariIndex(*A,*n,x)
	if index == -1 {
		fmt.Println("Produk tidak ditemukan")
	} else {
		*n -= 1
		for index < *n {
			A[index] = A[index+1]
			index++
		}
	}
	fmt.Println(" ")
}

func tambahTrans(T *tabTrans, P tabProd, nT *int, nP int){
	/*	I.S terdefinisi array P, bilangan bulat nP
		F.S array T berisi data produk sebanyak nT	*/
	var nom, banyak int

	tampilProd(P,nP)
	fmt.Println("Tambah transaksi: (nomor) (banyak)")
	fmt.Println("Cancel: input -1")
	fmt.Scan(&nom)
	for nom != -1 {
		fmt.Scan(&banyak)
		T[*nT].prod = P[nom-1]
		T[*nT].banyak = banyak
		T[*nT].subtotal = P[nom-1].harga*banyak
		*nT++

		fmt.Scan(&nom)	
	}
}

func tampilTrans(T tabTrans, n int){
	/*	I.S terdefinisi array T dan n
		F.S menampilkan seluruh data dalam array T	*/
	var i int

	i=0
	for i<n {
		fmt.Println(i+1,T[i].prod.nama,T[i].prod.harga,T[i].banyak,T[i].subtotal)
		i++
	}
	fmt.Println(" ")
}