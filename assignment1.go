package main

import ("fmt"
		"os"
		"strconv")


type People struct {
	name string
	address string
	jobs string
	reason string
}

func main() {

	peoples := []People{
		{
			name: "Hilmi 1",
			address: "MNC Kebon Sirih",
			jobs: "Programmer",
			reason: "Biar Pinter",
		},
		{
			name: "Hilmi 2",
			address: "MNC Kebon Sirih",
			jobs: "Programmer",
			reason: "Biar Pinter",
		},
		{
			name: "Hilmi 3",
			address: "MNC Kebon Sirih",
			jobs: "Programmer",
			reason: "Biar Pinter",
		},
		{
			name: "Hilmi 4",
			address: "MNC Kebon Sirih",
			jobs: "Programmer",
			reason: "Biar Pinter",
		},
		{
			name: "Hilmi 5",
			address: "MNC Kebon Sirih",
			jobs: "Programmer",
			reason: "Biar Pinter",
		}, 
	}
	index, err := validateInput(len(peoples))
	if err == "" {
		findRecord(index, peoples)
	} else {
		fmt.Println(err)
	}

}

func validateInput(numrecord int)(int, string) {
	
	if len(os.Args) > 1 {
		input, err  := strconv.Atoi(os.Args[1])
		if err == nil {
			if input <= 0 {
				return 0, "Nomor absen harus angka dan minimal 1"
			} else {
				if input > numrecord {
					return 0, "Data hanya ada "+ strconv.Itoa(numrecord)
				} else {
					return input, ""
				}
			}
		} else {
			return 0, "nomor absen harus angka dan minimal 1"
		}
	} else {
		return 0, "Anda belum memasukkan nomor absen !!!"
	}
	
}

func findRecord(index int, people []People) {
	record := people[index-1]
	fmt.Printf("Data pada nomor absen ke: %d \n", index)
	fmt.Printf("Nama: %s \n", record.name)
	fmt.Printf("Alamat: %s \n", record.address)
	fmt.Printf("Pekerjaan: %s \n", record.jobs)
	fmt.Printf("Alasan Memilih kelas Golang: %s \n", record.reason)
}

