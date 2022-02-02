CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE characters (
    "user" BIGINT PRIMARY KEY,
    "id" UUID DEFAULT uuid_generate_v4(),
    "name" VARCHAR(32) UNIQUE NOT NULL,
    "specks" BIGINT DEFAULT 50
);

-- INSERT INTO characters ("user", "name") VALUES (123456789, 'dom');