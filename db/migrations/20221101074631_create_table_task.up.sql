CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    image VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    parent_task_id INT DEFAULT NULL,
    poin INT NOT NULL);

-- migrate -database postgres://admin:admin123@localhost:5432/new-app -path db/migrations up
-- migrate -database "postgres://admin:admin123@localhost:5432/new-app?sslmode=disable" -path db/migrations up
-- migrate -database "postgres://postgres:admin123@localhost:2022/new-app?sslmode=disable" -path db/migrations up
