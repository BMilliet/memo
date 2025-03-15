-- +goose Up
ALTER TABLE snippets ADD COLUMN isTemporary BOOLEAN DEFAULT FALSE;

-- +goose Down
ALTER TABLE snippets DROP COLUMN isTemporary;

