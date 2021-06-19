
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `sex_mst` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `code` tinyint(3) unsigned NOT NULL,
    `name` varchar(4) COLLATE utf8mb4_unicode_ci NOT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `sex_mst_code_unique` (`code`),
    KEY `sex_mst_code_index` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `sex_mst`;
