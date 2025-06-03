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

func searchByName(data []Pemain, target string) []Pemain {
	var result []Pemain
	target = strings.ToLower(target)
	for _, d := range data {
		if strings.HasPrefix(strings.ToLower(d.Name), target) {
			result = append(result, d)
		}
	}
	return result
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

func tambahPemain(reader *bufio.Reader, pemainList *[]Pemain) {
	fmt.Println("Masukkan data pemain baru:")
	var p Pemain

	fmt.Print("Nama: ")
	p.Name, _ = reader.ReadString('\n')
	p.Name = strings.TrimSpace(p.Name)

	fmt.Print("OVR: ")
	p.OVR, _ = reader.ReadString('\n')
	p.OVR = strings.TrimSpace(p.OVR)

	fmt.Print("PAC: ")
	p.PAC, _ = reader.ReadString('\n')
	p.PAC = strings.TrimSpace(p.PAC)

	fmt.Print("SHO: ")
	p.SHO, _ = reader.ReadString('\n')
	p.SHO = strings.TrimSpace(p.SHO)

	fmt.Print("PAS: ")
	p.PAS, _ = reader.ReadString('\n')
	p.PAS = strings.TrimSpace(p.PAS)

	fmt.Print("DRI: ")
	p.DRI, _ = reader.ReadString('\n')
	p.DRI = strings.TrimSpace(p.DRI)

	fmt.Print("DEF: ")
	p.DEF, _ = reader.ReadString('\n')
	p.DEF = strings.TrimSpace(p.DEF)

	fmt.Print("PHY: ")
	p.PHY, _ = reader.ReadString('\n')
	p.PHY = strings.TrimSpace(p.PHY)

	fmt.Print("Posisi: ")
	p.Position, _ = reader.ReadString('\n')
	p.Position = strings.TrimSpace(p.Position)

	fmt.Print("Preferred kaki: ")
	p.Preferred, _ = reader.ReadString('\n')
	p.Preferred = strings.TrimSpace(p.Preferred)

	fmt.Print("Tinggi: ")
	p.Height, _ = reader.ReadString('\n')
	p.Height = strings.TrimSpace(p.Height)

	fmt.Print("Berat: ")
	p.Weight, _ = reader.ReadString('\n')
	p.Weight = strings.TrimSpace(p.Weight)

	fmt.Print("Posisi Alternatif: ")
	p.AltPos, _ = reader.ReadString('\n')
	p.AltPos = strings.TrimSpace(p.AltPos)

	fmt.Print("Umur: ")
	p.Age, _ = reader.ReadString('\n')
	p.Age = strings.TrimSpace(p.Age)

	fmt.Print("Negara: ")
	p.Nation, _ = reader.ReadString('\n')
	p.Nation = strings.TrimSpace(p.Nation)

	fmt.Print("Liga: ")
	p.League, _ = reader.ReadString('\n')
	p.League = strings.TrimSpace(p.League)

	fmt.Print("Tim: ")
	p.Team, _ = reader.ReadString('\n')
	p.Team = strings.TrimSpace(p.Team)

	*pemainList = append(*pemainList, p)
	fmt.Println("Pemain berhasil ditambahkan!")
}

func tulisCSV(filename string, data []Pemain) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{
		"Name", "OVR", "PAC", "SHO", "PAS", "DRI", "DEF", "PHY",
		"Position", "Preferred", "Height", "Weight", "AltPos",
		"Age", "Nation", "League", "Team",
	})

	for _, p := range data {
		writer.Write([]string{
			p.Name, p.OVR, p.PAC, p.SHO, p.PAS, p.DRI, p.DEF, p.PHY,
			p.Position, p.Preferred, p.Height, p.Weight, p.AltPos,
			p.Age, p.Nation, p.League, p.Team,
		})
	}
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
		fmt.Println("5. Tambah pemain")
		fmt.Println("6. Keluar")

		fmt.Print("Pilih menu (1-6): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Masukkan nama pemain: ")
			cari, _ := reader.ReadString('\n')
			cari = strings.TrimSpace(cari)
			results := searchByName(pemainList, cari)
			if len(results) == 0 {
				fmt.Println("Tidak ada pemain ditemukan.")
			} else {
				fmt.Println("Hasil pencarian pemain dengan awalan:", cari)
				for _, p := range results {
					fmt.Printf("- %s (%s tahun), %s - %s | OVR: %s | Posisi: %s\n", p.Name, p.Age, p.Team, p.Nation, p.OVR, p.Position)
				}
			}

		case "2":
			fmt.Print("Masukkan nama team: ")
			team, _ := reader.ReadString('\n')
			team = strings.TrimSpace(team)
			results := searchByTeam(pemainList, team)
			if len(results) == 0 {
				fmt.Println("Tidak ada pemain ditemukan dalam team tersebut.")
			} else {
				fmt.Println("Pemain dalam team", team)
				for _, p := range results {
					fmt.Printf("- %s (%s tahun), %s - %s | OVR: %s | Posisi: %s\n", p.Name, p.Age, p.Team, p.Nation, p.OVR, p.Position)
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
					fmt.Printf("- %s (%s tahun), %s - %s | OVR: %s | Posisi: %s\n", p.Name, p.Age, p.Team, p.Nation, p.OVR, p.Position)
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
					fmt.Printf("- %s (%s tahun), %s - %s | OVR: %s | Posisi: %s\n", p.Name, p.Age, p.Team, p.Nation, p.OVR, p.Position)
				}
			}

		case "5":
			tambahPemain(reader, &pemainList)
			selectionSortByName(pemainList)
			tulisCSV("player.csv", pemainList)

		case "6":
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Tolong belajar membaca")
		}
	}
}
