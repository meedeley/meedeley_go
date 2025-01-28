CREATE SEQUENCE IF NOT EXISTS user_sec_id;

CREATE TABLE
    IF NOT EXIST users (
        "id" INT PRIMARY KEY DEFAULT nextval ("user_sec_id") NOT NULL,
        "name" VARCHAR(100),
        "email" VARCHAR(100),
        "password" VARCHAR(100),
        "created_at" TIMESTAMP,
        "updated_at" TIMESTAMP
    )