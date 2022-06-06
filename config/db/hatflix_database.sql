/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=1 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

SET foreign_key_checks = 1;

-- Dumping database structure for hatflix
CREATE DATABASE IF NOT EXISTS `hatflix` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `hatflix`;

-- Dumping structure for table hatflix.categorias
CREATE TABLE IF NOT EXISTS `categorias` (
  `id` INT(11) unsigned NOT NULL AUTO_INCREMENT,
  `nome` varchar(32) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- Dumping structure for table hatflix.lojas
CREATE TABLE IF NOT EXISTS `lojas` (
  `id` INT(11) unsigned NOT NULL AUTO_INCREMENT,
  `nome` varchar(32) NOT NULL,
  `cnpj` varchar(14) NOT NULL,
  `telefone` varchar(11) NOT NULL,
  `endereco` varchar(64) NOT NULL,
  `id_categoria` INT(11) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_categoria` (`id_categoria`),
  CONSTRAINT `loja_categoria` FOREIGN KEY (`id_categoria`) REFERENCES `categorias` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

# INSERT INTO `lojas` (`horario_abertura` ,`horario_fechamento` ,`dias_funcionamento` ,`id_cidade` ,`nome` ,`descricao` ,`telefone` ,`endereco`) VALUES
# 	("15:00:00", "21:30:00", 3, 2, "Restaurante do zé", "Melhor comida feita pelo Zé", "31985467513", "Rua das Pétalas, numero 12, bairro Família Sagrada"),
# 	("10:45:00", "16:00:00", 9, 3, "Maria das Massas", "Massas Artesanais", "33985467513", "Rua das Rosas, número 12, bairro Santa Família"),
#   ("09:00:00", "16:30:00", 55, 5, "Omelette du Frumage", "Restaurante Francês", "3182378172", "Rua das Orquídeas, número 332, bairro Sagrada Família");


-- Dumping data for table hatflix.categorias: ~0 rows (approximately)
/*!40000 ALTER TABLE `categorias` DISABLE KEYS */;
/*!40000 ALTER TABLE `categorias` ENABLE KEYS */;


-- COMENTADO POR NAO SER USADA EM V0


-- CREATE TABLE IF NOT EXISTS `restaurante-categoria` (
--   `id_restaurante` INT(11) unsigned NOT NULL,
--   `id_categoria` INT(11) unsigned NOT NULL,
--   PRIMARY KEY (`id_restaurante`, `id_categoria`),
--   CONSTRAINT `restaurante_fk` FOREIGN KEY (`id_restaurante`) REFERENCES `restaurantes` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
--   CONSTRAINT `categoria_fk` FOREIGN KEY (`id_categoria`) REFERENCES `categorias` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
-- ) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

-- Dumping structure for table hatflix.pratos
CREATE TABLE IF NOT EXISTS `produtos` (
  `id` INT(11) unsigned NOT NULL AUTO_INCREMENT,
  `nome` char(50) NOT NULL DEFAULT '0',
  `id_loja` INT(11) unsigned NOT NULL DEFAULT 0,
  `id_categoria` INT(11) unsigned DEFAULT NULL,
  `tamanho` varchar(8) NOT NULL,
  `preco` decimal(10,2) unsigned NOT NULL DEFAULT 0.00,
  `quantidade` INT(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `id_loja` (`id_loja`),
  KEY `id_categoria` (`id_categoria`),
  CONSTRAINT `pratos_categoria_fk` FOREIGN KEY (`id_categoria`) REFERENCES `categorias` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `pratos_loja_fk` FOREIGN KEY (`id_loja`) REFERENCES `lojas` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Dumping data for table hatflix.pratos: ~0 rows (approximately)
# /*!40000 ALTER TABLE `pratos` DISABLE KEYS */;
# INSERT INTO `produtos` (`id_restaurante`, `id_categoria`, `nome`, `preco`, `tempo_de_preparo`) VALUES
# 	(1, 1, 'Tilápia', 45.00, 30),
# 	(1, 3, 'Batata Frita', 15.00, 15),
# 	(1, 3, 'Batata Frita com Queijo', 18.00, 16),
# 	(2, 4, 'Iscas de frango acebolada', 20.00, 20),
# 	(2, 4, 'Filé Parmegiana', 40.00, 40),
#   (3, 5, 'Ratatouille', 65.00, 40),
#   (3, 5, 'Raclette', 110.00, 60);
# /*!40000 ALTER TABLE `pratos` ENABLE KEYS */;

-- Dumping data for table hatflix.restaurantes: ~0 rows (approximately)
# /*!40000 ALTER TABLE `restaurantes` DISABLE KEYS */;
# /*!40000 ALTER TABLE `restaurantes` ENABLE KEYS */;
#
# /*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
# /*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
# /*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;

# CREATE TABLE IF NOT EXISTS `usuarios` (
#   `id` INT(11) unsigned NOT NULL AUTO_INCREMENT,
#   `primeiro_nome` varchar(32) NOT NULL,
#   `ultimo_nome` varchar(32) NOT NULL,
#   `telefone` varchar(11) NOT NULL,
#   `email` varchar(50) NOT NULL UNIQUE,
#   `senha_hash` BINARY(64) NOT NULL,
#   PRIMARY KEY (`id`)
# ) ENGINE=InnoDB DEFAULT CHARSET=UTF8;
#
# INSERT INTO `usuarios` (`primeiro_nome`, `ultimo_nome`, `telefone`, `email`, `senha_hash`) VALUES
#   ('Joao', 'Pedro', '31984464729', 'joaopedro@gmail.com', SHA1('joaopedro2010')),
#   ('Maria', 'Lima', '35987432164', 'marialima@hotmail.com', SHA1('ml15122015')),
#   ('Carlos', 'Antunes', '37984455792', 'carlos_antunes12@outlook.com', SHA1('carlitos1212'));

