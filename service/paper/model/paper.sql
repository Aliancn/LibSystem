-- 删除并重新创建 paper 表
DROP TABLE IF EXISTS paper;

CREATE TABLE paper (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    department VARCHAR(100),
    year INT NOT NULL,
    status VARCHAR(50) NOT NULL, -- e.g., available, archived
    download_times INT NOT NULL,
    file_path VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);