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
	selectionSortByName(pemainList)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n==== Menu Pencarian Pemain ====")
		fmt.Println("1. Cari berdasarkan nama")
		fmt.Println("2. Cari berdasarkan tim")
		fmt.Println("3. Cari berdasarkan umur")
		fmt.Println("4. Cari berdasarkan negara")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu (1-5): ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Masukkan nama pemain: ")
			cari, _ := reader.ReadString('\n')
			cari = strings.TrimSpace(cari)
			idx := searchByName(pemainList, cari)
			if idx != -1 {
				fmt.Println("Ditemukan:")
				fmt.Printf("%s (%s), %s - %s | OVR: %s | Position : %s\n", pemainList[idx].Name, pemainList[idx].Age, pemainList[idx].Team, pemainList[idx].Nation, pemainList[idx].OVR, pemainList[idx].Position)
			} else {
				fmt.Println("Pemain tidak ditemukan.")
			}

		case "2":
			fmt.Print("Masukkan nama tim: ")
			team, _ := reader.ReadString('\n')
			team = strings.TrimSpace(team)
			results := searchByTeam(pemainList, team)
			if len(results) == 0 {
				fmt.Println("Tidak ada pemain ditemukan dalam tim tersebut.")
			} else {
				fmt.Println("Pemain dalam tim", team)
				for _, p := range results {
					fmt.Printf("- %s (%s tahun)\n", p.Name, p.Age)
				}
			}

		case "3":
			fmt.Print("Masukkan umur pemain: ")
			age, _ := reader.ReadString('\n')
			age = strings.TrimSpace(age)
			results := searchByAge(pemainList, age)
			if len(results) == 0 {
				fmt.Println("Tidak ada pemain dengan umur tersebut.")
			} else {
				fmt.Println("Pemain dengan umur", age)
				for _, p := range results {
					fmt.Printf("- %s (%s tahun) - %s\n", p.Name, p.Age, p.Team)
				}
			}

		case "4":
			fmt.Print("Masukkan negara pemain: ")
			nation, _ := reader.ReadString('\n')
			nation = strings.TrimSpace(nation)
			results := searchByNation(pemainList, nation)
			if len(results) == 0 {
				fmt.Println("Tidak ada pemain dari negara tersebut.")
			} else {
				fmt.Println("Pemain dari", nation)
				for _, p := range results {
					fmt.Printf("- %s (%s tahun) - %s\n", p.Name, p.Age, p.Team)
				}
			}

		case "5":
			fmt.Println("Keluar dari program.")
			return

		default:
			fmt.Println("Pilihan tidak valid. Coba lagi.")
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

func searchByNation(data []Pemain, nation string) []Pemain {
	var result []Pemain
	for _, d := range data {
		if strings.EqualFold(d.Nation, nation) {
			result = append(result, d)
		}
	}
	return result
}

func searchByTeam(data []Pemain, team string) []Pemain {
	var result []Pemain
	for _, d := range data {
		if strings.EqualFold(d.Team, team) {
			result = append(result, d)
		}
	}
	return result
}

func searchByAge(data []Pemain, age string) []Pemain {
	var result []Pemain
	for _, d := range data {
		if d.Age == age {
			result = append(result, d)
		}
	}
	return result
}