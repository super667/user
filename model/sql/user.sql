CREATE TABLE `user` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT,
`number`        varchar(16) NOT NULL DEFAULT '' COMMENT '用户编号',
`name`          varchar(255) NOT NULL DEFAULT '' COMMENT '用户姓名',
`gender_code`   tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户性别',
`age`           int NOT NULL DEFAULT 0 COMMENT '年龄',
`dept_code`     varchar(32) NOT NULL DEFAULT '' COMMENT '部门编码',
`dept_name`     varchar(255) NOT NULL DEFAULT '' COMMENT '部门名称',
`manager_code`  varchar(32) NOT NULL DEFAULT '' COMMENT '管理者编码',
`manager_name`  varchar(255) NOT NULL DEFAULT '' COMMENT '管理者名称',
`phone`         varchar(32) NOT NULL DEFAULT '' COMMENT '电话号码',
`email`         varchar(32) NOT NULL DEFAULT '' COMMENT '邮箱',
`password`      varchar(128) NOT NULL DEFAULT '' COMMENT '用户密码',
`create_time`   datetime NULL DEFAULT CURRENT_TIMESTAMP,
`update_time`   datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
`delete_time`   datetime NULL,
PRIMARY KEY (`id`),
UNIQUE KEY `number` (`number`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;