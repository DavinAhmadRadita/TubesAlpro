package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type PemainInfo struct {
	Name   string
	Age    string
	Height string
	Weight string
	Nation string
}

type StatistikPemain struct {
	OVR string
	PAC string
	SHO string
	PAS string
	DRI string
	DEF string
	PHY string
}

type KarierPemain struct {
	Position  string
	Preferred string
	AltPos    string
	League    string
	Team      string
}

type Pemain struct {
	Info   PemainInfo
	Stat   StatistikPemain
	Karier KarierPemain
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
		if i == 0 || len(row) < 17 {
			continue
		}
		list = append(list, Pemain{
			Info: PemainInfo{
				Name:   row[0],
				Age:    row[13],
				Height: row[10],
				Weight: row[11],
				Nation: row[14],
			},
			Stat: StatistikPemain{
				OVR: row[1],
				PAC: row[2],
				SHO: row[3],
				PAS: row[4],
				DRI: row[5],
				DEF: row[6],
				PHY: row[7],
			},
			Karier: KarierPemain{
				Position:  row[8],
				Preferred: row[9],
				AltPos:    row[12],
				League:    row[15],
				Team:      row[16],
			},
		})
	}
	return list
}

func selectionSortByOVR(data []Pemain) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if toInt(data[j].Stat.OVR) > toInt(data[maxIdx].Stat.OVR) {
				maxIdx = j
			}
		}
		data[i], data[maxIdx] = data[maxIdx], data[i]
	}
}

func toInt(s string) int {
	var val int
	fmt.Sscanf(s, "%d", &val)
	return val
}

func searchByName(data []Pemain, target string) []Pemain {
	var result []Pemain
	target = strings.ToLower(target)
	for _, d := range data {
		if strings.HasPrefix(strings.ToLower(d.Info.Name), target) {
			result = append(result, d)
		}
	}
	return result
}

func searchByNation(data []Pemain, nation string) []Pemain {
	var result []Pemain
	for _, d := range data {
		if strings.EqualFold(d.Info.Nation, nation) {
			result = append(result, d)
		}
	}
	return result
}

func searchByTeam(data []Pemain, team string) []Pemain {
	var result []Pemain
	for _, d := range data {
		if strings.EqualFold(d.Karier.Team, team) {
			result = append(result, d)
		}
	}
	return result
}

func searchByAge(data []Pemain, age string) []Pemain {
	var result []Pemain
	for _, d := range data {
		if d.Info.Age == age {
			result = append(result, d)
		}
	}
	return result
}

func tambahPemain(reader *bufio.Reader, pemainList *[]Pemain) {
	fmt.Println("Masukkan data pemain baru:")
	var p Pemain

	fmt.Print("Nama: ")
	p.Info.Name, _ = reader.ReadString('\n')
	p.Info.Name = strings.TrimSpace(p.Info.Name)

	fmt.Print("OVR: ")
	p.Stat.OVR, _ = reader.ReadString('\n')
	p.Stat.OVR = strings.TrimSpace(p.Stat.OVR)

	fmt.Print("PAC: ")
	p.Stat.PAC, _ = reader.ReadString('\n')
	p.Stat.PAC = strings.TrimSpace(p.Stat.PAC)

	fmt.Print("SHO: ")
	p.Stat.SHO, _ = reader.ReadString('\n')
	p.Stat.SHO = strings.TrimSpace(p.Stat.SHO)

	fmt.Print("PAS: ")
	p.Stat.PAS, _ = reader.ReadString('\n')
	p.Stat.PAS = strings.TrimSpace(p.Stat.PAS)

	fmt.Print("DRI: ")
	p.Stat.DRI, _ = reader.ReadString('\n')
	p.Stat.DRI = strings.TrimSpace(p.Stat.DRI)

	fmt.Print("DEF: ")
	p.Stat.DEF, _ = reader.ReadString('\n')
	p.Stat.DEF = strings.TrimSpace(p.Stat.DEF)

	fmt.Print("PHY: ")
	p.Stat.PHY, _ = reader.ReadString('\n')
	p.Stat.PHY = strings.TrimSpace(p.Stat.PHY)

	fmt.Print("Posisi: ")
	p.Karier.Position, _ = reader.ReadString('\n')
	p.Karier.Position = strings.TrimSpace(p.Karier.Position)

	fmt.Print("Preferred kaki: ")
	p.Karier.Preferred, _ = reader.ReadString('\n')
	p.Karier.Preferred = strings.TrimSpace(p.Karier.Preferred)

	fmt.Print("Tinggi: ")
	p.Info.Height, _ = reader.ReadString('\n')
	p.Info.Height = strings.TrimSpace(p.Info.Height)

	fmt.Print("Berat: ")
	p.Info.Weight, _ = reader.ReadString('\n')
	p.Info.Weight = strings.TrimSpace(p.Info.Weight)

	fmt.Print("Posisi Alternatif: ")
	p.Karier.AltPos, _ = reader.ReadString('\n')
	p.Karier.AltPos = strings.TrimSpace(p.Karier.AltPos)

	fmt.Print("Umur: ")
	p.Info.Age, _ = reader.ReadString('\n')
	p.Info.Age = strings.TrimSpace(p.Info.Age)

	fmt.Print("Negara: ")
	p.Info.Nation, _ = reader.ReadString('\n')
	p.Info.Nation = strings.TrimSpace(p.Info.Nation)

	fmt.Print("Liga: ")
	p.Karier.League, _ = reader.ReadString('\n')
	p.Karier.League = strings.TrimSpace(p.Karier.League)

	fmt.Print("Tim: ")
	p.Karier.Team, _ = reader.ReadString('\n')
	p.Karier.Team = strings.TrimSpace(p.Karier.Team)

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
			p.Info.Name, p.Stat.OVR, p.Stat.PAC, p.Stat.SHO, p.Stat.PAS, p.Stat.DRI, p.Stat.DEF, p.Stat.PHY,
			p.Karier.Position, p.Karier.Preferred, p.Info.Height, p.Info.Weight, p.Karier.AltPos,
			p.Info.Age, p.Info.Nation, p.Karier.League, p.Karier.Team,
		})
	}
}

func main() {
	pemainList := bacaCSV("player.csv")
	selectionSortByOVR(pemainList)
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
					fmt.Printf("- %s (%s tahun), %s - %s | OVR: %s | Posisi: %s\n",
						p.Info.Name, p.Info.Age, p.Karier.Team, p.Info.Nation, p.Stat.OVR, p.Karier.Position)
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
					fmt.Printf("- %s (%s tahun), %s - %s | OVR: %s | Posisi: %s\n",
						p.Info.Name, p.Info.Age, p.Karier.Team, p.Info.Nation, p.Stat.OVR, p.Karier.Position)
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
					fmt.Printf("- %s (%s tahun), %s - %s | OVR: %s | Posisi: %s\n",
						p.Info.Name, p.Info.Age, p.Karier.Team, p.Info.Nation, p.Stat.OVR, p.Karier.Position)
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
					fmt.Printf("- %s (%s tahun), %s - %s | OVR: %s | Posisi: %s\n",
						p.Info.Name, p.Info.Age, p.Karier.Team, p.Info.Nation, p.Stat.OVR, p.Karier.Position)
				}
			}
		case "5":
			tambahPemain(reader, &pemainList)
			selectionSortByOVR(pemainList)
			tulisCSV("player.csv", pemainList)
		case "6":
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Coba lagi.")
		}
	}
}
