package models

type Karyawan struct {
    IDKaryawan   int    `json:"id_karyawan"`
    NamaLengkap  string `json:"nama_lengkap"`
    Email        string `json:"email"`
    Password     string `json:"password"`
    NoTelepon    string `json:"no_telepon"`
    Alamat       string `json:"alamat"`
    TanggalLahir string `json:"tanggal_lahir"`
    JenisKelamin string `json:"jenis_kelamin"`
    Jabatan      string `json:"jabatan"`
    Departemen   string `json:"departemen"`
    TanggalMasuk string `json:"tanggal_masuk"`
    StatusKaryawan string `json:"status_karyawan"`
}
