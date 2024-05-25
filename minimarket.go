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

	for pilih != 8 {
		menu()
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			tambahProd(&listProd,&nProd)
			inserSortProd(&listProd,nProd)
		case 2:
			tampilProd(listProd,nProd)
		case 3:
			tambahTrans(&listTrans,listProd,&nTrans,nProd)
		case 4:
			selSortTrans(&listTrans,nTrans)
			tampilTrans(listTrans,nTrans)
		case 5:
			fmt.Scan(&cari)
			cariProd(listProd,nProd,cari)
		case 6:
			ubahProd(&listProd,nProd)
		case 7:
			fmt.Scan(&cari)
			hapusProd(&listProd,&nProd,cari)
		case 9:
			fmt.Print("Tanggal transaksi: ")
			fmt.Scan(&tgl)
			fmt.Printf("Omzet pada hari %s: %d\n",tgl,hitungOmzet(listTrans,nTrans,tgl))
		default:
			fmt.Println("Pilih menu 1-9")
		}
	}

}

func menu(){
	fmt.Println("MAIN MENU")
	fmt.Println("1. Tambah produk\n2. Lihat produk\n3. Tambah transaksi\n4. Tampil transaksi\n5. Cari produk\n6. Ubah produk\n7. Hapus produk\n8. Exit\n9. Omzet hari ini")
	fmt.Print("Pilih: ")
}

func tambahProd(A *tabProd, n *int){
	/*	I.S terdefinisi n adalah banyak data dalam array
		F.S Array A berisi data produk sebanyak n	*/
	var nama string
	
	fmt.Println("Tambah produk: (nama) (harga)\nKetik done ketika selesai")
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
	fmt.Println("\t\tLIST PRODUK")
	fmt.Printf("\tnomor \tnama produk \tharga\n")
	for i<n {
		// fmt.Println(i+1,".",A[i].nama,A[i].harga)
		fmt.Printf("\t%d. \t%s \t%d\n",i+1,A[i].nama,A[i].harga)
		i++
	}
	fmt.Println(" ")
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
	}
	fmt.Println(" ")
}

func tambahTrans(T *tabTrans, P tabProd, nT *int, nP int){
	/*	I.S terdefinisi array P, bilangan bulat nP
		F.S array T berisi data produk sebanyak nT	*/
	var nom, banyak int
	var tgl string

	tampilProd(P,nP)
	fmt.Print("Masukkan tanggal transaksi: ")
	fmt.Scan(&tgl)
	fmt.Println("Tambah transaksi: (nomor) (banyak)")
	fmt.Println("Cancel: input -1")
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
	for i<n {
		fmt.Println(i+1,T[i].prod.nama,T[i].prod.harga,T[i].banyak,T[i].subtotal,T[i].tanggal)
		i++
	}
	fmt.Println(" ")
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
	var i, j, max, min, pass, sort int
	var temp transaksi

	fmt.Println("Urutkan berdasarkan:\n1. Terbesar\n2. Terkecil")
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