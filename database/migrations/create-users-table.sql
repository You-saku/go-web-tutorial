CREATE TABLE IF NOT EXISTS users(
    `id` BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR (50),
    `email` VARCHAR (50),
    `age` INTEGER,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `deleted_at` DATETIME
);
