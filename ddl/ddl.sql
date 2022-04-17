CREATE TABLE IF NOT EXISTS `users` (
        `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
        `name` varchar(255) NOT NULL UNIQUE,
        `token` varchar(255) NOT NULL,
        `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (`id`)
);
