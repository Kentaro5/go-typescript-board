
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `pref_mst` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `pref_code` int(10) unsigned NOT NULL,
    `pref` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `pref_mst_pref_code_index` (`pref_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `pref_mst`;
