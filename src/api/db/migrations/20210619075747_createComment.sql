-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `comment_tbl` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` int(10) unsigned NOT NULL,
    `comment` text COLLATE utf8mb4_unicode_ci NOT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `comment_tbl_user_id_foreign` (`user_id`),
    CONSTRAINT `comment_tbl_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `user_tbl` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `comment_tbl`;

