package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)


type DBConfig struct {
	Addr         string
	MaxIdleConns int
	MaxOpenConns int
	MaxIdleTime  string
}

func NewDBConn(cfg DBConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.Addr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	duration, err := time.ParseDuration(cfg.MaxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func NewPostgreStore(db *sql.DB) Storage {
	return &PostgreStore{db}
}

type PostgreStore struct {
	db *sql.DB
}


func (s *PostgreStore) Get(ctx context.Context, params GetUserParams) (*User, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	var column string
	switch params.QueryType {
	case Email:
		column = "email"
	case Username:
		column = "username"
	case Id:
		column = "id"
	default:
		panic("error expected one of the following: 'email', 'username', 'id', got: "+params.QueryType)
	}
	query := fmt.Sprintf(`
		SELECT id, first_name, last_name, username, email, password, birth_date, picture_path,
		       ST_Y(loc::geometry), ST_X(loc::geometry)
		FROM users
		WHERE %s = $1
	`, column)
 
	var u User
	var lastName sql.NullString
	var lat, lng sql.NullFloat64
 
	err := s.db.QueryRowContext(ctx, query, params.Val).Scan(
		&u.ID,
		&u.FirstName,
		&lastName,
		&u.Username,
		&u.Email,
		&u.Password.hash,
		&u.BirthDate,
		&u.PicturePath,
		&lat,
		&lng,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrResourceNotFound
		}
		return nil, err
	}
	u.LastName = lastName.String

	if lat.Valid && lng.Valid {
		u.Location = &Location{Lat: lat.Float64, Lng: lng.Float64}
	}
	return &u, nil
}

func (s *PostgreStore) Update(ctx context.Context, u *User) error {
	query := `
		UPDATE users
		SET first_name = $1,
		    last_name = $2,
		    username = $3,
		    email = $4,
		    password = $5,
		    birth_date = $6,
		    picture_path = $7,
		    loc = ST_GeogFromText($8)
		WHERE id = $9
	`
 
	res, err := s.db.ExecContext(ctx, query,
		u.FirstName,
		nullString(u.LastName),
		u.Username,
		u.Email,
		u.Password.hash,
		u.BirthDate,
		u.PicturePath,
		locationToEWKT(u.Location),
		u.ID,
	)
	if err != nil {
		return err
	}
 
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (s *PostgreStore) Create(ctx context.Context, u *User) error {
	query := `
		INSERT INTO users (first_name, last_name, username, email, password, birth_date, picture_path, loc)
		VALUES ($1, $2, $3, $4, $5, $6, $7, ST_GeogFromText($8)) RETURNING id
	`
 
	err := s.db.QueryRowContext(ctx, query,
		u.FirstName,
		nullString(u.LastName),
		u.Username,
		u.Email,
		u.Password.hash,
		u.BirthDate,
		u.PicturePath,
		locationToEWKT(u.Location),
	).Scan(
		&u.ID,
	)
	return err
}

func (s *PostgreStore) Delete(ctx context.Context, id string) error {
	res, err := s.db.ExecContext(ctx, `DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return err
	}
 
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
 
	return nil
}

func (s *PostgreStore) GetFeed(ctx context.Context, id string) ([]UserCards, error) {
	return nil, nil
}

