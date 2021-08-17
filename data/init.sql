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

-- 功能
DROP TABLE IF EXISTS `feature`;
CREATE TABLE `feature` (
	`id` varchar(50) NOT NULL,
	`name` varchar(50) NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8;

INSERT INTO `feature` (`id`,`name`) VALUES ('f9781469-de65-4e68-8a0d-5592dbfcbe7a', 'EC2-Create');
INSERT INTO `feature` (`id`,`name`) VALUES ('0ba98484-5b11-4304-ba73-b1db422aeb8e', 'EC2-Query');
INSERT INTO `feature` (`id`,`name`) VALUES ('147b0e15-69b8-47ae-9f69-42540d0f4a6f', 'EC2-Update');
INSERT INTO `feature` (`id`,`name`) VALUES ('91ff2c24-a5c6-451f-b55e-3b8846ccb162', 'EC2-Delete');