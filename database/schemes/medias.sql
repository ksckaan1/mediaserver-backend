CREATE TABLE medias (
    id VARCHAR(50) PRIMARY KEY,
    created_at DATETIME NOT NULL,
    path TEXT NOT NULL,
    type TEXT NOT NULL,
    storage_type TEXT NOT NULL,
    mime_type TEXT NOT NULL,
    size INTEGER NOT NULL
)