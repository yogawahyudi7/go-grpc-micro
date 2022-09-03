package routes

import (
	configs "go-grpc-micro/servers/articles/configs"
	article_ctrl "go-grpc-micro/servers/articles/delivery/controllers"
	pb_article "go-grpc-micro/servers/articles/protobuf_files/articles"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type ManagementServer struct {
	dbconn *gorm.DB
	pb_article.UnimplementedArticlesServer
	article_ctrl *article_ctrl.ArticlesController
}

func NewManagementServer(
	db *gorm.DB,
	pb_article *pb_article.UnimplementedArticlesServer,
	article_ctrl *article_ctrl.ArticlesController,
) *ManagementServer {
	return &ManagementServer{
		dbconn:                      db,
		UnimplementedArticlesServer: *pb_article,
		article_ctrl:                article_ctrl,
	}
}

func (mServer *ManagementServer) Run() error {
	listen, err := net.Listen("tcp", ":"+configs.GetConfig().Port)
	if err != nil {
		log.Fatalln("error in listen Server", err.Error())
	}
	s := grpc.NewServer()
	pb_article.RegisterArticlesServer(s, mServer)
	log.Printf("Server listening at %v", listen.Addr())
	return s.Serve(listen)
}
