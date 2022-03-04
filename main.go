package main

import (
	"fmt"
	"os"
	"strconv"
)

// type Biodata struct {
// 	ID        int
// 	Nama      string
// 	Alamat    string
// 	Pekerjaan string
// 	Alasan    string
// }

func main() {
	reader := os.Args[1]

	// Biodata := make(map[int]string)

	// Biodata[1]= "Nama:Evrin \n Alamat: Medan \n Pekerjaan: BackEnd Engineer \n Alasan: Pingin Pintar"
	var Biodata = map[int]string{
		1: " Nama:Evrin \n Alamat: Medan \n Pekerjaan: BackEnd Engineer \n Alasan: Pingin Pintar",
	}
	// fmt.Println(Biodata)

	data,_ := strconv.Atoi(reader)
	dataPeserta := Biodata[data]

	fmt.Println(dataPeserta)

}
