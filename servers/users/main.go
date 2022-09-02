package main

import (
	configs "go-grpc-micro/servers/users/configs"
	user_ctrl "go-grpc-micro/servers/users/delivery/controlles"
	routes "go-grpc-micro/servers/users/delivery/routes"
	pb_user "go-grpc-micro/servers/users/protobuf_files/users"
	repo_user "go-grpc-micro/servers/users/repository/users"
	utils_database "go-grpc-micro/servers/users/utils/databases"
)

func main() {

	config := configs.GetConfig()
	db := utils_database.InitDB(config)

	repo_user := repo_user.NewUsersRepo(db)
	user_ctrl := user_ctrl.NewUsersControllers(repo_user)

	routes.NewManagementServer(
		db,
		&routes.NewManagementServer(db,
			&pb_user.UnimplementedUsersServer{},
			user_ctrl).
			UnimplementedUsersServer,
		user_ctrl).
		Run()
}
