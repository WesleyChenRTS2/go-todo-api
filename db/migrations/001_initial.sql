-- +goose Up
CREATE TABLE IF NOT EXISTS todos (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  completed BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Seed the database with examples
INSERT INTO todos (title, completed) VALUES
  ('Example 1', FALSE),
  ('Example 2', TRUE),
  ('Example 3', FALSE);

-- +goose Down
DROP TABLE IF EXISTS todos;
