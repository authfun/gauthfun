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

-- Api
DROP TABLE IF EXISTS `api`;
CREATE TABLE `api` (
	`id` varchar(50) NOT NULL,
	`name` varchar(50) NULL,
	`group` varchar(50) NULL,
	`Method` varchar(10) NULL,
	`Route` varchar(10) NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8;

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

-- 组织架构
DROP TABLE IF EXISTS `organization`;
CREATE TABLE `organization` (
	`id` varchar(50) NOT NULL,
	`parent_id` varchar(50) NULL,
	`tenant_id` varchar(50) NOT NULL,
	`name` varchar(50) NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8;

INSERT INTO `organization` (`id`,`parent_id`,`tenant_id`,`name`) VALUES ("6bb9cc35-22b3-43e4-b03e-03c224a9d17c", NULL, "edf087cd-5e65-4a16-9f1d-2027776b757e", 'CompanyA');
INSERT INTO `organization` (`id`,`parent_id`,`tenant_id`,`name`) VALUES ("68b2b5b1-8cac-4db3-ae12-e3021ffffabb", "6bb9cc35-22b3-43e4-b03e-03c224a9d17c", "edf087cd-5e65-4a16-9f1d-2027776b757e", 'Department1');
INSERT INTO `organization` (`id`,`parent_id`,`tenant_id`,`name`) VALUES ("f4ab04b1-6247-445f-a358-945a16d246f2", "6bb9cc35-22b3-43e4-b03e-03c224a9d17c", "edf087cd-5e65-4a16-9f1d-2027776b757e", 'Department2');
INSERT INTO `organization` (`id`,`parent_id`,`tenant_id`,`name`) VALUES ("8f3f96ce-423a-4a50-b1ac-cd2c9e9e82bc", NULL, "d6381a15-e2bd-4337-8f22-787dfb72ac09", 'CompanyB');
INSERT INTO `organization` (`id`,`parent_id`,`tenant_id`,`name`) VALUES ("bbecf91c-bf30-4a19-bcc5-866be6898d11", "8f3f96ce-423a-4a50-b1ac-cd2c9e9e82bc", "d6381a15-e2bd-4337-8f22-787dfb72ac09", 'Department1');
INSERT INTO `organization` (`id`,`parent_id`,`tenant_id`,`name`) VALUES ("83e5a973-5981-4079-9ff7-1d47d6e9ae89", "8f3f96ce-423a-4a50-b1ac-cd2c9e9e82bc", "d6381a15-e2bd-4337-8f22-787dfb72ac09", 'Department2');