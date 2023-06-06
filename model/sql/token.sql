CREATE TABLE `token` (
`id` bigint unsigned NOT NULL AUTO_INCREMENT,
`user_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
`fresh_token`     varchar(32)  NOT NULL DEFAULT '' COMMENT '刷新token',
`expired_at`    int unsigned  NOT NULL DEFAULT 0 COMMENT '过期时间',
`create_time`   datetime     NULL DEFAULT CURRENT_TIMESTAMP,
`update_time`   datetime     NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
`delete_time`   datetime     NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;