package models

type Absensi struct {
    IDAbsensi       int     `json:"id_absensi"`
    IDKaryawan      int     `json:"id_karyawan"`
    TanggalAbsensi  string  `json:"tanggal_absensi"`
    JamAbsensi      string  `json:"jam_absensi"`
    KoordinatLat    float64 `json:"koordinat_lat"`
    KoordinatLong   float64 `json:"koordinat_long"`
    StatusAbsensi   string  `json:"status_absensi"`
}
