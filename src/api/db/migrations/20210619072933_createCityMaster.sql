
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `city_mst` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `pref_code` int(10) unsigned NOT NULL,
    `city_code` int(10) unsigned NOT NULL,
    `city` varchar(8) COLLATE utf8mb4_unicode_ci NOT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `city_mst_city_code_index` (`city_code`),
    KEY `city_mst_pref_code_foreign` (`pref_code`),
    CONSTRAINT `city_mst_pref_code_foreign` FOREIGN KEY (`pref_code`) REFERENCES `pref_mst` (`pref_code`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `city_mst`;
