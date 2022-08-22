CREATE TABLE IF NOT EXISTS `kchat_chatroom`(
   `room_id` INT UNSIGNED AUTO_INCREMENT,
   `name` VARCHAR(100) NOT NULL DEFAULT 'kchatuser',
   `img_url` VARCHAR(200) NOT NULL DEFAULT '',
   `users` TEXT NOT NULL,
   PRIMARY KEY ( `room_id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
INSERT INTO `kchat_service`.`kchat_chatroom` (`room_id`, `name`, `img_url`, `users`) VALUES ('1', '环球加速翠翠钡盐团BFBuhnCuonCo', 'https://s3.bmp.ovh/imgs/2022/08/20/87e87448071da92a.jpg', '[1]');
INSERT INTO `kchat_service`.`kchat_chatroom` (`room_id`, `name`, `img_url`, `users`) VALUES ('2', 'TestRoom', 'https://s3.bmp.ovh/imgs/2022/08/20/fdf16a1aab38bd21.jpg', '[1]');