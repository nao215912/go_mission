CREATE TABLE IF NOT EXISTS `users` (
    `id` char(36) NOT NULL,
    `name` varchar(255) NOT NULL UNIQUE,
    `token` char(36) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `characters` (
    `id` char(36) NOT NULL,
    `name` varchar(255) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `user_characters` (
    `user_id` char(36) NOT NULL,
    `character_id` char(36) NOT NULL,
    CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    CONSTRAINT `fk_character_id` FOREIGN KEY (`character_id`) REFERENCES `characters` (`id`)
);

