begin;

CREATE TABLE `groups` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `type` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE `users` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `username` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `group_id` INT NOT NULL,
    `first_name` VARCHAR(255) NOT NULL,
    `last_name` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

ALTER TABLE `users` ADD FOREIGN KEY (`group_id`) REFERENCES `groups`(`id`);

CREATE TABLE `menus` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `path` VARCHAR(255) NOT NULL,
    `parent_id` INT NOT NULL,
    `ordering` INT NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE `group_menus` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `group_id` INT NOT NULL,
    `menu_id` INT NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

ALTER TABLE `group_menus` ADD FOREIGN KEY (`group_id`) REFERENCES `groups`(`id`);
ALTER TABLE `group_menus` ADD FOREIGN KEY (`menu_id`) REFERENCES `menus`(`id`);

INSERT INTO `groups` (`name`, `type`) VALUES ('Admin', 'admin');
INSERT INTO `users` (`username`, `password`, `group_id`, `first_name`, `last_name`) VALUES ('admin', '$2y$10$HAfrrDzZXjDgy9XMf/eff.mSyYo2iRZlZ9LAsuXeYxc1cdoe4gxXS', (SELECT `id` FROM `groups` WHERE `name` = 'admin' LIMIT 1), 'Admin', 'Admin');
INSERT INTO `menus` (`name`, `path`, `parent_id`, `ordering`) VALUES ('Users', 'users', 0, 1);
INSERT INTO `group_menus` (`group_id`, `menu_id`) VALUES ((SELECT `id` FROM `groups` WHERE `name` = 'admin' LIMIT 1), (SELECT `id` FROM `menus` WHERE `name` = 'Users' LIMIT 1));


commit;
