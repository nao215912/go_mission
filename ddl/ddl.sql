CREATE TABLE IF NOT EXISTS `user` (
        `id` bigint(20) NOT NULL AUTO_INCREMENT,
        `name` varchar(255) NOT NULL UNIQUE,
        `token` varchar(255) NOT NULL,
        `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (`id`)
);
