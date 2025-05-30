CREATE TABLE articles (
    id UUID NOT NULL PRIMARY KEY,
    title  TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
)