CREATE TABLE  IF NOT EXISTS  transactions(
    id VARCHAR(255) PRIMARY KEY NOT NULL,
    account_id VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    amount BIGINT NOT NULL,
    event_ts BIGINT NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    deleted_at BIGINT
);

