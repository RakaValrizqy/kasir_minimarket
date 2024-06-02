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
	tanggal string
}

type tabProd [nMax]produk
type tabTrans [nMax]transaksi

func main(){
	var listProd tabProd
	var listTrans tabTrans
	var pilih, nProd, nTrans int
	var cari, tgl string

	nProd = 0
	nTrans = 0

	for pilih != 9 {
		menu()
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			tambahProd(&listProd,&nProd)
			inserSortProd(&listProd,nProd)
		case 2:
			fmt.Println("Masukkan nama produk yang ingin dicari:")
			fmt.Scan(&cari)
			cariProd(listProd,nProd,cari)
		case 3:
			ubahProd(&listProd,nProd)
		case 4:
			tampilProd(listProd,nProd)
			fmt.Println("Masukkan nama produk yang ingin dihapus:")
			fmt.Scan(&cari)
			hapusProd(&listProd,&nProd,cari)
		case 5:
			tampilProd(listProd,nProd)
		case 6:
			tambahTrans(&listTrans,listProd,&nTrans,nProd)
		case 7:
			selSortTrans(&listTrans,nTrans)
			tampilTrans(listTrans,nTrans)
		case 8:
			fmt.Print("Tanggal transaksi: ")
			fmt.Scan(&tgl)
			fmt.Printf("Omzet pada tanggal %s: %d\n",tgl,hitungOmzet(listTrans,nTrans,tgl))
		default:
			fmt.Println("Pilih menu 1-9")
		}
	}

}

func menu(){
	fmt.Println("MAIN MENU")
	fmt.Println("1. Tambah produk\n2. Cari produk\n3. Ubah produk\n4. Hapus produk\n5. Lihat produk\n6. Tambah transaksi\n7. Tampil transaksi\n8. Omzet berdasarkan tanggal\n9. Exit")
	fmt.Print("Pilih: ")
}

func tambahProd(A *tabProd, n *int){
	/*	I.S terdefinisi n adalah banyak data dalam array
		F.S Array A berisi data produk sebanyak n	*/
	var nama string
	
	fmt.Println("Tambah produk: (nama) (harga)\nKetik done untuk mengakhiri")
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
	fmt.Println("---------------------------------------------")
	fmt.Printf("%27s\n","LIST PRODUK")
	fmt.Println("---------------------------------------------")
	fmt.Printf("|%5s| %20s| %15s|\n","Nomor","Nama Produk","Harga")
	for i<n {
		// fmt.Println(i+1,".",A[i].nama,A[i].harga)
		fmt.Printf("|%5d| %20s| %15d|\n",i+1,A[i].nama,A[i].harga)
		i++
	}
	fmt.Println("---------------------------------------------")
}

func cariIndex(A tabProd, n int, x string)int{
	//mengembalikan index dari array A jika terdapat data dalam array A yang sama dengan x
	var index, ki, te, ka int

	index = -1
	ki=0
	ka=n-1
	
	for index == -1 && ki<=ka {
		te = (ki+ka)/2
		if x < A[te].nama {
			ka = te-1
		} else if x > A[te].nama {
			ki = te+1
		} else {
			index = te
		}
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
		fmt.Println("Produk berhasil dihapus")
		tampilProd(*A,*n)
	}
}

func tambahTrans(T *tabTrans, P tabProd, nT *int, nP int){
	/*	I.S terdefinisi array P, bilangan bulat nP
		F.S array T berisi data produk sebanyak nT	*/
	var nom, banyak int
	var tgl string

	tampilProd(P,nP)
	fmt.Print("Masukkan tanggal transaksi: ")
	fmt.Scan(&tgl)
	fmt.Println("Tambah transaksi: (nomor produk) (banyak produk)")
	fmt.Println("Ketik -1 untuk mengakhiri")
	fmt.Scan(&nom)
	for nom != -1 {
		fmt.Scan(&banyak)
		T[*nT].prod = P[nom-1]
		T[*nT].banyak = banyak
		T[*nT].subtotal = P[nom-1].harga*banyak
		T[*nT].tanggal = tgl
		*nT++

		fmt.Scan(&nom)	
	}
}

func tampilTrans(T tabTrans, n int){
	/*	I.S terdefinisi array T dan n
		F.S menampilkan seluruh data dalam array T	*/
	var i int

	i=0

	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Printf("%50s\n","LIST TRANSAKSI")
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Printf("|%5s| %20s| %15s| %8s| %15s| %17s|\n","Nomor","Nama Produk","Harga","Banyak","Subtotal","Tanggal transaksi")
	for i<n {
		fmt.Printf("|%5d| %20s| %15d| %8d| %15d| %17s|\n",i+1,T[i].prod.nama,T[i].prod.harga,T[i].banyak,T[i].subtotal,T[i].tanggal)
		i++
	}
	fmt.Println("--------------------------------------------------------------------------------------------")
}

func hitungOmzet(T tabTrans, n int, tgl string)int{
	//mengembalikan nilai omzet dari array T jika T[].tanggal == tgl
	var i, omzet int

	i=0
	omzet = 0

	for i<n {
		if T[i].tanggal == tgl {
			omzet += T[i].subtotal
		}
		i++
	}
	return omzet
}

func selSortTrans(A *tabTrans, n int){
	/*	I.S Terdefinisi Array list transaksi A dengan panjang n dan diasumsikan belum terurut
		F.S Array list transaksi A dengan panjang n terurut berdasarkan subtotal secara ascending atau descending sesuai keingininan pengguna menggunakan metode selection sorting*/
	var i, j, max, min, pass, sort int
	var temp transaksi

	fmt.Println("Urutkan berdasarkan subtotal:\n1. Terbesar\n2. Terkecil")
	fmt.Print("Pilihan anda: ")
	fmt.Scan(&sort)

	pass = n-1

	switch sort {
		case 1:
			for i=0; i<pass; i++ {
				max = i
				for j=i+1; j<n; j++ {
					if A[max].subtotal < A[j].subtotal {
						max = j
					}
					
				}
		
				temp = A[i]
				A[i] = A[max]
				A[max] = temp
			}
		case 2:
			for i=0; i<pass; i++ {
				min = i
				for j=i+1; j<n; j++ {
					if A[min].subtotal > A[j].subtotal {
						min = j
					}
					
				}
		
				temp = A[i]
				A[i] = A[min]
				A[min] = temp
			}
		default:
			fmt.Println("Pilih 1-2")
	}
	
}

func inserSortProd(A *tabProd, n int){
	/*	I.S Terdefinisi Array list transaksi A dengan panjang n, n adalah banyak data dalam array
		F.S Array A berisi data produk sebanyak n dan mengurutkan sesuai nama secara ascending menggunakan metode insertion sorting	*/
	var i, pass int
	var temp produk

	pass = 1

	for pass < n {
		i = pass
		temp = A[pass]
		for i > 0 && temp.nama < A[i-1].nama {
			A[i] = A[i-1]
			i -= 1
		}
		A[i] = temp
		pass += 1
	}
}