CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    image VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    parent_task_id INT DEFAULT NULL,
    poin INT NOT NULL);