// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package db

import (
	"context"
	"time"
)

const createCategory = `-- name: CreateCategory :exec
INSERT INTO category (id, name) VALUES ($1, $2) RETURNING id, name
`

type CreateCategoryParams struct {
	ID   int32
	Name string
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, createCategory, arg.ID, arg.Name)
	return err
}

const createObjective = `-- name: CreateObjective :exec
INSERT INTO objectives (id, user_id, category_id, subcategory_id, how_am_i, how_do_i_want_to_be, what_will_i_do, when_will_i_do_it) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, user_id, category_id, subcategory_id, how_am_i, how_do_i_want_to_be, what_will_i_do, when_will_i_do_it
`

type CreateObjectiveParams struct {
	ID             int32
	UserID         int32
	CategoryID     int32
	SubcategoryID  int32
	HowAmI         string
	HowDoIWantToBe string
	WhatWillIDo    string
	WhenWillIDoIt  string
}

func (q *Queries) CreateObjective(ctx context.Context, arg CreateObjectiveParams) error {
	_, err := q.db.ExecContext(ctx, createObjective,
		arg.ID,
		arg.UserID,
		arg.CategoryID,
		arg.SubcategoryID,
		arg.HowAmI,
		arg.HowDoIWantToBe,
		arg.WhatWillIDo,
		arg.WhenWillIDoIt,
	)
	return err
}

const createSubcategory = `-- name: CreateSubcategory :exec
INSERT INTO subcategory (id, name) VALUES ($1, $2) RETURNING id, name
`

type CreateSubcategoryParams struct {
	ID   int32
	Name string
}

func (q *Queries) CreateSubcategory(ctx context.Context, arg CreateSubcategoryParams) error {
	_, err := q.db.ExecContext(ctx, createSubcategory, arg.ID, arg.Name)
	return err
}

const createUser = `-- name: CreateUser :exec
INSERT INTO users (id, first_name, last_name, email, password, is_active, is_staff, is_superuser, last_login, date_joined) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id, first_name, last_name, email, password, is_active, is_staff, is_superuser, last_login, date_joined
`

type CreateUserParams struct {
	ID          int32
	FirstName   string
	LastName    string
	Email       string
	Password    string
	IsActive    bool
	IsStaff     bool
	IsSuperuser bool
	LastLogin   time.Time
	DateJoined  time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Password,
		arg.IsActive,
		arg.IsStaff,
		arg.IsSuperuser,
		arg.LastLogin,
		arg.DateJoined,
	)
	return err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM category WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, id)
	return err
}

const deleteObjective = `-- name: DeleteObjective :exec
DELETE FROM objectives WHERE id = $1
`

func (q *Queries) DeleteObjective(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteObjective, id)
	return err
}

const deleteSubcategory = `-- name: DeleteSubcategory :exec
DELETE FROM subcategory WHERE id = $1
`

func (q *Queries) DeleteSubcategory(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteSubcategory, id)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getCategories = `-- name: GetCategories :many
SELECT id, name FROM category
`

func (q *Queries) GetCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, getCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCategoryById = `-- name: GetCategoryById :one
SELECT id, name FROM category WHERE id = $1
`

func (q *Queries) GetCategoryById(ctx context.Context, id int32) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategoryById, id)
	var i Category
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getCategoryByName = `-- name: GetCategoryByName :one
SELECT id, name FROM category WHERE name = $1
`

func (q *Queries) GetCategoryByName(ctx context.Context, name string) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategoryByName, name)
	var i Category
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getObjectiveById = `-- name: GetObjectiveById :one
SELECT id, user_id, category_id, subcategory_id, how_am_i, how_do_i_want_to_be, what_will_i_do, when_will_i_do_it FROM objectives WHERE id = $1
`

func (q *Queries) GetObjectiveById(ctx context.Context, id int32) (Objective, error) {
	row := q.db.QueryRowContext(ctx, getObjectiveById, id)
	var i Objective
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CategoryID,
		&i.SubcategoryID,
		&i.HowAmI,
		&i.HowDoIWantToBe,
		&i.WhatWillIDo,
		&i.WhenWillIDoIt,
	)
	return i, err
}

const getObjectives = `-- name: GetObjectives :many
SELECT id, user_id, category_id, subcategory_id, how_am_i, how_do_i_want_to_be, what_will_i_do, when_will_i_do_it FROM objectives
`

func (q *Queries) GetObjectives(ctx context.Context) ([]Objective, error) {
	rows, err := q.db.QueryContext(ctx, getObjectives)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Objective
	for rows.Next() {
		var i Objective
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CategoryID,
			&i.SubcategoryID,
			&i.HowAmI,
			&i.HowDoIWantToBe,
			&i.WhatWillIDo,
			&i.WhenWillIDoIt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getObjectivesByCategoryId = `-- name: GetObjectivesByCategoryId :many
SELECT id, user_id, category_id, subcategory_id, how_am_i, how_do_i_want_to_be, what_will_i_do, when_will_i_do_it FROM objectives WHERE category_id = $1
`

func (q *Queries) GetObjectivesByCategoryId(ctx context.Context, categoryID int32) ([]Objective, error) {
	rows, err := q.db.QueryContext(ctx, getObjectivesByCategoryId, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Objective
	for rows.Next() {
		var i Objective
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CategoryID,
			&i.SubcategoryID,
			&i.HowAmI,
			&i.HowDoIWantToBe,
			&i.WhatWillIDo,
			&i.WhenWillIDoIt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getObjectivesByCategoryIdAndSubcategoryId = `-- name: GetObjectivesByCategoryIdAndSubcategoryId :many
SELECT id, user_id, category_id, subcategory_id, how_am_i, how_do_i_want_to_be, what_will_i_do, when_will_i_do_it FROM objectives WHERE category_id = $1 AND subcategory_id = $2
`

type GetObjectivesByCategoryIdAndSubcategoryIdParams struct {
	CategoryID    int32
	SubcategoryID int32
}

func (q *Queries) GetObjectivesByCategoryIdAndSubcategoryId(ctx context.Context, arg GetObjectivesByCategoryIdAndSubcategoryIdParams) ([]Objective, error) {
	rows, err := q.db.QueryContext(ctx, getObjectivesByCategoryIdAndSubcategoryId, arg.CategoryID, arg.SubcategoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Objective
	for rows.Next() {
		var i Objective
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CategoryID,
			&i.SubcategoryID,
			&i.HowAmI,
			&i.HowDoIWantToBe,
			&i.WhatWillIDo,
			&i.WhenWillIDoIt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getObjectivesBySubcategoryId = `-- name: GetObjectivesBySubcategoryId :many
SELECT id, user_id, category_id, subcategory_id, how_am_i, how_do_i_want_to_be, what_will_i_do, when_will_i_do_it FROM objectives WHERE subcategory_id = $1
`

func (q *Queries) GetObjectivesBySubcategoryId(ctx context.Context, subcategoryID int32) ([]Objective, error) {
	rows, err := q.db.QueryContext(ctx, getObjectivesBySubcategoryId, subcategoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Objective
	for rows.Next() {
		var i Objective
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CategoryID,
			&i.SubcategoryID,
			&i.HowAmI,
			&i.HowDoIWantToBe,
			&i.WhatWillIDo,
			&i.WhenWillIDoIt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getObjectivesByUserId = `-- name: GetObjectivesByUserId :many
SELECT id, user_id, category_id, subcategory_id, how_am_i, how_do_i_want_to_be, what_will_i_do, when_will_i_do_it FROM objectives WHERE user_id = $1
`

func (q *Queries) GetObjectivesByUserId(ctx context.Context, userID int32) ([]Objective, error) {
	rows, err := q.db.QueryContext(ctx, getObjectivesByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Objective
	for rows.Next() {
		var i Objective
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CategoryID,
			&i.SubcategoryID,
			&i.HowAmI,
			&i.HowDoIWantToBe,
			&i.WhatWillIDo,
			&i.WhenWillIDoIt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getObjectivesByUserIdAndCategoryId = `-- name: GetObjectivesByUserIdAndCategoryId :many
SELECT id, user_id, category_id, subcategory_id, how_am_i, how_do_i_want_to_be, what_will_i_do, when_will_i_do_it FROM objectives WHERE user_id = $1 AND category_id = $2
`

type GetObjectivesByUserIdAndCategoryIdParams struct {
	UserID     int32
	CategoryID int32
}

func (q *Queries) GetObjectivesByUserIdAndCategoryId(ctx context.Context, arg GetObjectivesByUserIdAndCategoryIdParams) ([]Objective, error) {
	rows, err := q.db.QueryContext(ctx, getObjectivesByUserIdAndCategoryId, arg.UserID, arg.CategoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Objective
	for rows.Next() {
		var i Objective
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CategoryID,
			&i.SubcategoryID,
			&i.HowAmI,
			&i.HowDoIWantToBe,
			&i.WhatWillIDo,
			&i.WhenWillIDoIt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getObjectivesByUserIdAndCategoryIdAndSubcategoryId = `-- name: GetObjectivesByUserIdAndCategoryIdAndSubcategoryId :many
SELECT id, user_id, category_id, subcategory_id, how_am_i, how_do_i_want_to_be, what_will_i_do, when_will_i_do_it FROM objectives WHERE user_id = $1 AND category_id = $2 AND subcategory_id = $3
`

type GetObjectivesByUserIdAndCategoryIdAndSubcategoryIdParams struct {
	UserID        int32
	CategoryID    int32
	SubcategoryID int32
}

func (q *Queries) GetObjectivesByUserIdAndCategoryIdAndSubcategoryId(ctx context.Context, arg GetObjectivesByUserIdAndCategoryIdAndSubcategoryIdParams) ([]Objective, error) {
	rows, err := q.db.QueryContext(ctx, getObjectivesByUserIdAndCategoryIdAndSubcategoryId, arg.UserID, arg.CategoryID, arg.SubcategoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Objective
	for rows.Next() {
		var i Objective
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CategoryID,
			&i.SubcategoryID,
			&i.HowAmI,
			&i.HowDoIWantToBe,
			&i.WhatWillIDo,
			&i.WhenWillIDoIt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getObjectivesByUserIdAndSubcategoryId = `-- name: GetObjectivesByUserIdAndSubcategoryId :many

SELECT id, user_id, category_id, subcategory_id, how_am_i, how_do_i_want_to_be, what_will_i_do, when_will_i_do_it FROM objectives WHERE user_id = $1 AND subcategory_id = $2
`

type GetObjectivesByUserIdAndSubcategoryIdParams struct {
	UserID        int32
	SubcategoryID int32
}

func (q *Queries) GetObjectivesByUserIdAndSubcategoryId(ctx context.Context, arg GetObjectivesByUserIdAndSubcategoryIdParams) ([]Objective, error) {
	rows, err := q.db.QueryContext(ctx, getObjectivesByUserIdAndSubcategoryId, arg.UserID, arg.SubcategoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Objective
	for rows.Next() {
		var i Objective
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CategoryID,
			&i.SubcategoryID,
			&i.HowAmI,
			&i.HowDoIWantToBe,
			&i.WhatWillIDo,
			&i.WhenWillIDoIt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSubcategories = `-- name: GetSubcategories :many
SELECT id, name FROM subcategory
`

func (q *Queries) GetSubcategories(ctx context.Context) ([]Subcategory, error) {
	rows, err := q.db.QueryContext(ctx, getSubcategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Subcategory
	for rows.Next() {
		var i Subcategory
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSubcategoryById = `-- name: GetSubcategoryById :one
SELECT id, name FROM subcategory WHERE id = $1
`

func (q *Queries) GetSubcategoryById(ctx context.Context, id int32) (Subcategory, error) {
	row := q.db.QueryRowContext(ctx, getSubcategoryById, id)
	var i Subcategory
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getSubcategoryByName = `-- name: GetSubcategoryByName :one
SELECT id, name FROM subcategory WHERE name = $1
`

func (q *Queries) GetSubcategoryByName(ctx context.Context, name string) (Subcategory, error) {
	row := q.db.QueryRowContext(ctx, getSubcategoryByName, name)
	var i Subcategory
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, first_name, last_name, email, password, is_active, is_staff, is_superuser, last_login, date_joined FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.IsActive,
		&i.IsStaff,
		&i.IsSuperuser,
		&i.LastLogin,
		&i.DateJoined,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, first_name, last_name, email, password, is_active, is_staff, is_superuser, last_login, date_joined FROM users WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.IsActive,
		&i.IsStaff,
		&i.IsSuperuser,
		&i.LastLogin,
		&i.DateJoined,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, first_name, last_name, email, password, is_active, is_staff, is_superuser, last_login, date_joined FROM users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Password,
			&i.IsActive,
			&i.IsStaff,
			&i.IsSuperuser,
			&i.LastLogin,
			&i.DateJoined,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCategory = `-- name: UpdateCategory :exec
UPDATE category SET name = $2 WHERE id = $1 RETURNING id, name
`

type UpdateCategoryParams struct {
	ID   int32
	Name string
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, updateCategory, arg.ID, arg.Name)
	return err
}

const updateObjective = `-- name: UpdateObjective :exec
UPDATE objectives SET user_id = $2, category_id = $3, subcategory_id = $4, how_am_i = $5, how_do_i_want_to_be = $6, what_will_i_do = $7, when_will_i_do_it = $8 WHERE id = $1 RETURNING id, user_id, category_id, subcategory_id, how_am_i, how_do_i_want_to_be, what_will_i_do, when_will_i_do_it
`

type UpdateObjectiveParams struct {
	ID             int32
	UserID         int32
	CategoryID     int32
	SubcategoryID  int32
	HowAmI         string
	HowDoIWantToBe string
	WhatWillIDo    string
	WhenWillIDoIt  string
}

func (q *Queries) UpdateObjective(ctx context.Context, arg UpdateObjectiveParams) error {
	_, err := q.db.ExecContext(ctx, updateObjective,
		arg.ID,
		arg.UserID,
		arg.CategoryID,
		arg.SubcategoryID,
		arg.HowAmI,
		arg.HowDoIWantToBe,
		arg.WhatWillIDo,
		arg.WhenWillIDoIt,
	)
	return err
}

const updateSubcategory = `-- name: UpdateSubcategory :exec
UPDATE subcategory SET name = $2 WHERE id = $1 RETURNING id, name
`

type UpdateSubcategoryParams struct {
	ID   int32
	Name string
}

func (q *Queries) UpdateSubcategory(ctx context.Context, arg UpdateSubcategoryParams) error {
	_, err := q.db.ExecContext(ctx, updateSubcategory, arg.ID, arg.Name)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users SET first_name = $2, last_name = $3, password = $4, is_active = $5, is_staff = $6, is_superuser = $7, last_login = $8, date_joined = $9 WHERE id = $1 RETURNING id, first_name, last_name, email, password, is_active, is_staff, is_superuser, last_login, date_joined
`

type UpdateUserParams struct {
	ID          int32
	FirstName   string
	LastName    string
	Password    string
	IsActive    bool
	IsStaff     bool
	IsSuperuser bool
	LastLogin   time.Time
	DateJoined  time.Time
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Password,
		arg.IsActive,
		arg.IsStaff,
		arg.IsSuperuser,
		arg.LastLogin,
		arg.DateJoined,
	)
	return err
}
