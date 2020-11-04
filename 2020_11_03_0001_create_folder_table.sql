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