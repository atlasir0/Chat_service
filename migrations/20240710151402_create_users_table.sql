-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- TODO: раскидать по разным файлам 
CREATE TABLE users (
                       id BIGSERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       email VARCHAR(255) NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       role INT NOT NULL,
                       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS log_transaction(
    id INT GENERATED ALWAYS AS IDENTITY NOT NULL,
    info TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT pk_log_transaction_id PRIMARY KEY(id)
);




-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE users;

