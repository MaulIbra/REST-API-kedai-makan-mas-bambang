package models

type Profile struct {
	ProfileID    string `json:"profileId"`
	User         User	`json:"user"`
	NamaLengkap  string `json:"namaLengkap"`
	JenisKelamin int	`json:"jenisKelamin"`
	Alamat       string `json:"alamat"`
}
