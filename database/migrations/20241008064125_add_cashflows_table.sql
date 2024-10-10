-- +goose Up
-- +goose StatementBegin
CREATE TABLE cashflows (
    id SERIAL PRIMARY KEY,
    amount DECIMAL(15, 2) NOT NULL,
    type VARCHAR(6) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
