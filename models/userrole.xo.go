// Package models contains the types for schema 'test'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
)

// UserRole represents a row from 'test.user_roles'.
type UserRole struct {
	ID           int           `json:"id"`             // id
	UserID       int           `json:"user_id"`        // user_id
	RoleID       int           `json:"role_id"`        // role_id
	CreatedAt    time.Time     `json:"created_at"`     // created_at
	UpdatedAt    time.Time     `json:"updated_at"`     // updated_at
	DeletedAt    pq.NullTime   `json:"deleted_at"`     // deleted_at
	CreateUserID sql.NullInt64 `json:"create_user_id"` // create_user_id
	UpdateUserID sql.NullInt64 `json:"update_user_id"` // update_user_id
	DeleteUserID sql.NullInt64 `json:"delete_user_id"` // delete_user_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the UserRole exists in the database.
func (ur *UserRole) Exists() bool {
	return ur._exists
}

// Deleted provides information if the UserRole has been deleted from the database.
func (ur *UserRole) Deleted() bool {
	return ur._deleted
}

// Insert inserts the UserRole to the database.
func (ur *UserRole) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if ur._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO test.user_roles (` +
		`user_id, role_id, created_at, updated_at, deleted_at, create_user_id, update_user_id, delete_user_id` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, ur.UserID, ur.RoleID, ur.CreatedAt, ur.UpdatedAt, ur.DeletedAt, ur.CreateUserID, ur.UpdateUserID, ur.DeleteUserID)
	err = db.QueryRow(sqlstr, ur.UserID, ur.RoleID, ur.CreatedAt, ur.UpdatedAt, ur.DeletedAt, ur.CreateUserID, ur.UpdateUserID, ur.DeleteUserID).Scan(&ur.ID)
	if err != nil {
		return err
	}

	// set existence
	ur._exists = true

	return nil
}

// Update updates the UserRole in the database.
func (ur *UserRole) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ur._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if ur._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE test.user_roles SET (` +
		`user_id, role_id, created_at, updated_at, deleted_at, create_user_id, update_user_id, delete_user_id` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) WHERE id = $9`

	// run query
	XOLog(sqlstr, ur.UserID, ur.RoleID, ur.CreatedAt, ur.UpdatedAt, ur.DeletedAt, ur.CreateUserID, ur.UpdateUserID, ur.DeleteUserID, ur.ID)
	_, err = db.Exec(sqlstr, ur.UserID, ur.RoleID, ur.CreatedAt, ur.UpdatedAt, ur.DeletedAt, ur.CreateUserID, ur.UpdateUserID, ur.DeleteUserID, ur.ID)
	return err
}

// Save saves the UserRole to the database.
func (ur *UserRole) Save(db XODB) error {
	if ur.Exists() {
		return ur.Update(db)
	}

	return ur.Insert(db)
}

// Upsert performs an upsert for UserRole.
//
// NOTE: PostgreSQL 9.5+ only
func (ur *UserRole) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if ur._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO test.user_roles (` +
		`id, user_id, role_id, created_at, updated_at, deleted_at, create_user_id, update_user_id, delete_user_id` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, user_id, role_id, created_at, updated_at, deleted_at, create_user_id, update_user_id, delete_user_id` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.user_id, EXCLUDED.role_id, EXCLUDED.created_at, EXCLUDED.updated_at, EXCLUDED.deleted_at, EXCLUDED.create_user_id, EXCLUDED.update_user_id, EXCLUDED.delete_user_id` +
		`)`

	// run query
	XOLog(sqlstr, ur.ID, ur.UserID, ur.RoleID, ur.CreatedAt, ur.UpdatedAt, ur.DeletedAt, ur.CreateUserID, ur.UpdateUserID, ur.DeleteUserID)
	_, err = db.Exec(sqlstr, ur.ID, ur.UserID, ur.RoleID, ur.CreatedAt, ur.UpdatedAt, ur.DeletedAt, ur.CreateUserID, ur.UpdateUserID, ur.DeleteUserID)
	if err != nil {
		return err
	}

	// set existence
	ur._exists = true

	return nil
}

// Delete deletes the UserRole from the database.
func (ur *UserRole) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ur._exists {
		return nil
	}

	// if deleted, bail
	if ur._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM test.user_roles WHERE id = $1`

	// run query
	XOLog(sqlstr, ur.ID)
	_, err = db.Exec(sqlstr, ur.ID)
	if err != nil {
		return err
	}

	// set deleted
	ur._deleted = true

	return nil
}

// UserByCreateUserID returns the User associated with the UserRole's CreateUserID (create_user_id).
//
// Generated from foreign key 'user_roles_create_user_id_fkey'.
func (ur *UserRole) UserByCreateUserID(db XODB) (*User, error) {
	return UserByID(db, int(ur.CreateUserID.Int64))
}

// UserByDeleteUserID returns the User associated with the UserRole's DeleteUserID (delete_user_id).
//
// Generated from foreign key 'user_roles_delete_user_id_fkey'.
func (ur *UserRole) UserByDeleteUserID(db XODB) (*User, error) {
	return UserByID(db, int(ur.DeleteUserID.Int64))
}

// Role returns the Role associated with the UserRole's RoleID (role_id).
//
// Generated from foreign key 'user_roles_role_id_fkey'.
func (ur *UserRole) Role(db XODB) (*Role, error) {
	return RoleByID(db, ur.RoleID)
}

// UserByUpdateUserID returns the User associated with the UserRole's UpdateUserID (update_user_id).
//
// Generated from foreign key 'user_roles_update_user_id_fkey'.
func (ur *UserRole) UserByUpdateUserID(db XODB) (*User, error) {
	return UserByID(db, int(ur.UpdateUserID.Int64))
}

// UserByUserID returns the User associated with the UserRole's UserID (user_id).
//
// Generated from foreign key 'user_roles_user_id_fkey'.
func (ur *UserRole) UserByUserID(db XODB) (*User, error) {
	return UserByID(db, ur.UserID)
}

// UserRoleByID retrieves a row from 'test.user_roles' as a UserRole.
//
// Generated from index 'user_roles_pkey'.
func UserRoleByID(db XODB, id int) (*UserRole, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, user_id, role_id, created_at, updated_at, deleted_at, create_user_id, update_user_id, delete_user_id ` +
		`FROM test.user_roles ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	ur := UserRole{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&ur.ID, &ur.UserID, &ur.RoleID, &ur.CreatedAt, &ur.UpdatedAt, &ur.DeletedAt, &ur.CreateUserID, &ur.UpdateUserID, &ur.DeleteUserID)
	if err != nil {
		return nil, err
	}

	return &ur, nil
}
