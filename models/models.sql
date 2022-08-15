CREATE TABLE `dag` (
                       `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'dag ID',
                       `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '名称',
                       `config` json DEFAULT NULL COMMENT 'dag 配置',
                       `status` int DEFAULT NULL COMMENT 'dag 状态',
                       `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                       `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                       PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='DAG表';

CREATE TABLE `processor` (
                             `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'processor ID',
                             `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '名称',
                             `handler` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '处理器(英文)',
                             `status` int DEFAULT NULL COMMENT '处理器 状态',
                             `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='处理器表';

CREATE TABLE `resource` (
                            `resource_id` bigint unsigned NOT NULL COMMENT '资源 id',
                            `dag_id` int NOT NULL COMMENT 'DAG id',
                            `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '名称',
                            `status` int DEFAULT NULL COMMENT '状态',
                            `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                            PRIMARY KEY (`resource_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='资源表';

CREATE TABLE `resource_state` (
                                  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
                                  `resource_id` bigint DEFAULT NULL COMMENT '资源 id',
                                  `processor_id` int DEFAULT NULL COMMENT '处理器 id',
                                  `process_state` int DEFAULT NULL COMMENT '处理状态',
                                  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='资源处理状态表';