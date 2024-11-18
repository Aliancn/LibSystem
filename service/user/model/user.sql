DROP TABLE IF EXISTS "user";

CREATE TABLE "user" (
                        id SERIAL PRIMARY KEY,
                        username VARCHAR(255) UNIQUE NOT NULL,
                        password VARCHAR(255) NOT NULL,
                        name VARCHAR(255),
                        role VARCHAR(50) NOT NULL DEFAULT 'user',
                        email VARCHAR(100),
                        phone VARCHAR(20),
                        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);