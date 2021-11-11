-- +goose Up
ALTER TABLE artists ALTER COLUMN name SET NOT NULL;
ALTER TABLE artists ADD CONSTRAINT namechk CHECK (name != '');
-- +goose Down
ALTER TABLE artists ALTER COLUMN name DROP NOT NULL;
ALTER TABLE artists DROP CONSTRAINT namechk;
