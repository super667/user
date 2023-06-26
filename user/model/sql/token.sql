CREATE TABLE `token` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT,
`user_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '用户名称',
`token`     varchar(2048)  NOT NULL DEFAULT '' COMMENT '刷新token',
`expired_at`    datetime     NULL COMMENT '过期时间',
`create_time`   datetime     NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_time`   datetime     NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
`delete_time`   datetime     NULL COMMENT '删除时间',
PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;