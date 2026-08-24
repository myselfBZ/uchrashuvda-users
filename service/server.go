package service

import (
	"net"

	pb "github.com/myselfBZ/uchrashuvda-isc/users"
	"github.com/myselfBZ/user-service/store"
	"google.golang.org/grpc"
)


func New(s store.Storage) *Server {
	server := grpc.NewServer(
		grpc.MaxRecvMsgSize(10*1024*1024),
	)
	srv := &service{store: s}
	pb.RegisterUserServiceServer(server, srv)
	return &Server{
		service: srv,
		server: server,
	} 
}

type Server struct {
	service pb.UserServiceServer
	server *grpc.Server
}

func (s *Server) Run(port string) error {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return err
	} 
	return s.server.Serve(ln)
}
