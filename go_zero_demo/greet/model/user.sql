CREATE TABLE `user` (
                            `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
                            `username` varchar(30) NOT NULL COMMENT '用户账号',
                            `password` varchar(100) NOT NULL DEFAULT '' COMMENT '密码',
                            `nickname` varchar(30) NOT NULL COMMENT '用户昵称',
                            `remark` varchar(500) DEFAULT NULL COMMENT '备注',
                            `dept_id` bigint(20) DEFAULT NULL COMMENT '部门ID',
                            `post_ids` varchar(255) DEFAULT NULL COMMENT '岗位编号数组',
                            `email` varchar(50) DEFAULT '' COMMENT '用户邮箱',
                            `mobile` varchar(11) DEFAULT '' COMMENT '手机号码',
                            `sex` tinyint(4) DEFAULT '0' COMMENT '用户性别',
                            `avatar` varchar(100) DEFAULT '' COMMENT '头像地址',
                            `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '帐号状态（0正常 1停用）',
                            `login_ip` varchar(50) DEFAULT '' COMMENT '最后登录IP',
                            `login_date` datetime DEFAULT NULL COMMENT '最后登录时间',
                            `creator` varchar(64) DEFAULT '' COMMENT '创建者',
                            `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `updater` varchar(64) DEFAULT '' COMMENT '更新者',
                            `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                            `deleted` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否删除',
                            PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表';