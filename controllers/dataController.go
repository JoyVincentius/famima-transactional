package controllers

import (
	"famima-transactional/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllKaryawan(c *gin.Context) {
	var karyawanList []models.Karyawan
	if err := GetGormDB().Find(&karyawanList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data karyawan"})
		return
	}

	var karyawanResponses []models.Karyawan
	for _, karyawan := range karyawanList {
		karyawanResponses = append(karyawanResponses, models.Karyawan{
			IDKaryawan:     karyawan.IDKaryawan,
			NamaLengkap:    karyawan.NamaLengkap,
			Email:          karyawan.Email,
			Password:       "******", // Masking password
			NoTelepon:      karyawan.NoTelepon,
			Alamat:         karyawan.Alamat,
			TanggalLahir:   karyawan.TanggalLahir,
			JenisKelamin:   karyawan.JenisKelamin,
			Jabatan:        karyawan.Jabatan,
			Departemen:     karyawan.Departemen,
			TanggalMasuk:   karyawan.TanggalMasuk,
			StatusKaryawan: karyawan.StatusKaryawan,
		})
	}

	c.JSON(http.StatusOK, karyawanResponses)
}

func GetKaryawanByID(c *gin.Context) {
	id := c.Param("id")
	var karyawan models.Karyawan

	if err := GetGormDB().First(&karyawan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Karyawan tidak ditemukan"})
		return
	}

	karyawan.Password = "******" // Masking password
	c.JSON(http.StatusOK, karyawan)
}
