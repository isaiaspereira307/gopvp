-- name: CreateUser :exec
INSERT INTO users (id, first_name, last_name, email, password, is_active, is_staff, is_superuser, last_login, date_joined) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING *;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUsers :many
SELECT * FROM users;

-- name: UpdateUser :exec
UPDATE users SET first_name = $2, last_name = $3, password = $4, is_active = $5, is_staff = $6, is_superuser = $7, last_login = $8, date_joined = $9 WHERE id = $1 RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: CreateCategory :exec
INSERT INTO category (id, name) VALUES ($1, $2) RETURNING *;

-- name: GetCategoryById :one
SELECT * FROM category WHERE id = $1;

-- name: GetCategoryByName :one
SELECT * FROM category WHERE name = $1;

-- name: GetCategories :many
SELECT * FROM category;

-- name: UpdateCategory :exec
UPDATE category SET name = $2 WHERE id = $1 RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM category WHERE id = $1;

-- name: CreateSubcategory :exec
INSERT INTO subcategory (id, name) VALUES ($1, $2) RETURNING *;

-- name: GetSubcategoryById :one
SELECT * FROM subcategory WHERE id = $1;

-- name: GetSubcategoryByName :one
SELECT * FROM subcategory WHERE name = $1;

-- name: GetSubcategories :many
SELECT * FROM subcategory;

-- name: UpdateSubcategory :exec
UPDATE subcategory SET name = $2 WHERE id = $1 RETURNING *;

-- name: DeleteSubcategory :exec
DELETE FROM subcategory WHERE id = $1;

-- name: CreateObjective :exec
INSERT INTO objectives (id, user_id, category_id, subcategory_id, how_am_i, how_do_i_want_to_be, what_will_i_do, when_will_i_do_it) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: GetObjectiveById :one
SELECT * FROM objectives WHERE id = $1;

-- name: GetObjectives :many
SELECT * FROM objectives;

-- name: UpdateObjective :exec
UPDATE objectives SET user_id = $2, category_id = $3, subcategory_id = $4, how_am_i = $5, how_do_i_want_to_be = $6, what_will_i_do = $7, when_will_i_do_it = $8 WHERE id = $1 RETURNING *;

-- name: DeleteObjective :exec
DELETE FROM objectives WHERE id = $1;

-- name: GetObjectivesByUserId :many
SELECT * FROM objectives WHERE user_id = $1;

-- name: GetObjectivesByCategoryId :many
SELECT * FROM objectives WHERE category_id = $1;

-- name: GetObjectivesBySubcategoryId :many
SELECT * FROM objectives WHERE subcategory_id = $1;

-- name: GetObjectivesByUserIdAndCategoryId :many
SELECT * FROM objectives WHERE user_id = $1 AND category_id = $2;

-- name: GetObjectivesByUserIdAndSubcategoryId :many

SELECT * FROM objectives WHERE user_id = $1 AND subcategory_id = $2;

-- name: GetObjectivesByCategoryIdAndSubcategoryId :many
SELECT * FROM objectives WHERE category_id = $1 AND subcategory_id = $2;

-- name: GetObjectivesByUserIdAndCategoryIdAndSubcategoryId :many
SELECT * FROM objectives WHERE user_id = $1 AND category_id = $2 AND subcategory_id = $3;
