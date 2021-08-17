-- 租户
DROP TABLE IF EXISTS `tenant`;
CREATE TABLE `tenant` (
	`id` varchar(50) NOT NULL,
	`name` varchar(50) NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8;

INSERT INTO `tenant` (`id`,`name`) VALUES ('edf087cd-5e65-4a16-9f1d-2027776b757e', 'CompanyA');
INSERT INTO `tenant` (`id`,`name`) VALUES ('d6381a15-e2bd-4337-8f22-787dfb72ac09', 'CompanyB');

-- 菜单
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
	`id` varchar(50) NOT NULL,
	`name` varchar(50) NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8;

INSERT INTO `menu` (`id`,`name`) VALUES ('d85c19d9-8d0a-4f20-84f6-ec9d1449182d', 'EC2');
INSERT INTO `menu` (`id`,`name`) VALUES ('50ea4deb-4333-49e7-8d94-ee92818e2ce3', 'S3');