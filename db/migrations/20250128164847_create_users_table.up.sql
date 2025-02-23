CREATE SEQUENCE IF NOT EXISTS user_sec_id;
CREATE TABLE
    IF NOT EXISTS users (
        "id" INT PRIMARY KEY DEFAULT nextval ('user_sec_id') NOT NULL,
        "name" VARCHAR(100) NOT NULL,
        "email" VARCHAR(100) NOT NULL,
        "password" VARCHAR(100) NOT NULL,
        "created_at" TIMESTAMP,
        "updated_at" TIMESTAMP
    )
