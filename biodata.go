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
		Nama:      "Dedi Chandra",
		Alamat:    "Balige",
		Pekerjaan: "Backend Engineer",
		Alasan:    "Mencari kehidupan yang lebih baik",
	}

	data3 := Biodata{
		Nama:      "Faikar Achmad Luthfi",
		Alamat:    "Jakarta",
		Pekerjaan: "Backend Engineer",
		Alasan:    "Mencari Nafkah",
	}

	data4 := Biodata{
		Nama:      "Alvin Immanuel Simbolon",
		Alamat:    "Medan",
		Pekerjaan: "Backend Engineer",
		Alasan:    "Mencari Nafkah",
	}

	data5 := Biodata{
		Nama:      "Muhammad Ghifari",
		Alamat:    "Jakarta",
		Pekerjaan: "Backend Engineer",
		Alasan:    "Mencari Nafkah",
	}

	data6 := Biodata{
		Nama:      "Khairul Abdi Dongoran",
		Alamat:    "Sibolga",
		Pekerjaan: "Backend Engineer",
		Alasan:    "Mencari Nafkah",
	}

	datas := map[int]Biodata{1: data1, 2: data2, 3: data3, 4: data4, 5: data5, 6: data6}

	data, _ := strconv.Atoi(reader)

	dataPeserta := datas[data]

	Response(dataPeserta)

}

func Response(data Biodata) {
	var newBiodata Biodata
	result := fmt.Sprintf("Nama : %s\nAlamat : %s\nPekerjaan : %s\nAlasan : %s", data.Nama, data.Alamat, data.Pekerjaan, data.Alasan)
	if data == newBiodata {
		result = "input harus angka 1 - 6"
	}
	fmt.Println(result)
}
