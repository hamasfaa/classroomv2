-- Create "kelas" table
CREATE TABLE `kelas` (
  `k_id` varchar(255) NOT NULL,
  `k_mata_kuliah` varchar(50) NULL,
  `k_nama_kelas` varchar(50) NULL,
  `k_tanggal_dibuat` date NULL,
  `k_kode_kelas` char(6) NULL,
  PRIMARY KEY (`k_id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "users" table
CREATE TABLE `users` (
  `uid` varchar(255) NOT NULL,
  `u_nama` varchar(100) NULL,
  `u_email` varchar(100) NULL,
  `u_password` varchar(255) NULL,
  `u_role` enum('dosen','mahasiswa') NULL,
  `u_tanggal_lahir` date NULL,
  `u_no_ponsel` varchar(15) NULL,
  `u_alamat` varchar(255) NULL,
  `u_foto` varchar(255) NULL,
  PRIMARY KEY (`uid`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "absen_dosens" table
CREATE TABLE `absen_dosens` (
  `ad_id` varchar(255) NOT NULL,
  `ad_tanggal_dibuat` date NULL,
  `ad_deskripsi` longtext NULL,
  `ad_pertemuan` smallint NULL,
  `ad_kode` char(6) NULL,
  `kelas_k_id` varchar(255) NULL,
  `user_uid` varchar(255) NULL,
  PRIMARY KEY (`ad_id`),
  INDEX `fk_absen_dosens_kelas` (`kelas_k_id`),
  INDEX `fk_absen_dosens_user` (`user_uid`),
  CONSTRAINT `fk_absen_dosens_kelas` FOREIGN KEY (`kelas_k_id`) REFERENCES `kelas` (`k_id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `fk_absen_dosens_user` FOREIGN KEY (`user_uid`) REFERENCES `users` (`uid`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "absen_mahasiswas" table
CREATE TABLE `absen_mahasiswas` (
  `am_id` varchar(255) NOT NULL,
  `am_status` tinyint NULL,
  `am_deskripsi` longtext NULL,
  `absen_dosen_ad_id` varchar(255) NULL,
  `user_uid` varchar(255) NULL,
  `kelas_k_id` varchar(255) NULL,
  PRIMARY KEY (`am_id`),
  INDEX `fk_absen_mahasiswas_absen_dosen` (`absen_dosen_ad_id`),
  INDEX `fk_absen_mahasiswas_kelas` (`kelas_k_id`),
  INDEX `fk_absen_mahasiswas_user` (`user_uid`),
  CONSTRAINT `fk_absen_mahasiswas_absen_dosen` FOREIGN KEY (`absen_dosen_ad_id`) REFERENCES `absen_dosens` (`ad_id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `fk_absen_mahasiswas_kelas` FOREIGN KEY (`kelas_k_id`) REFERENCES `kelas` (`k_id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `fk_absen_mahasiswas_user` FOREIGN KEY (`user_uid`) REFERENCES `users` (`uid`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "tugas_dosens" table
CREATE TABLE `tugas_dosens` (
  `td_id` varchar(255) NOT NULL,
  `td_judul` varchar(100) NULL,
  `td_deskripsi` longtext NULL,
  `td_tanggal_dibuat` date NULL,
  `td_deadline` date NULL,
  `td_status` bool NULL,
  `td_file_soal` varchar(255) NULL,
  `kelas_k_id` varchar(255) NULL,
  `user_uid` varchar(255) NULL,
  PRIMARY KEY (`td_id`),
  INDEX `fk_tugas_dosens_kelas` (`kelas_k_id`),
  INDEX `fk_tugas_dosens_user` (`user_uid`),
  CONSTRAINT `fk_tugas_dosens_kelas` FOREIGN KEY (`kelas_k_id`) REFERENCES `kelas` (`k_id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `fk_tugas_dosens_user` FOREIGN KEY (`user_uid`) REFERENCES `users` (`uid`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "tugas_mahasiswas" table
CREATE TABLE `tugas_mahasiswas` (
  `tm_id` varchar(255) NOT NULL,
  `tm_waktu_pengumpulan` datetime NULL,
  `tm_status` bool NULL,
  `tm_file_tugas` varchar(255) NULL,
  `tm_nilai_tugas` bigint NULL,
  `tugas_dosen_td_id` varchar(255) NULL,
  `kelas_k_id` varchar(255) NULL,
  `user_uid` varchar(255) NULL,
  PRIMARY KEY (`tm_id`),
  INDEX `fk_tugas_mahasiswas_kelas` (`kelas_k_id`),
  INDEX `fk_tugas_mahasiswas_tugas_dosen` (`tugas_dosen_td_id`),
  INDEX `fk_tugas_mahasiswas_user` (`user_uid`),
  CONSTRAINT `fk_tugas_mahasiswas_kelas` FOREIGN KEY (`kelas_k_id`) REFERENCES `kelas` (`k_id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `fk_tugas_mahasiswas_tugas_dosen` FOREIGN KEY (`tugas_dosen_td_id`) REFERENCES `tugas_dosens` (`td_id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `fk_tugas_mahasiswas_user` FOREIGN KEY (`user_uid`) REFERENCES `users` (`uid`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "user_kelas" table
CREATE TABLE `user_kelas` (
  `kelas_k_id` varchar(255) NULL,
  `user_uid` varchar(255) NULL,
  INDEX `fk_user_kelas_kelas` (`kelas_k_id`),
  INDEX `fk_user_kelas_user` (`user_uid`),
  CONSTRAINT `fk_user_kelas_kelas` FOREIGN KEY (`kelas_k_id`) REFERENCES `kelas` (`k_id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `fk_user_kelas_user` FOREIGN KEY (`user_uid`) REFERENCES `users` (`uid`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
