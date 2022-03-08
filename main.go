package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	reader := os.Args[1]

	data1 := Biodata{
		Nama:      "Evrin",
		Alamat:    "Tarutung",
		Pekerjaan: "Backend",
		Alasan:    "Bekerja",
	}

	data2 := Biodata{
		Nama:      "Dedi Chandra",
		Alamat:    "Balige",
		Pekerjaan: "Backend",
		Alasan:    "Bekerja",
	}


	datas := map[int]Biodata{1: data1, 2: data2}

	data, _ := strconv.Atoi(reader)
	dataPeserta := datas[data]

	fmt.Printf("Nama : %s, Alamat : %s, Pekerjaan : %s, Alasan : %s\n", dataPeserta.Nama,dataPeserta.Alamat, dataPeserta.Pekerjaan, dataPeserta.Alasan)

}