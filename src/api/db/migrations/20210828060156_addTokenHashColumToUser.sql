
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE user_tbl
    ADD token_hash VARCHAR(15) NOT NULL
    AFTER password;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE user_tbl
DROP COLUMN token_hash;
