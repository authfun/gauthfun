
CREATE TABLE `tenant` (
	`id` varchar(50) NOT NULL,
	`name` varchar(50) NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8;

INSERT INTO `tenant` (`id`,`name`) VALUES ('edf087cd-5e65-4a16-9f1d-2027776b757e', 'AWS');
INSERT INTO `tenant` (`id`,`name`) VALUES ('d6381a15-e2bd-4337-8f22-787dfb72ac09', 'Aliyun');