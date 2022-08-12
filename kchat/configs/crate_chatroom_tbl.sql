CREATE TABLE IF NOT EXISTS `kchat_chatroom`(
   `room_id` INT UNSIGNED AUTO_INCREMENT,
   `name` VARCHAR(100) NOT NULL DEFAULT 'kchatuser',
   `users` TEXT NOT NULL,
   PRIMARY KEY ( `room_id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;