// Package models contains the types for schema 'test'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Reaction represents a row from 'test.reactions'.
type Reaction struct {
	ID          int            `json:"id"`          // id
	Key         string         `json:"key"`         // key
	Icon        string         `json:"icon"`        // icon
	LongName    string         `json:"long_name"`   // long_name
	Description sql.NullString `json:"description"` // description
	DeletedAt   pq.NullTime    `json:"deleted_at"`  // deleted_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Reaction exists in the database.
func (r *Reaction) Exists() bool {
	return r._exists
}

// Deleted provides information if the Reaction has been deleted from the database.
func (r *Reaction) Deleted() bool {
	return r._deleted
}

// Insert inserts the Reaction to the database.
func (r *Reaction) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if r._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO test.reactions (` +
		`key, icon, long_name, description, deleted_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, r.Key, r.Icon, r.LongName, r.Description, r.DeletedAt)
	err = db.QueryRow(sqlstr, r.Key, r.Icon, r.LongName, r.Description, r.DeletedAt).Scan(&r.ID)
	if err != nil {
		return err
	}

	// set existence
	r._exists = true

	return nil
}

// Update updates the Reaction in the database.
func (r *Reaction) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !r._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if r._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE test.reactions SET (` +
		`key, icon, long_name, description, deleted_at` +
		`) = ( ` +
		`$1, $2, $3, $4, $5` +
		`) WHERE id = $6`

	// run query
	XOLog(sqlstr, r.Key, r.Icon, r.LongName, r.Description, r.DeletedAt, r.ID)
	_, err = db.Exec(sqlstr, r.Key, r.Icon, r.LongName, r.Description, r.DeletedAt, r.ID)
	return err
}

// Save saves the Reaction to the database.
func (r *Reaction) Save(db XODB) error {
	if r.Exists() {
		return r.Update(db)
	}

	return r.Insert(db)
}

// Upsert performs an upsert for Reaction.
//
// NOTE: PostgreSQL 9.5+ only
func (r *Reaction) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if r._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO test.reactions (` +
		`id, key, icon, long_name, description, deleted_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, key, icon, long_name, description, deleted_at` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.key, EXCLUDED.icon, EXCLUDED.long_name, EXCLUDED.description, EXCLUDED.deleted_at` +
		`)`

	// run query
	XOLog(sqlstr, r.ID, r.Key, r.Icon, r.LongName, r.Description, r.DeletedAt)
	_, err = db.Exec(sqlstr, r.ID, r.Key, r.Icon, r.LongName, r.Description, r.DeletedAt)
	if err != nil {
		return err
	}

	// set existence
	r._exists = true

	return nil
}

// Delete deletes the Reaction from the database.
func (r *Reaction) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !r._exists {
		return nil
	}

	// if deleted, bail
	if r._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM test.reactions WHERE id = $1`

	// run query
	XOLog(sqlstr, r.ID)
	_, err = db.Exec(sqlstr, r.ID)
	if err != nil {
		return err
	}

	// set deleted
	r._deleted = true

	return nil
}

// ReactionByKey retrieves a row from 'test.reactions' as a Reaction.
//
// Generated from index 'reactions_key_key'.
func ReactionByKey(db XODB, key string) (*Reaction, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, key, icon, long_name, description, deleted_at ` +
		`FROM test.reactions ` +
		`WHERE key = $1`

	// run query
	XOLog(sqlstr, key)
	r := Reaction{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, key).Scan(&r.ID, &r.Key, &r.Icon, &r.LongName, &r.Description, &r.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// ReactionByID retrieves a row from 'test.reactions' as a Reaction.
//
// Generated from index 'reactions_pkey'.
func ReactionByID(db XODB, id int) (*Reaction, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, key, icon, long_name, description, deleted_at ` +
		`FROM test.reactions ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	r := Reaction{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&r.ID, &r.Key, &r.Icon, &r.LongName, &r.Description, &r.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
