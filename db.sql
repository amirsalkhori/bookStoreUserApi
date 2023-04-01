DROP TABLE IF EXISTS users;
CREATE TABLE users (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NULL,
    `family` varchar(100) NULL,
    `email` varchar(100) NULL,
    `created_at` varchar(100) NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `email_UNIQUE` (`email` ASC ));