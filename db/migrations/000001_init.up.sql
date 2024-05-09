-- users table
CREATE TABLE users (
id SERIAL PRIMARY KEY,
username VARCHAR(50) UNIQUE NOT NULL,
password VARCHAR(100) NOT NULL,
email VARCHAR(100) UNIQUE NOT NULL
);

-- groups table
CREATE TABLE groups (
id SERIAL PRIMARY KEY,
name VARCHAR(50) NOT NULL,
description TEXT
);

-- user_groups table (many-to-many relationship between users and groups)
CREATE TABLE user_groups (
user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
group_id INTEGER REFERENCES groups(id) ON DELETE CASCADE,
PRIMARY KEY (user_id, group_id)
);

-- tasks table
CREATE TABLE tasks (
id SERIAL PRIMARY KEY,
title VARCHAR(100) NOT NULL,
description TEXT,
due_date TIMESTAMP,
status VARCHAR(20) NOT NULL,
group_id INTEGER REFERENCES groups(id) ON DELETE CASCADE,
user_id INTEGER REFERENCES users(id) ON DELETE CASCADE
);

-- subtasks table
CREATE TABLE subtasks (
id SERIAL PRIMARY KEY,
title VARCHAR(100) NOT NULL,
description TEXT,
due_date TIMESTAMP,
status VARCHAR(20) NOT NULL,
task_id INTEGER REFERENCES tasks(id) ON DELETE CASCADE
);