package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type Pemain struct {
	Name      string
	OVR       string
	PAC       string
	SHO       string
	PAS       string
	DRI       string
	DEF       string
	PHY       string
	Position  string
	Preferred string
	Height    string
	Weight    string
	AltPos    string
	Age       string
	Nation    string
	League    string
	Team      string
}

func main() {
	pemainList := bacaCSV("player.csv")

	// Sort by name A-Z
	selectionSortByName(pemainList)

	// Langsung ke input pencarian
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Masukkan nama pemain yang ingin dicari (atau ketik 'sudah' untuk keluar): ")
		cari, _ := reader.ReadString('\n')
		cari = strings.TrimSpace(cari)

		if strings.EqualFold(cari, "sudah") {
			break
		}

		idx := searchByName(pemainList, cari)
		if idx != -1 {
			fmt.Println("Ditemukan:")
			fmt.Println("Name		:", pemainList[idx].Name)
			fmt.Println("Age		:", pemainList[idx].Age)
			fmt.Println("Team		:", pemainList[idx].Team)
			fmt.Println("Nation		:", pemainList[idx].Nation)
			fmt.Println("OVR		:", pemainList[idx].OVR)
			fmt.Println("Position	:", pemainList[idx].Position)
		} else {
			fmt.Println("Pemain tidak ditemukan.")
		}
	}
}


func bacaCSV(filename string) []Pemain {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var list []Pemain
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if len(row) < 17 {
			continue
		}
		list = append(list, Pemain{
			Name:      row[0],
			OVR:       row[1],
			PAC:       row[2],
			SHO:       row[3],
			PAS:       row[4],
			DRI:       row[5],
			DEF:       row[6],
			PHY:       row[7],
			Position:  row[8],
			Preferred: row[9],
			Height:    row[10],
			Weight:    row[11],
			AltPos:    row[12],
			Age:       row[13],
			Nation:    row[14],
			League:    row[15],
			Team:      row[16],
		})
	}
	return list
}

func selectionSortByName(data []Pemain) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if strings.ToLower(data[j].Name) < strings.ToLower(data[minIdx].Name) {
				minIdx = j
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
}

func searchByName(data []Pemain, target string) int {
	for i, d := range data {
		if strings.EqualFold(d.Name, target) {
			return i
		}
	}
	return -1
}
