CREATE TABLE user_role (
   `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
   user_id INT NOT NULL DEFAULT 0 COMMENT '用户ID',
   role_id INT NOT NULL DEFAULT 0 COMMENT '角色ID',
   `desc` text COMMENT '描述',
   PRIMARY KEY(`id`),
   UNIQUE KEY `idx_role_user` (`user_id`, `role_id`)
);
