CREATE TABLE IF NOT EXISTS todos(
    `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL,
    `name` VARCHAR (50),
    `description` VARCHAR (50),
    `is_done` BOOLEAN DEFAULT FALSE,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `deleted_at` DATETIME,
    FOREIGN KEY (`user_id`) REFERENCES users(`id`)
);
