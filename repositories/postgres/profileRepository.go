package postgres

import (
	"github.com/bombergame/common/errs"
	"github.com/bombergame/profiles-service/domains"
	"strings"
)

type ProfileRepository struct {
	conn *Connection
}

func NewProfileRepository(conn *Connection) *ProfileRepository {
	return &ProfileRepository{
		conn: conn,
	}
}

func (r *ProfileRepository) Create(p domains.Profile) error {
	query := `SELECT * FROM create_profile($1,$2,$3,$4);`

	statement, err := r.conn.db.Prepare(query)
	if err != nil {
		return errs.NewServiceError(err)
	}

	_, err = statement.Exec(p.Username, p.Password.Hash, p.Password.Salt, p.Email)
	if err != nil {
		return wrapError(err)
	}

	return nil
}

func (r *ProfileRepository) FindByID(id int64) (*domains.Profile, error) {
	query := `SELECT * FROM get_profile($1);`

	statement, err := r.conn.db.Prepare(query)
	if err != nil {
		return nil, errs.NewServiceError(err)
	}

	row := statement.QueryRow(id)

	p := new(domains.Profile)
	if err := row.Scan(&p.ID, &p.Username, &p.Email, &p.Score); err != nil {
		return nil, wrapError(err)
	}

	return p, nil
}

func (r *ProfileRepository) FindByUsername(username string) (*domains.Profile, error) {
	query := `SELECT * FROM get_profile($1);`

	statement, err := r.conn.db.Prepare(query)
	if err != nil {
		return nil, errs.NewServiceError(err)
	}

	row := statement.QueryRow(username)

	p := new(domains.Profile)
	if err := row.Scan(&p.ID, &p.Password.Hash, &p.Password.Salt); err != nil {
		return nil, wrapError(err)
	}

	if err != nil {
		return nil, wrapError(err)
	}

	return p, nil
}

func (r *ProfileRepository) GetAllPaginated(pageIndex, pageSize int32) ([]domains.Profile, error) {
	if pageIndex < 1 {
		return nil, errs.NewInvalidFormatError("constraint violation: page index < 1")
	}
	if pageSize < 1 {
		return nil, errs.NewInvalidFormatError("constraint violation: page size < 1")
	}

	query := `SELECT * FROM get_all_profiles_paginated($1,$2);`

	statement, err := r.conn.db.Prepare(query)
	if err != nil {
		return nil, errs.NewServiceError(err)
	}

	rows, err := statement.Query(pageIndex-1, pageSize)
	if err != nil {
		return nil, errs.NewServiceError(err)
	}
	defer rows.Close()

	pf := make([]domains.Profile, 0)
	for rows.Next() {
		var p domains.Profile
		if err := rows.Scan(&p.ID, &p.Username, &p.Email, &p.Score); err != nil {
			return nil, wrapError(err)
		}

		if err != nil {
			return nil, errs.NewServiceError(err)
		}
		pf = append(pf, p)
	}

	return pf, nil
}

func (r *ProfileRepository) Update(id int64, p domains.Profile) error {
	query := `SELECT * FROM update_profile($1,$2,$3,$4,$5);`

	statement, err := r.conn.db.Prepare(query)
	if err != nil {
		return errs.NewServiceError(err)
	}

	_, err = statement.Exec(id, p.Username, p.Password.Hash, p.Password.Salt, p.Email)
	if err != nil {
		return wrapError(err)
	}

	return nil
}

func (r *ProfileRepository) Delete(id int64) error {
	query := `SELECT * FROM delete_profile($1);`

	statement, err := r.conn.db.Prepare(query)
	if err != nil {
		return errs.NewServiceError(err)
	}

	_, err = statement.Exec(id)
	if err != nil {
		return wrapError(err)
	}

	return nil
}

func parseProfile(scanFunc func(...interface{}) error) (*domains.Profile, error) {
	p := new(domains.Profile)
	if err := scanFunc(&p.ID, &p.Username, &p.Password, &p.Email, &p.Score); err != nil {
		return nil, err
	}
	return p, nil
}

func wrapError(err error) error {
	msg := err.Error()

	if strings.Contains(msg, "profile_username_unique") {
		return errs.NewDuplicateError("username duplicate")
	}
	if strings.Contains(msg, "profile_username_check") {
		return errs.NewInvalidFormatError("username pattern mismatch")
	}

	if strings.Contains(msg, "profile_password_check") {
		return errs.NewInvalidFormatError("password pattern mismatch")
	}

	if strings.Contains(msg, "profile_email_unique") {
		return errs.NewDuplicateError("email duplicate")
	}
	if strings.Contains(msg, "profile_email_check") {
		return errs.NewInvalidFormatError("email pattern mismatch")
	}

	if strings.Contains(msg, "not found") {
		return errs.NewNotFoundError("profile not found")
	}

	return errs.NewServiceError(err)
}
