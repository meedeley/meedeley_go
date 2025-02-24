CREATE SEQUENCE IF NOT EXISTS user_seq_id;
CREATE TABLE
    IF NOT EXISTS users (
        "id" INT PRIMARY KEY DEFAULT nextval('user_seq_id') NOT NULL,
        "name" VARCHAR(100) NOT NULL,
        "email" VARCHAR(100) NOT NULL UNIQUE,
        "password" VARCHAR(100) NOT NULL,
        "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMP WITH TIME ZONE
    )
