package mahasiswa

import "time"

type Mahasiswa struct {
	ID                 int       `json:"id"`
	Nama               string    `json:"nama"`
	Usia               int       `json:"usia"`
	Gender             int       `json:"gender"`
	Tanggal_Registrasi time.Time `json:"tanggal_registrasi"`
	Jurusan            Jurusan   `json:"jurusan"`
	Hobi               []Hobi    `json:"hobi"`
}

type Jurusan struct {
	ID   int    `json:"id"`
	NamaJurusan string `json:"nama_jurusan"`
}

type Hobi struct {
	ID   int    `json:"id"`
	NamaHobi string `json:"nama_hobi"`
}
