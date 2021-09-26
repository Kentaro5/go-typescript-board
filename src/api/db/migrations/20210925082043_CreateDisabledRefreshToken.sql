
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `disabled_refresh_token_tbl` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `refresh_token` text NOT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='執行されたリフレッシュトークンのリストを保存するテーブル';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `disabled_refresh_token_tbl`;
