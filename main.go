package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Mahasiswa struct {
	No   int
	Nama string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to my home page!")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/get_convert_array_to_json_object", getConvertArrayToJSONObject)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
	handleRequests()
}

func getConvertArrayToJSONObject(w http.ResponseWriter, r *http.Request) {
	//set variabel
	var no int
	var mhs Mahasiswa
	var mhsx []Mahasiswa
	var array_data = make([]string, 5) //set max index pada array

	//isi variabel array
	array_data = []string{"Ari Mahadi", "Naufal Kusnanto", "Tigor Manurung", "Surya Nagasaputra", "Abdul Puteh"}

	//looping untuk menampung data
	for i := 0; i < len(array_data); i++ {
		//set nomor (increment++)
		no = i + 1
		//set data siswa pada objek dengan tipe struct
		mhs = Mahasiswa{no, string(array_data[i])}
		//Tambahkan data siswa
		mhsx = append(mhsx, mhs)
	}
	//encoding data struct to JSON string
	json_data, err := json.Marshal(mhsx)
	if err != nil {
		panic(err)
	}
	//menampilkan data pada terminal atau console
	fmt.Printf("%s\n", string(json_data))
	//menampilkan data dengan tipe struct ke JSON Object
	json.NewEncoder(w).Encode(mhsx)
}
