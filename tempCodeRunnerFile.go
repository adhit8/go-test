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
	http.HandleFunc("/test", getWelcomeTest)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
	handleRequests()
}

func getWelcomeTest(w http.ResponseWriter, r *http.Request) {
	//set variabel
	var no int
	var mhs Mahasiswa
	var mhsx []Mahasiswa
	var array_data = make([]string, 5)

	//isi variabel array
	array_data = []string{"Ari Mahadi", "Naufal Kusnanto", "Tigor Manurung", "Surya Nagasaputra", "Abdul Puteh"}

	for i := 0; i < len(array_data); i++ {
		no = i + 1
		mhs = Mahasiswa{no, string(array_data[i])}
		mhsx = append(mhsx, mhs)
	}

	json_data, err := json.Marshal(mhsx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", string(json_data))
	json.NewEncoder(w).Encode(mhsx)
}
