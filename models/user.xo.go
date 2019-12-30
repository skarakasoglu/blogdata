// Package models contains the types for schema 'test'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
)

// User represents a row from 'test.users'.
type User struct {
	ID           int            `json:"id"`             // id
	Username     string         `json:"username"`       // username
	Password     string         `json:"password"`       // password
	Name         string         `json:"name"`           // name
	Surname      string         `json:"surname"`        // surname
	Email        string         `json:"email"`          // email
	Description  sql.NullString `json:"description"`    // description
	CreatedAt    time.Time      `json:"created_at"`     // created_at
	UpdatedAt    time.Time      `json:"updated_at"`     // updated_at
	DeletedAt    pq.NullTime    `json:"deleted_at"`     // deleted_at
	CreateUserID sql.NullInt64  `json:"create_user_id"` // create_user_id
	UpdateUserID sql.NullInt64  `json:"update_user_id"` // update_user_id
	DeleteUserID sql.NullInt64  `json:"delete_user_id"` // delete_user_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the User exists in the database.
func (u *User) Exists() bool {
	return u._exists
}

// Deleted provides information if the User has been deleted from the database.
func (u *User) Deleted() bool {
	return u._deleted
}

// Insert inserts the User to the database.
func (u *User) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if u._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO test.users (` +
		`username, password, name, surname, email, description, created_at, updated_at, deleted_at, create_user_id, update_user_id, delete_user_id` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, u.Username, u.Password, u.Name, u.Surname, u.Email, u.Description, u.CreatedAt, u.UpdatedAt, u.DeletedAt, u.CreateUserID, u.UpdateUserID, u.DeleteUserID)
	err = db.QueryRow(sqlstr, u.Username, u.Password, u.Name, u.Surname, u.Email, u.Description, u.CreatedAt, u.UpdatedAt, u.DeletedAt, u.CreateUserID, u.UpdateUserID, u.DeleteUserID).Scan(&u.ID)
	if err != nil {
		return err
	}

	// set existence
	u._exists = true

	return nil
}

// Update updates the User in the database.
func (u *User) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if u._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE test.users SET (` +
		`username, password, name, surname, email, description, created_at, updated_at, deleted_at, create_user_id, update_user_id, delete_user_id` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`) WHERE id = $13`

	// run query
	XOLog(sqlstr, u.Username, u.Password, u.Name, u.Surname, u.Email, u.Description, u.CreatedAt, u.UpdatedAt, u.DeletedAt, u.CreateUserID, u.UpdateUserID, u.DeleteUserID, u.ID)
	_, err = db.Exec(sqlstr, u.Username, u.Password, u.Name, u.Surname, u.Email, u.Description, u.CreatedAt, u.UpdatedAt, u.DeletedAt, u.CreateUserID, u.UpdateUserID, u.DeleteUserID, u.ID)
	return err
}

// Save saves the User to the database.
func (u *User) Save(db XODB) error {
	if u.Exists() {
		return u.Update(db)
	}

	return u.Insert(db)
}

// Upsert performs an upsert for User.
//
// NOTE: PostgreSQL 9.5+ only
func (u *User) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if u._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO test.users (` +
		`id, username, password, name, surname, email, description, created_at, updated_at, deleted_at, create_user_id, update_user_id, delete_user_id` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, username, password, name, surname, email, description, created_at, updated_at, deleted_at, create_user_id, update_user_id, delete_user_id` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.username, EXCLUDED.password, EXCLUDED.name, EXCLUDED.surname, EXCLUDED.email, EXCLUDED.description, EXCLUDED.created_at, EXCLUDED.updated_at, EXCLUDED.deleted_at, EXCLUDED.create_user_id, EXCLUDED.update_user_id, EXCLUDED.delete_user_id` +
		`)`

	// run query
	XOLog(sqlstr, u.ID, u.Username, u.Password, u.Name, u.Surname, u.Email, u.Description, u.CreatedAt, u.UpdatedAt, u.DeletedAt, u.CreateUserID, u.UpdateUserID, u.DeleteUserID)
	_, err = db.Exec(sqlstr, u.ID, u.Username, u.Password, u.Name, u.Surname, u.Email, u.Description, u.CreatedAt, u.UpdatedAt, u.DeletedAt, u.CreateUserID, u.UpdateUserID, u.DeleteUserID)
	if err != nil {
		return err
	}

	// set existence
	u._exists = true

	return nil
}

// Delete deletes the User from the database.
func (u *User) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return nil
	}

	// if deleted, bail
	if u._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM test.users WHERE id = $1`

	// run query
	XOLog(sqlstr, u.ID)
	_, err = db.Exec(sqlstr, u.ID)
	if err != nil {
		return err
	}

	// set deleted
	u._deleted = true

	return nil
}

// UserByCreateUserID returns the User associated with the User's CreateUserID (create_user_id).
//
// Generated from foreign key 'users_create_user_id_fkey'.
func (u *User) UserByCreateUserID(db XODB) (*User, error) {
	return UserByID(db, int(u.CreateUserID.Int64))
}

// UserByDeleteUserID returns the User associated with the User's DeleteUserID (delete_user_id).
//
// Generated from foreign key 'users_delete_user_id_fkey'.
func (u *User) UserByDeleteUserID(db XODB) (*User, error) {
	return UserByID(db, int(u.DeleteUserID.Int64))
}

// UserByUpdateUserID returns the User associated with the User's UpdateUserID (update_user_id).
//
// Generated from foreign key 'users_update_user_id_fkey'.
func (u *User) UserByUpdateUserID(db XODB) (*User, error) {
	return UserByID(db, int(u.UpdateUserID.Int64))
}

// UserByID retrieves a row from 'test.users' as a User.
//
// Generated from index 'users_pkey'.
func UserByID(db XODB, id int) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, username, password, name, surname, email, description, created_at, updated_at, deleted_at, create_user_id, update_user_id, delete_user_id ` +
		`FROM test.users ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	u := User{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&u.ID, &u.Username, &u.Password, &u.Name, &u.Surname, &u.Email, &u.Description, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt, &u.CreateUserID, &u.UpdateUserID, &u.DeleteUserID)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// UserByUsername retrieves a row from 'test.users' as a User.
//
// Generated from index 'users_username_key'.
func UserByUsername(db XODB, username string) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, username, password, name, surname, email, description, created_at, updated_at, deleted_at, create_user_id, update_user_id, delete_user_id ` +
		`FROM test.users ` +
		`WHERE username = $1`

	// run query
	XOLog(sqlstr, username)
	u := User{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, username).Scan(&u.ID, &u.Username, &u.Password, &u.Name, &u.Surname, &u.Email, &u.Description, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt, &u.CreateUserID, &u.UpdateUserID, &u.DeleteUserID)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
