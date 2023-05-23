CREATE TABLE `permission` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT,
`perm`          varchar(255) NOT NULL DEFAULT '' COMMENT '权限点',
`desc`         varchar(1024) NOT NULL DEFAULT '' COMMENT '权限点描述',
`create_time`   datetime NULL DEFAULT CURRENT_TIMESTAMP,
`update_time`   datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
`delete_time`   datetime NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;