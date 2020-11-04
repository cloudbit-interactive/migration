CREATE TABLE IF NOT EXISTS `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uuid` varchar(50) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `email` varchar(50) NOT NULL,
  `password` varchar(500) NOT NULL,
  `status` enum('ENABLED','DISABLED','CONFIRMING_EMAIL') NOT NULL DEFAULT 'ENABLED',
  `date` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `uuid` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `folder` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uuid` varchar(50) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `userUuid` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uuid` (`uuid`),
  UNIQUE KEY `name_userUuid` (`name`,`userUuid`),
  KEY `FK_folder_user` (`userUuid`),
  CONSTRAINT `FK_folder_user` FOREIGN KEY (`userUuid`) REFERENCES `user` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `text` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uuid` varchar(50) DEFAULT NULL,
  `title` varchar(500) DEFAULT NULL,
  `content` longtext DEFAULT NULL,
  `voice` text DEFAULT NULL,
  `date` timestamp NOT NULL DEFAULT current_timestamp(),
  `userUuid` varchar(50) DEFAULT NULL,
  `folderUuid` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uuid` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8mb4;