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

}