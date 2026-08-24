package service

import (
	"context"
	"time"

	pb "github.com/myselfBZ/uchrashuvda-isc/users"
	"github.com/myselfBZ/uchrashuvda-users/store"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type service struct {
	store store.Storage
	pb.UnimplementedUserServiceServer
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
// type UserServiceServer interface {
// 	Create(context.Context, *CreateUserRequest) (*User, error)
// 	GetByID(context.Context, *GetByIDRequest) (*User, error)
// 	GetByEmail(context.Context, *GetByEmailRequest) (*User, error)
// 	Verify(context.Context, *VerifyUserRequest) (*User, error)
// 	mustEmbedUnimplementedUserServiceServer()
// }

func (s *service) GetByEmail(ctx context.Context, r *pb.GetByEmailRequest) (*pb.User, error) {
	user, err := s.store.Get(ctx, store.GetUserParams{
		QueryType: store.Email,
		Val:       r.Email,
	})

	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:          user.ID.String(),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		BirthDate:   timestamppb.New(user.BirthDate),
		PicturePath: user.PicturePath,
		Email:       user.Email,
		Username:    user.Username,
		// Location: &pb.Location{
		// 	Lat: user.Location.Lat,
		// 	Lng: user.Location.Lng,
		// },
	}, nil
}

func (s *service) Verify(ctx context.Context, r *pb.VerifyUserRequest) (*pb.User, error) {
	user, err := s.store.Get(ctx, store.GetUserParams{
		QueryType: store.Email,
		Val:       r.Email,
	})

	if err != nil {
		return nil, err
	}

	if err := user.Password.Compare(r.Password); err != nil {
		return nil, err
	}

	return &pb.User{
		Id:          user.ID.String(),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		BirthDate:   timestamppb.New(user.BirthDate),
		PicturePath: user.PicturePath,
		Email:       user.Email,
		Username:    user.Username,
		// Location: &pb.Location{
		// 	Lat: user.Location.Lat,
		// 	Lng: user.Location.Lng,
		// },
	}, nil
}

func (s *service) Create(ctx context.Context, r *pb.CreateUserRequest) (*pb.User, error) {

	u := &store.User{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Username:  r.Username,
		Email:     r.Email,
		BirthDate: r.BirthDate.AsTime(),
	}

	if err := u.Password.Set(r.PasswordText); err != nil {
		return nil, err
	}

	if err := s.store.Create(ctx, u); err != nil {
		return nil, err
	}

	return &pb.User{
		Id:        u.ID.String(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		BirthDate: r.BirthDate,
	}, nil
}

func (s *service) GetByID(ctx context.Context, p *pb.GetByIDRequest) (*pb.User, error) {
	return &pb.User{
		Id:          "xxxx-xxx-xxx-xxxx",
		Username:    "myselfBZ",
		Email:       "boburforfun@gmail.com",
		BirthDate:   timestamppb.New(time.Now()),
		FirstName:   "bobur",
		LastName:    "alivobjonov",
		PicturePath: "/sexy/boys/picture",
	}, nil
}
