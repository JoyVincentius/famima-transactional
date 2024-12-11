-- init.sql

CREATE TABLE IF NOT EXISTS karyawan (
    id_karyawan INT AUTO_INCREMENT PRIMARY KEY,
    nama_lengkap VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    no_telepon VARCHAR(15),
    alamat TEXT,
    tanggal_lahir DATE,
    jenis_kelamin ENUM('L', 'P'),
    jabatan VARCHAR(50),
    departemen VARCHAR(50),
    tanggal_masuk DATE,
    status_karyawan ENUM('Aktif', 'Tidak Aktif') DEFAULT 'Aktif',
    foto BLOB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS absensi (
    id_absensi INT AUTO_INCREMENT PRIMARY KEY,
    id_karyawan INT,
    tanggal_absensi DATE,
    jam_absensi TIME,
    koordinat_lat DECIMAL(9,6),
    koordinat_long DECIMAL(9,6),
    status_absensi ENUM('Hadir', 'Tidak Hadir') DEFAULT 'Hadir',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_karyawan) REFERENCES karyawan(id_karyawan)
);

DELIMITER //

CREATE PROCEDURE sp_register(
    IN p_nama_lengkap VARCHAR(100),
    IN p_email VARCHAR(100),
    IN p_password VARCHAR(255),
    IN p_no_telepon VARCHAR(15),
    IN p_alamat TEXT,
    IN p_tanggal_lahir DATE,
    IN p_jenis_kelamin ENUM('L', 'P'),
    IN p_jabatan VARCHAR(50),
    IN p_departemen VARCHAR(50),
    IN p_tanggal_masuk DATE
)
BEGIN
    INSERT INTO karyawan (nama_lengkap, email, password, no_telepon, alamat, tanggal_lahir, jenis_kelamin, jabatan, departemen, tanggal_masuk)
    VALUES (p_nama_lengkap, p_email, p_password, p_no_telepon, p_alamat, p_tanggal_lahir, p_jenis_kelamin, p_jabatan, p_departemen, p_tanggal_masuk);
END //

CREATE PROCEDURE sp_absensi(
    IN p_id_karyawan INT,
    IN p_tanggal_absensi DATE,
    IN p_jam_absensi TIME,
    IN p_koordinat_lat DECIMAL(9,6),
    IN p_koordinat_long DECIMAL(9,6),
    IN p_status_absensi ENUM('Hadir', 'Tidak Hadir')
)
BEGIN
    INSERT INTO absensi (id_karyawan, tanggal_absensi, jam_absensi, koordinat_lat, koordinat_long, status_absensi)
    VALUES (p_id_karyawan, p_tanggal_absensi, p_jam_absensi, p_koordinat_lat, p_koordinat_long, p_status_absensi);
END //

DELIMITER ;
