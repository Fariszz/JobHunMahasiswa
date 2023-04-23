package mahasiswa

import "time"

type GetMahasiswaDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateMahasiswaInput struct {
	Nama               string    `json:"nama" binding:"required"`
	Usia               int       `json:"usia" binding:"required"`
	Gender             int       `json:"gender" binding:"required"`
	Tanggal_Registrasi time.Time `json:"tanggal_registrasi" binding:"required"`
	Jurusan            Jurusan   `json:"jurusan"`
	Hobi               []Hobi    `json:"hobi"`
}

type CreateHobiInput struct {
	NamaHobi string `json:"nama_hobi" binding:"required"`
}
