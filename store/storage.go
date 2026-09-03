package store

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const(
	QueryTimeoutDuration = time.Second * 2
)

var(
	ErrResourceNotFound = errors.New("resource not found")
	ErrEmailIsTaken = errors.New("this email address is already taken")
	ErrUsernameIsTaken 	= errors.New("this username is already taken")
	ErrEmptyPassword = errors.New("empty password provided")
)

type GetUserParams struct {
	QueryType fetchUserBy
	Val       string
}

type fetchUserBy string

const (
	Email    fetchUserBy = "email"
	Username fetchUserBy = "username"
	Id       fetchUserBy = "id"
)

type User struct {
	ID          uuid.UUID 
	PicturePath string    
	Username    string   
	FirstName   string   
	LastName    string   
	Password    password  
	BirthDate   time.Time
	Location    *Location
	Email       string    
}


type password struct {
	text *string
	hash []byte
}

func (p *password) Set(text string) error {
	if text == "" {
		return ErrEmptyPassword 
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.text = &text
	p.hash = hash

	return nil
}

func (p *password) Compare(text string) error {
	return bcrypt.CompareHashAndPassword(p.hash, []byte(text))
}


type Location struct {
	Lat float64
	Lng float64
}

type UserCards struct {
	ID          uuid.UUID `json:"id"`
	PicturePath string    `json:"picture_path"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Age         int       `json:"age"`
	Distance    string    `json:"distance"`
}

type Storage interface {
	GetFeed(ctx context.Context, userId string) ([]UserCards, error)
	Get(ctx context.Context, params GetUserParams) (*User, error)
	Update(ctx context.Context, u *User) error
	Create(ctx context.Context, u *User) error
	Delete(ctx context.Context, id string) error
}
