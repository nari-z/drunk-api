CREATE TABLE liquors (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL,
    image_file_path TEXT NOT NULL,
    created_at DATE NOT NULL,
    updated_at DATE NOT NULL,
    deleted_at DATE NOT NULL
);