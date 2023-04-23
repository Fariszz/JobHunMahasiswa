package mahasiswa

type Service interface {
	GetAllMahasiswa() ([]Mahasiswa, error)
	GetMahasiswaByID(input GetMahasiswaDetailInput) (Mahasiswa, error)
	CreateMahasiswa(input CreateMahasiswaInput) (Mahasiswa, error)
	UpdateMahasiswa(inputID GetMahasiswaDetailInput, inputData CreateMahasiswaInput) (Mahasiswa, error)
	DeleteMahasiswa(input GetMahasiswaDetailInput) (Mahasiswa, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllMahasiswa() ([]Mahasiswa, error) {

	mahasiswas, err := s.repository.FindAll()

	if err != nil {
		return mahasiswas, err
	}

	return mahasiswas, nil
}

func (s *service) GetMahasiswaByID(input GetMahasiswaDetailInput) (Mahasiswa, error) {
	mahasiswa, err := s.repository.FindByMahasiswaId(input.ID)

	if err != nil {
		return mahasiswa, err
	}

	return mahasiswa, nil
}

func (s *service) CreateMahasiswa(input CreateMahasiswaInput) (Mahasiswa, error) {
	mahasiswa := Mahasiswa{}
	mahasiswa.Nama = input.Nama
	mahasiswa.Usia = input.Usia
	mahasiswa.Gender = input.Gender
	mahasiswa.Tanggal_Registrasi = input.Tanggal_Registrasi
	mahasiswa.Jurusan = input.Jurusan
	mahasiswa.Hobi = input.Hobi

	newMahasiswa, err := s.repository.Save(mahasiswa)

	if err != nil {
		return newMahasiswa, err
	}

	return newMahasiswa, nil
}

func (s *service) UpdateMahasiswa(inputID GetMahasiswaDetailInput, inputData CreateMahasiswaInput) (Mahasiswa, error) {
	mahasiswa, err := s.repository.FindByMahasiswaId(inputID.ID)

	if err != nil {
		return mahasiswa, err
	}

	mahasiswa.Nama = inputData.Nama
	mahasiswa.Usia = inputData.Usia
	mahasiswa.Gender = inputData.Gender
	mahasiswa.Tanggal_Registrasi = inputData.Tanggal_Registrasi

	updatedMahasiswa, err := s.repository.Update(mahasiswa)

	if err != nil {
		return updatedMahasiswa, err
	}

	return updatedMahasiswa, nil
}

func (s *service) DeleteMahasiswa(input GetMahasiswaDetailInput) (Mahasiswa, error) {
	mahasiswa, err := s.repository.FindByMahasiswaId(input.ID)

	if err != nil {
		return mahasiswa, err
	}

	deleteMahasiswa, err := s.repository.Delete(mahasiswa)

	if err != nil {
		return mahasiswa, err
	}

	return deleteMahasiswa, nil
}
