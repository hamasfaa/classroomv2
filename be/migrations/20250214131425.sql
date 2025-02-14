-- Create "tugas_files" table
CREATE TABLE `tugas_files` (
  `tf_id` varchar(255) NOT NULL,
  `tf_nama` varchar(255) NULL,
  `tf_path` varchar(255) NULL,
  `tugas_td_id` varchar(255) NULL,
  PRIMARY KEY (`tf_id`),
  INDEX `fk_tugas_files_tugas_dosen` (`tugas_td_id`),
  CONSTRAINT `fk_tugas_files_tugas_dosen` FOREIGN KEY (`tugas_td_id`) REFERENCES `tugas_dosens` (`td_id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
