CREATE TABLE `friends`
(
    `id`         bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id`    bigint(20) DEFAULT NULL,
    `to_user_id` bigint(20) DEFAULT NULL,
    `create_at`  timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY          `friends_user_id_index` (`user_id`),
    KEY          `friends_to_user_id_index` (`to_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4

