package controllers

import (
	"database/sql"
	"famima-transactional/models"
	"famima-transactional/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func Register(c *gin.Context) {
	var karyawan models.Karyawan
	if err := c.ShouldBindJSON(&karyawan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(karyawan.Password), bcrypt.DefaultCost)
	karyawan.Password = string(hashedPassword)

	query := "CALL sp_register(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, karyawan.NamaLengkap, karyawan.Email, karyawan.Password, karyawan.NoTelepon, karyawan.Alamat, karyawan.TanggalLahir, karyawan.JenisKelamin, karyawan.Jabatan, karyawan.Departemen, karyawan.TanggalMasuk)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Karyawan berhasil didaftarkan"})
}

func Login(c *gin.Context) {
	var karyawan models.Karyawan
	if err := c.ShouldBindJSON(&karyawan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storedPassword string
	query := "SELECT password FROM karyawan WHERE email = ?"
	err := db.QueryRow(query, karyawan.Email).Scan(&storedPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email tidak ditemukan"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(karyawan.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	token, err := utils.GenerateJWT(karyawan.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
