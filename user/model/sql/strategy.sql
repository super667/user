CREATE TABLE `strategy` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT,
`subject`         varchar(1024) NOT NULL DEFAULT '' COMMENT '权限赋予的主体',
`subject_name`         varchar(1024) NOT NULL DEFAULT '' COMMENT '权限赋予的主体名称',
`subject_type`         varchar(1024) NOT NULL DEFAULT '' COMMENT '主体类型',
`subject_type_name`         varchar(1024) NOT NULL DEFAULT '' COMMENT '主体类型名称',
`resource`         varchar(1024) NOT NULL DEFAULT '' COMMENT '资源',
`resource_name`         varchar(1024) NOT NULL DEFAULT '' COMMENT '资源名称',
`perm`          varchar(255) NOT NULL DEFAULT '' COMMENT '权限点',
`perm_name`          varchar(255) NOT NULL DEFAULT '' COMMENT '权限点名称',
`scope`         varchar(255) NOT NULL DEFAULT '' COMMENT '数据范围',
`create_time`   datetime NULL DEFAULT CURRENT_TIMESTAMP,
`update_time`   datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
`delete_time`   datetime NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;