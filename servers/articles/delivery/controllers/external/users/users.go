package users

import (
	"context"
	"fmt"
	configs "go-grpc-micro/servers/articles/configs"
	helpers "go-grpc-micro/servers/articles/utils/generic"
	pb_user "go-grpc-micro/servers/users/protobuf_files/users"
	"log"
	"time"

	"google.golang.org/grpc"
)

var Response *pb_user.Response = &pb_user.Response{}

func Get_data_user_by_user_id(client pb_user.UsersClient, user_id string) (*pb_user.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := pb_user.Request{Id: user_id}

	res_user, err_client := client.Read(ctx, &s)
	if err_client != nil {
		return res_user, err_client
	}

	return res_user, nil
}

func External_user(user_id string) (*pb_user.Response, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	url := fmt.Sprintf("%s:%s", configs.GetConfig().Address, helpers.Server_user_port)
	conn, err_dial := grpc.Dial(url, opts...)
	if err_dial != nil {
		helpers.Errorf(err_dial.Error())
		log.Print(&helpers.Buf)
		return Response, err_dial
	}
	defer conn.Close()
	client := pb_user.NewUsersClient(conn)

	res_user, err_user := Get_data_user_by_user_id(client, user_id)
	if err_user != nil {
		helpers.Errorf(err_user.Error())
		log.Print(&helpers.Buf)
		return Response, err_user
	}

	return res_user, nil

}
