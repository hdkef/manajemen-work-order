package table

import "fmt"

var ENTITY_CREATION = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INT AUTO_INCREMENT PRIMARY KEY, fullname VARCHAR(50) NOT NULL, email VARCHAR(25) UNIQUE NOT NULL,username VARCHAR(10) UNIQUE NOT NULL, PASSWORD TEXT NOT NULL, ROLE ENUM('USER','BDMU','BDMUP','KELA','KELB','PPK','PPE','ULP', 'Super-Admin') NOT NULL, signature VARCHAR(50) NOT NULL UNIQUE)", ENTITY)

var PPP_CREATION = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INT AUTO_INCREMENT PRIMARY KEY, reason TEXT,creator_id INT NOT NULL, date_created DATE NOT NULL,doc VARCHAR(250) NOT NULL, photo VARCHAR(250), status VARCHAR(25) NOT NULL, perihal VARCHAR(25) NOT NULL, nota VARCHAR(50) NOT NULL UNIQUE, pekerjaan TEXT NOT NULL, sifat VARCHAR(15) NOT NULL ,bdmu_id INT, bdmup_id INT, kela_id INT, FOREIGN KEY (creator_id) REFERENCES entity(id))", PPP)

var RP_CREATION = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INT AUTO_INCREMENT PRIMARY KEY, reason TEXT,creator_id INT NOT NULL, date_created DATE NOT NULL, doc VARCHAR(250) NOT NULL, ppp_id INT NOT NULL, status VARCHAR(25) NOT NULL, kela_id INT, bdmup_id INT, bdmu_id INT, FOREIGN KEY (ppp_id) REFERENCES ppp(id))", RP)

var PERKIRAAN_BIAYA_CREATION = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INT AUTO_INCREMENT PRIMARY KEY, creator_id INT NOT NULL, date_created DATE NOT NULL, doc VARCHAR(250) NOT NULL, rp_id INT NOT NULL, est_cost DECIMAL(12,5) NOT NULL, FOREIGN KEY (rp_id) REFERENCES rp(id))", PERKIRAAN_BIAYA)

var PENGADAAN_CREATION = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INT AUTO_INCREMENT PRIMARY KEY, creator_id INT NOT NULL, date_created DATE NOT NULL, doc VARCHAR(250) NOT NULL, perkiraan_biaya_id INT NOT NULL, FOREIGN KEY (perkiraan_biaya_id) REFERENCES perkiraan_biaya(id))", PENGADAAN)

var SPK_CREATION = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INT AUTO_INCREMENT PRIMARY KEY, creator_id INT NOT NULL, date_created DATE NOT NULL, status VARCHAR(25) NOT NULL,doc VARCHAR(250) NOT NULL, pengadaan_id INT NOT NULL, worker_email VARCHAR(25) NOT NULL,FOREIGN KEY (pengadaan_id) REFERENCES pengadaan(id))", SPK)

var BDMU_PPP_CREATION = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INT AUTO_INCREMENT PRIMARY KEY, ppp_id INT NOT NULL, date_created DATE NOT NULL, FOREIGN KEY (ppp_id) REFERENCES ppp(id))", BDMU_PPP)

var BDMU_RP_CREATION = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INT AUTO_INCREMENT PRIMARY KEY, rp_id INT NOT NULL, date_created DATE NOT NULL, FOREIGN KEY (rp_id) REFERENCES rp(id))", BDMU_RP)

var BDMUP_PPP_CREATION = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INT AUTO_INCREMENT PRIMARY KEY, ppp_id INT NOT NULL, date_created DATE NOT NULL, FOREIGN KEY (ppp_id) REFERENCES ppp(id))", BDMUP_PPP)

var BDMUP_RP_CREATION = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INT AUTO_INCREMENT PRIMARY KEY, rp_id INT NOT NULL, date_created DATE NOT NULL, FOREIGN KEY (rp_id) REFERENCES rp(id))", BDMUP_RP)

var KELA_PPP_CREATION = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INT AUTO_INCREMENT PRIMARY KEY, ppp_id INT NOT NULL, date_created DATE NOT NULL, FOREIGN KEY (ppp_id) REFERENCES ppp(id))", KELA_PPP)

var KELA_RP_CREATION = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INT AUTO_INCREMENT PRIMARY KEY, rp_id INT NOT NULL, date_created DATE NOT NULL, FOREIGN KEY (rp_id) REFERENCES rp(id))", KELA_RP)

var KELB_PPP_CREATION = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INT AUTO_INCREMENT PRIMARY KEY, ppp_id INT NOT NULL, date_created DATE NOT NULL, FOREIGN KEY (ppp_id) REFERENCES ppp(id))", KELB_PPP)

var ULP_PERKIRAAN_BIAYA_CREATION = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INT AUTO_INCREMENT PRIMARY KEY, date_created DATE NOT NULL, perkiraan_biaya_id INT NOT NULL, FOREIGN KEY (perkiraan_biaya_id) REFERENCES perkiraan_biaya(id))", ULP_PERKIRAAN_BIAYA)

var PPE_PERKIRAAN_BIAYA_CREATION = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INT AUTO_INCREMENT PRIMARY KEY, date_created DATE NOT NULL, perkiraan_biaya_id INT NOT NULL, FOREIGN KEY (perkiraan_biaya_id) REFERENCES perkiraan_biaya(id))", PPE_PERKIRAAN_BIAYA)

var PPK_PENGADAAN_CREATION = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INT AUTO_INCREMENT PRIMARY KEY, date_created DATE NOT NULL, pengadaan_id INT NOT NULL, FOREIGN KEY (pengadaan_id) REFERENCES pengadaan(id))", PPK_PENGADAAN)

var PPK_RP_CREATION = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INT AUTO_INCREMENT PRIMARY KEY, rp_id INT NOT NULL, date_created DATE NOT NULL, FOREIGN KEY (rp_id) REFERENCES rp(id))", PPK_RP)

var PPK_SPK_CREATION = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INT AUTO_INCREMENT PRIMARY KEY, date_created DATE NOT NULL, spk_id INT NOT NULL, FOREIGN KEY (spk_id) REFERENCES spk(id))", PPK_SPK)

var EMAIL_SESSION_CREATION = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INT AUTO_INCREMENT PRIMARY KEY, PIN INT NOT NULL, spk_id INT NOT NULL, FOREIGN KEY(spk_id) REFERENCES spk(id))", EMAIL_SESSION)
