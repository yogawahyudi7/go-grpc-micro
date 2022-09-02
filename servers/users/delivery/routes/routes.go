package routes

import (
	configs "go-grpc-micro/servers/users/configs"
	user_ctrl "go-grpc-micro/servers/users/delivery/controlles"
	pb_user "go-grpc-micro/servers/users/protobuf_files/users"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type ManagementServer struct {
	dbconn *gorm.DB
	pb_user.UnimplementedUsersServer
	user_ctrl *user_ctrl.UsersController
}

func NewManagementServer(
	db *gorm.DB,
	pb_user *pb_user.UnimplementedUsersServer,
	user_ctrl *user_ctrl.UsersController,
) *ManagementServer {
	return &ManagementServer{
		dbconn:                   db,
		UnimplementedUsersServer: *pb_user,
		user_ctrl:                user_ctrl,
	}
}

func (mServer *ManagementServer) Run() error {
	listen, err := net.Listen("tcp", ":"+configs.GetConfig().Port)
	if err != nil {
		log.Fatalln("error in listen Server", err.Error())
	}
	s := grpc.NewServer()
	pb_user.RegisterUsersServer(s, mServer)
	log.Printf("Server listening at %v", listen.Addr())
	return s.Serve(listen)
}
