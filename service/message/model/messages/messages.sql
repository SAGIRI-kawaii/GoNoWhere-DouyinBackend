CREATE TABLE `messages` (
                            `id` bigint(20) NOT NULL AUTO_INCREMENT,
                            `user_id` bigint(20) DEFAULT NULL,
                            `to_user_id` bigint(20) DEFAULT NULL,
                            `content` text,
                            `create_at` datetime DEFAULT NULL,
                            `delete_at` datetime DEFAULT NULL,
                            `update_at` datetime DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            KEY `messages_id_index` (`id`),
                            KEY `messages_to_user_id_index` (`to_user_id`),
                            KEY `messages_user_id_index` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4

