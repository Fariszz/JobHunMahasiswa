package mahasiswa

import "time"

type MahasiswaFormatter struct {	
	Nama               string    `json:"nama"`
	Usia               int       `json:"usia"`
	Gender             int       `json:"gender"`
	Tanggal_Registrasi time.Time `json:"tanggal_registrasi"`
}

func FormatMahasiswa(mahasiswa Mahasiswa) MahasiswaFormatter {
	mahasiswaFormatter := MahasiswaFormatter{}	
	mahasiswaFormatter.Nama = mahasiswa.Nama
	mahasiswaFormatter.Usia = mahasiswa.Usia
	mahasiswaFormatter.Gender = mahasiswa.Gender
	mahasiswaFormatter.Tanggal_Registrasi = mahasiswa.Tanggal_Registrasi

	return mahasiswaFormatter
}

func FormatMahasiswas(mahasiswa []Mahasiswa) []MahasiswaFormatter {
	var mahasiswaFormatter []MahasiswaFormatter

	for _, mahasiswa := range mahasiswa {
		mahasiswaFormatter = append(mahasiswaFormatter, FormatMahasiswa(mahasiswa))
	}

	return mahasiswaFormatter
}

type MahasiswaDetailFormatter struct {
	ID                 int       `json:"id"`
	Nama               string    `json:"nama"`
	Usia               int       `json:"usia"`
	Gender             int       `json:"gender"`
	Tanggal_Registrasi time.Time `json:"tanggal_registrasi"`
	Jurusan            Jurusan   `json:"jurusan"`
	Hobi               []HobiFormatter    `json:"hobi"`
}

func FormatMahasiswaDetail(mahasiswa Mahasiswa) MahasiswaDetailFormatter {
	mahasiswaDetailFormatter := MahasiswaDetailFormatter{}
	mahasiswaDetailFormatter.ID = mahasiswa.ID
	mahasiswaDetailFormatter.Nama = mahasiswa.Nama
	mahasiswaDetailFormatter.Usia = mahasiswa.Usia
	mahasiswaDetailFormatter.Gender = mahasiswa.Gender
	mahasiswaDetailFormatter.Tanggal_Registrasi = mahasiswa.Tanggal_Registrasi
	mahasiswaDetailFormatter.Jurusan = mahasiswa.Jurusan
	
	hobbies := []HobiFormatter{}

	for _, hobi := range mahasiswa.Hobi {
		hobbies = append(hobbies, FormatHobi(hobi))
	}

	mahasiswaDetailFormatter.Hobi = hobbies

	return mahasiswaDetailFormatter
}
type HobiFormatter struct {
	ID       int    `json:"id"`
	NamaHobi string `json:"nama_hobi"`
}

func FormatHobi(hobi Hobi) HobiFormatter {
	hobiFormatter := HobiFormatter{}
	hobiFormatter.ID = hobi.ID
	hobiFormatter.NamaHobi = hobi.NamaHobi

	return hobiFormatter
}

func FormatHobiList(hobi []Hobi) []HobiFormatter {
	var hobiFormatter []HobiFormatter

	for _, hobi := range hobi {
		hobiFormatter = append(hobiFormatter, FormatHobi(hobi))
	}

	return hobiFormatter
}

type CreateMahasiswaFormatter struct {
	Nama               string    `json:"nama"`
	Usia               int       `json:"usia"`
	Gender             int       `json:"gender"`
	Tanggal_Registrasi time.Time `json:"tanggal_registrasi"`
}
