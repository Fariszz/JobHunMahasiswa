package mahasiswa

import (
	"database/sql"
)

type Repository interface {
	FindAll() ([]Mahasiswa, error)
	FindByMahasiswaId(ID int) (Mahasiswa, error)
	Save(m Mahasiswa) (Mahasiswa, error)
	Update(m Mahasiswa) (Mahasiswa, error)
	Delete(m Mahasiswa) (Mahasiswa, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Mahasiswa, error) {
	rows, err := r.db.Query("SELECT Nama, Usia, Gender, Tanggal_Registrasi FROM Mahasiswa")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	mahasiswas := make([]Mahasiswa, 0)

	for rows.Next() {
		var m Mahasiswa
		err := rows.Scan(&m.Nama, &m.Usia, &m.Gender, &m.Tanggal_Registrasi)

		if err != nil {
			return nil, err
		}
		mahasiswas = append(mahasiswas, m)
	}

	return mahasiswas, nil
}

func (r *repository) FindByMahasiswaId(ID int) (Mahasiswa, error) {
	m := Mahasiswa{}
	h := Hobi{}

	query := `
	SELECT m.Id, m.Nama, m.Usia, m.Gender, m.Tanggal_Registrasi, h.Nama_Hobi
	from Mahasiswa m 	
	INNER JOIN Mahasiswa_Hobi mh ON m.Id = mh.Id_Mahasiswa
	INNER JOIN Hobi h ON h.Id = mh.Id_Hobi	
	WHERE m.Id = ?
	`

	rows, err := r.db.Query(query, ID)

	if err != nil {
		return m, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&m.ID, &m.Nama, &m.Usia, &m.Gender, &m.Tanggal_Registrasi, &h.NamaHobi)

		if err != nil {
			return m, err
		}

		m.Hobi = append(m.Hobi, h)
	}

	return m, nil
}

func (r *repository) Save(m Mahasiswa) (Mahasiswa, error) {
	// start transaction

	tx, err := r.db.Begin()

	if err != nil {
		return m, err
	}

	// insert into mahasiswa table
	mahasiswaStmt, err := tx.Prepare("INSERT INTO Mahasiswa (Nama, Usia, Gender, Tanggal_Registrasi) VALUES (?, ?, ?, ?)")

	if err != nil {
		return m, err
	}

	defer mahasiswaStmt.Close()

	mahasiswaRes, err := mahasiswaStmt.Exec(m.Nama, m.Usia, m.Gender, m.Tanggal_Registrasi)

	if err != nil {
		return m, err
	}

	mahasiswaId, err := mahasiswaRes.LastInsertId()

	if err != nil {
		return m, err
	}

	// insert into jurusan table
	jurusanStmt, err := tx.Prepare("INSERT INTO Jurusan (Nama_Jurusan) VALUES (?)")

	if err != nil {
		return m, err
	}

	defer jurusanStmt.Close()

	_, err = jurusanStmt.Exec(m.Jurusan.NamaJurusan)

	if err != nil {
		return m, err
	}

	// insert into hobi table and Mahasiswa_Hobi table
	for _, hobi := range m.Hobi {
		// insert data into table hobi
		hobiStmt, err := tx.Prepare("INSERT INTO Hobi (Nama_Hobi) VALUES (?)")

		if err != nil {
			return m, err
		}

		defer hobiStmt.Close()

		hobiRes, err := hobiStmt.Exec(hobi.NamaHobi)

		if err != nil {
			return m, err
		}

		hobiId, err := hobiRes.LastInsertId()

		if err != nil {
			return m, err
		}

		// insert data ke table Mahasiswa_Hobi
		mahasiswaHobiStmt, err := tx.Prepare("INSERT INTO Mahasiswa_Hobi (Id_Mahasiswa, Id_Hobi) VALUES (?, ?)")

		if err != nil {
			return m, err
		}

		defer mahasiswaHobiStmt.Close()

		_, err = mahasiswaHobiStmt.Exec(mahasiswaId, hobiId)

		if err != nil {
			return m, err
		}
	}

	// commit transaction
	err = tx.Commit()

	if err != nil {
		return m, err
	}

	return m, nil
}

func (r *repository) Update(m Mahasiswa) (Mahasiswa, error) {

	// update data in table Mahasiswa
	_, err := r.db.Query(`
	UPDATE mahasiswa SET Nama=?, Usia=?, Gender=?, Tanggal_Registrasi=?
	WHERE Id=?
	`, m.Nama, m.Usia, m.Gender, m.Tanggal_Registrasi, m.ID)

	if err != nil {
		return m, err
	}

	return m, nil
}

func (r *repository) Delete(m Mahasiswa) (Mahasiswa, error) {
	// start transaction
	tx, err := r.db.Begin()

	if err != nil {
		return m, err
	}

	// delete data from table Mahasiswa_Hobi
	mahasiswaHobiStmt, err := tx.Prepare("DELETE FROM Mahasiswa_Hobi WHERE Id_Mahasiswa = ?")

	if err != nil {
		return m, err
	}

	defer mahasiswaHobiStmt.Close()

	_, err = mahasiswaHobiStmt.Exec(m.ID)

	if err != nil {
		return m, err
	}

	// delete data from table Mahasiswa

	mahasiswaStmt, err := tx.Prepare("DELETE FROM Mahasiswa WHERE Id = ?")

	if err != nil {
		return m, err
	}

	defer mahasiswaStmt.Close()

	_, err = mahasiswaStmt.Exec(m.ID)

	if err != nil {
		return m, err
	}

	// Delete the related records in the hobi table that are no longer referenced in the mahasiswa_hobi table
	_, err = tx.Exec("DELETE FROM Hobi WHERE NOT EXISTS (SELECT 1 FROM Mahasiswa_Hobi WHERE Hobi.Id = Mahasiswa_Hobi.Id_Hobi)")
	if err != nil {
		tx.Rollback()
	}

	// commit transaction
	err = tx.Commit()

	if err != nil {
		return m, err
	}

	return m, nil
}
