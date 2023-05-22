CREATE TABLE `user` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT,
`number`        varchar(255) NOT NULL DEFAULT '' COMMENT '用户编号',
`name`          varchar(255) NOT NULL DEFAULT '' COMMENT '用户姓名',
`gender_code`   tinyint(3) UNSIGNED NULL DEFAULT NULL COMMENT '用户性别',
`age`           int NULL DEFAULT NULL COMMENT '年龄',
`dept_code`     varchar(255) NULL DEFAULT NULL COMMENT '部门编码',
`dept_name`     varchar(255) NULL DEFAULT NULL COMMENT '部门名称',
`manager_code`  varchar(255) NULL DEFAULT NULL COMMENT '管理者编码',
`manager_name`  varchar(255) NULL DEFAULT NULL COMMENT '管理者名称',
`phone`         varchar(255) NULL DEFAULT NULL COMMENT '电话号码',
`email`         varchar(255) NULL DEFAULT NULL COMMENT '邮箱',
`password`      varchar(255) NULL DEFAULT NULL COMMENT '用户密码',
`create_time`   datetime NULL DEFAULT CURRENT_TIMESTAMP,
`update_time`   datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
`delete_time`   datetime NULL,
PRIMARY KEY (`id`),
UNIQUE KEY `number` (`number`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;