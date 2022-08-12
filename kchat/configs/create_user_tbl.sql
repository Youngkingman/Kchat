-- USE kchat_service;
-- DROP TABLE `kchat_user`;
CREATE TABLE IF NOT EXISTS `kchat_user`(
   `uid` INT UNSIGNED AUTO_INCREMENT,
   `name` VARCHAR(100) NOT NULL DEFAULT 'kchatuser',
   `email` VARCHAR(40) NOT NULL UNIQUE,
   `password` VARCHAR(200) NOT NULL,
   `img_url` VARCHAR(200) NOT NULL DEFAULT '',
   `website` VARCHAR(200) NOT NULL DEFAULT '',
   PRIMARY KEY ( `uid` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;