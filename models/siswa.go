package models

type Siswa struct {
	SiswaID            uint   `gorm:"primaryKey" json:"siswa_id"`
	KelasID            uint   `json:"kelas_id"`
	KedinasanID        uint   `json:"kedinasan_id"`
	NomorInduk         string `json:"nomor_induk"`
	Nama               string `json:"nama"`
	JenisKelamin       string `json:"jenis_kelamin"`
	TempatLahir        string `json:"tempat_lahir"`
	TanggalLahir       string `json:"tanggal_lahir"`
	PendidikanTerakhir string `json:"pendidikan_terakhir"`
	Alamat             string `json:"alamat"`
	Telp               string `json:"telp"`
	NamaAyah           string `json:"nama_ayah"`
	TelpAyah           string `json:"telp_ayah"`
	NamaIbu            string `json:"nama_ibu"`
	TelpIbu            string `json:"telp_ibu"`
	TanggalMulaiBimbel string `json:"tanggal_mulai_bimbel"`
	Status             string `json:"status"`
	Aktif              string `json:"aktif"`
	BesaranSPP         int    `json:"besaran_spp"`
	Password           string `json:"-"`
	RFID               string `json:"rfid"`
}

func (Siswa) TableName() string {
	return "siswa"
}
