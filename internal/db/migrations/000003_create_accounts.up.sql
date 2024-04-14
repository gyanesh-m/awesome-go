CREATE TABLE  IF NOT EXISTS accounts(
    id VARCHAR(255) PRIMARY KEY NOT NULL,
    amount BIGINT NOT NULL,
    document_number VARCHAR(255) NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    deleted_at BIGINT
);

