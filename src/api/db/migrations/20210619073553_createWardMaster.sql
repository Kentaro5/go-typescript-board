
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `ward_mst` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `city_code` int(10) unsigned NOT NULL,
    `ward_code` int(10) unsigned NOT NULL,
    `ward` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `ward_mst_ward_code_index` (`ward_code`),
    KEY `ward_mst_city_code_foreign` (`city_code`),
    CONSTRAINT `ward_mst_city_code_foreign` FOREIGN KEY (`city_code`) REFERENCES `city_mst` (`city_code`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `ward_mst`;
