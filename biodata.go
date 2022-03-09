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
		Pekerjaan: "Backend Engineer",
		Alasan:    "Bekerja",
	}

	data2 := Biodata{
		Nama:      "Jordan",
		Alamat:    "Medan",
		Pekerjaan: "Human Resource",
		Alasan:    "Mencari kehidupan yang lebih baik",
	}

	data3 := Biodata{
		Nama:      "Maruli",
		Alamat:    "Palembang",
		Pekerjaan: "DevOps",
		Alasan:    "Mencari Nafkah",
	}

	datas := map[int]Biodata{1: data1, 2: data2, 3: data3}

	data, _ := strconv.Atoi(reader)
	dataPeserta := datas[data]

	result:= fmt.Sprintf("Nama : %s, Alamat : %s, Pekerjaan : %s, Alasan : %s", dataPeserta.Nama, dataPeserta.Alamat, dataPeserta.Pekerjaan, dataPeserta.Alasan)

	fmt.Println(result)
}
