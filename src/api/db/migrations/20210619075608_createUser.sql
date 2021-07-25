
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `user_tbl` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `sex_code` tinyint(3) unsigned NOT NULL,
    `pref_code` int(10) unsigned NOT NULL,
    `city_code` int(10) unsigned NOT NULL,
    `ward_code` int(10) unsigned DEFAULT NULL,
    `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_tbl_email_unique` (`email`),
    KEY `user_tbl_sex_code_foreign` (`sex_code`),
    KEY `user_tbl_pref_code_foreign` (`pref_code`),
    KEY `user_tbl_city_code_foreign` (`city_code`),
    KEY `user_tbl_ward_code_foreign` (`ward_code`),
    CONSTRAINT `user_tbl_city_code_foreign` FOREIGN KEY (`city_code`) REFERENCES `city_mst` (`city_code`) ON UPDATE CASCADE,
    CONSTRAINT `user_tbl_pref_code_foreign` FOREIGN KEY (`pref_code`) REFERENCES `pref_mst` (`pref_code`) ON UPDATE CASCADE,
    CONSTRAINT `user_tbl_sex_code_foreign` FOREIGN KEY (`sex_code`) REFERENCES `sex_mst` (`code`) ON UPDATE CASCADE,
    CONSTRAINT `user_tbl_ward_code_foreign` FOREIGN KEY (`ward_code`) REFERENCES `ward_mst` (`ward_code`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `user_tbl`;
