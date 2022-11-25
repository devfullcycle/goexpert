-- name: ListCategories :many
SELECT * FROM categories;

-- name: GetCategory :one
SELECT * FROM categories 
WHERE id = ?;

-- name: CreateCategory :exec
INSERT INTO categories (id, name, description) 
VALUES (?,?,?);

-- name: UpdateCategory :exec
UPDATE categories SET name = ?, description = ? 
WHERE id = ?;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = ?;

-- name: CreateCourse :exec
INSERT INTO courses (id, name, description, category_id, price)
VALUES (?,?,?,?,?);

-- name: ListCourses :many
SELECT c.*, ca.name as category_name 
FROM courses c JOIN categories ca ON c.category_id = ca.id;