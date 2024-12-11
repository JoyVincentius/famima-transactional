package controllers

import (
	"famima-transactional/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Absensi(c *gin.Context) {
	var absensi models.Absensi
	if err := c.ShouldBindJSON(&absensi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storedIDKaryawan string
	validateID := "SELECT id_karyawan FROM karyawan WHERE id_karyawan = ?"
	err := db.QueryRow(validateID, absensi.IDKaryawan).Scan(&storedIDKaryawan)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ID karyawan tidak ditemukan"})
		return
	}

	query := "CALL sp_absensi(?, ?, ?, ?, ?, ?)"
	_, err = db.Exec(query, absensi.IDKaryawan, absensi.TanggalAbsensi, absensi.JamAbsensi, absensi.KoordinatLat, absensi.KoordinatLong, absensi.StatusAbsensi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Absensi berhasil dicatat"})
}
