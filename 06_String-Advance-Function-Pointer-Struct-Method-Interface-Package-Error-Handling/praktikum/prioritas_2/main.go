package main

import (
	"fmt"
	"slices"
)

type Pelanggan struct {
	Nama                string
	Alamat              string
	DaftarKendaraanSewa []Kendaraan
}

type Kendaraan struct {
	Merk      string
	Tipe      string
	HargaSewa int
}

type JasaSewaMotor struct {
	DataPelanggan []Pelanggan
}

func (js *JasaSewaMotor) TampilkanPelanggan() {
	if len(js.DataPelanggan) != 0 {
		fmt.Println("\nDATA PELANGGAN")
		for i, cus := range js.DataPelanggan {
			fmt.Printf("%v. Nama: %v, Alamat: %v\n", i+1, cus.Nama, cus.Alamat)
		}
	} else {
		fmt.Println("Err: Data pelanggan masih kosong!")
	}
}

func (js *JasaSewaMotor) SewaKendaraan(p Pelanggan, k Kendaraan) {
	if len(js.DataPelanggan) == 0 {
		p.DaftarKendaraanSewa = append(p.DaftarKendaraanSewa, k)
		js.DataPelanggan = append(js.DataPelanggan, p)
		fmt.Println("Log: Berhasil menambah data pelanggan baru", p.Nama)
	} else {
		dataDitemukan := false
		for i, cus := range js.DataPelanggan {
			if p.Nama == cus.Nama && p.Alamat == cus.Alamat {
				js.DataPelanggan[i].DaftarKendaraanSewa = append(cus.DaftarKendaraanSewa, k)
				fmt.Println("Log: Berhasil menambah sewa motor", p.Nama, k)
				dataDitemukan = true
				break
			}
		}
		if !dataDitemukan {
			p.DaftarKendaraanSewa = append(p.DaftarKendaraanSewa, k)
			js.DataPelanggan = append(js.DataPelanggan, p)
			fmt.Println("Log: Berhasil menambah data pelanggan baru", p.Nama)
		}
	}
}

func (js *JasaSewaMotor) ReturnKendaraan(p Pelanggan, k Kendaraan) {
	if len(js.DataPelanggan) != 0 {
		for i, cus := range js.DataPelanggan {
			if p.Nama == cus.Nama && p.Alamat == cus.Alamat {
				idxMotor := slices.Index(js.DataPelanggan[i].DaftarKendaraanSewa, k)
				if idxMotor != -1 {
					js.DataPelanggan[i].DaftarKendaraanSewa = append(js.DataPelanggan[i].DaftarKendaraanSewa[:idxMotor], js.DataPelanggan[i].DaftarKendaraanSewa[idxMotor+1:]...)
					fmt.Println("Log: Berhasil mengembalikan kendaraan", k.Merk, k.Tipe, "dari pelanggan", p.Nama)
				} else {
					fmt.Println("Err: Kendaraan tidak ditemukan pada pelanggan", p.Nama)
				}
			}
		}
	} else {
		fmt.Println("Err: Data pelanggan masih kosong!")
	}
}

func (js *JasaSewaMotor) TampilkanDaftarKendaraanSewa(p Pelanggan) {
	if len(js.DataPelanggan) != 0 {
		for _, cus := range js.DataPelanggan {
			if p.Nama == cus.Nama && p.Alamat == cus.Alamat {
				if len(cus.DaftarKendaraanSewa) != 0 {
					fmt.Println("Daftar Kendaraan yang Disewa oleh", p.Nama, ":")
					for _, kendaraan := range cus.DaftarKendaraanSewa {
						fmt.Println("Merk:", kendaraan.Merk, "Tipe:", kendaraan.Tipe, "Harga Sewa:", kendaraan.HargaSewa)
					}
				} else {
					fmt.Println("Pelanggan", cus.Nama, "tidak ada menyewa motor")
				}
			}
		}
	} else {
		fmt.Println("Err: Data pelanggan masih kosong!")
	}
}

func main() {
	var PerusahaanABC JasaSewaMotor

	cus1 := Pelanggan{
		Nama:   "Ilham",
		Alamat: "Jalan Ninja",
	}
	cus2 := Pelanggan{
		Nama:   "Kurniawan",
		Alamat: "Jalan Samurai",
	}

	vec1 := Kendaraan{
		Merk:      "Yamaha",
		Tipe:      "Nmax",
		HargaSewa: 500000,
	}
	vec2 := Kendaraan{
		Merk:      "Honda",
		Tipe:      "Vario 125",
		HargaSewa: 200000,
	}
	vec3 := Kendaraan{
		Merk:      "Kawasaki",
		Tipe:      "KLX 150",
		HargaSewa: 300000,
	}

	PerusahaanABC.SewaKendaraan(cus1, vec1)
	PerusahaanABC.SewaKendaraan(cus1, vec2)
	PerusahaanABC.SewaKendaraan(cus1, vec3)

	PerusahaanABC.SewaKendaraan(cus2, vec2)
	PerusahaanABC.SewaKendaraan(cus2, vec3)

	PerusahaanABC.TampilkanPelanggan()
	fmt.Println()

	PerusahaanABC.ReturnKendaraan(cus1, vec2)
	fmt.Println()

	PerusahaanABC.TampilkanDaftarKendaraanSewa(cus1)
	fmt.Println()
	PerusahaanABC.TampilkanDaftarKendaraanSewa(cus2)

	// kira-kira mnrut saya gini sih kak, kyk sistem oop :D
}
