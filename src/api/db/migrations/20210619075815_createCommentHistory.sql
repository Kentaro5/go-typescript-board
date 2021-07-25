
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `comment_history_tbl` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `board_id` int(10) unsigned NOT NULL,
    `comment_id` int(10) unsigned NOT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `comment_history_tbl_board_id_foreign` (`board_id`),
    KEY `comment_history_tbl_comment_id_foreign` (`comment_id`),
    CONSTRAINT `comment_history_tbl_board_id_foreign` FOREIGN KEY (`board_id`) REFERENCES `board_tbl` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `comment_history_tbl_comment_id_foreign` FOREIGN KEY (`comment_id`) REFERENCES `comment_tbl` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `comment_history_tbl`;
