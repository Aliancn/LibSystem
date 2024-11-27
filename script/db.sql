

-- 检查并删除数据库
DO $$
    BEGIN
        -- 如果数据库 db_library 已存在，则删除它
        IF EXISTS (SELECT FROM pg_database WHERE datname = 'db_library') THEN
            PERFORM pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = 'db_library';
            EXECUTE 'DROP DATABASE db_library';
        END IF;
    END
$$;

-- 创建新的数据库
CREATE DATABASE db_library;

-- 连接到新创建的数据库
\c db_library


-- 删除并重新创建 user 表
DROP TABLE IF EXISTS "user";

CREATE TABLE "user" (
                        id SERIAL PRIMARY KEY,
                        username VARCHAR(255) UNIQUE NOT NULL,
                        password VARCHAR(255) NOT NULL,
                        name VARCHAR(255),
                        role VARCHAR(50) NOT NULL,
                        email VARCHAR(100),
                        phone VARCHAR(20),
                        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


-- 删除并重新创建 book 表
DROP TABLE IF EXISTS book;

CREATE TABLE book (
                      id SERIAL PRIMARY KEY,
                      title VARCHAR(255) NOT NULL,
                      author VARCHAR(255) NOT NULL,
                      publisher VARCHAR(255),
                      year INT NOT NULL,
                      genre VARCHAR(100),
                      status VARCHAR(50) NOT NULL, -- e.g., available, borrowed
                      location VARCHAR(100),
                      borrow_times INT NOT NULL,
                      created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                      updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


-- 删除并重新创建 borrow 表
DROP TABLE IF EXISTS borrow;

CREATE TABLE borrow (
                        id SERIAL PRIMARY KEY,
                        user_id INT NOT NULL,
                        book_id INT NOT NULL,
                        borrow_date DATE NOT NULL,
                        return_date DATE NOT NULL,
                        status VARCHAR(50) NOT NULL, -- e.g., borrowed, returned
                        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


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


-- 删除并重新创建 borrow 表
DROP TABLE IF EXISTS borrow;

CREATE TABLE borrow (
    id SERIAL PRIMARY KEY,
    borrow_date TIMESTAMP NOT NULL,
    return_date TIMESTAMP,
    status VARCHAR(50) NOT NULL, -- e.g., borrowed, returned
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    book_id INT NOT NULL,
    user_id INT NOT NULL
);

-- 删除并重新创建 info 表
DROP TABLE IF EXISTS info;
CREATE TABLE info (
    id SERIAL PRIMARY KEY,
    paper_id INT NOT NULL,
    download_time TIMESTAMP NOT NULL
);