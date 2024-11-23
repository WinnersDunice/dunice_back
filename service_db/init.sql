-- Создание таблицы users
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    name VARCHAR(100),
    surname VARCHAR(100),
    middlename VARCHAR(100),
    mac_address VARCHAR(17)
);

-- Создание таблицы offices
CREATE TABLE IF NOT EXISTS offices (
    officeid SERIAL PRIMARY KEY,
    address VARCHAR(255) NOT NULL
);

-- Создание промежуточной таблицы user_offices
CREATE TABLE IF NOT EXISTS user_offices (
    userid INT NOT NULL,
    officeid INT NOT NULL,
    PRIMARY KEY (userid, officeid),
    FOREIGN KEY (userid) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (officeid) REFERENCES offices(officeid) ON DELETE CASCADE
);

-- Создание таблицы admins
CREATE TABLE IF NOT EXISTS admins (
    userid INT NOT NULL,
    adminid SERIAL PRIMARY KEY,
    FOREIGN KEY (userid) REFERENCES users(id) ON DELETE CASCADE
);

-- Создание таблицы workspaces
CREATE TABLE IF NOT EXISTS workspaces (
    userid INT NOT NULL,
    workspaceid SERIAL PRIMARY KEY,
    FOREIGN KEY (userid) REFERENCES users(id) ON DELETE CASCADE
);
