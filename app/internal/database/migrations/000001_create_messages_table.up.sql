CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    content VARCHAR(255) NOT NULL,
    recipient VARCHAR(255) NOT NULL,
    -- status INTEGER NOT NULL DEFAULT 0
    tries SMALLINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    sending_at TIMESTAMP DEFAULT NULL,
    sent_at TIMESTAMP DEFAULT NULL
);

CREATE INDEX idx_messages_created_at ON messages (created_at);
CREATE INDEX idx_messages_sending_at ON messages (sending_at);
