package main

import (
	configs "go-grpc-micro/servers/articles/configs"
	article_ctrl "go-grpc-micro/servers/articles/delivery/controllers"
	routes "go-grpc-micro/servers/articles/delivery/routes"
	pb_article "go-grpc-micro/servers/articles/protobuf_files/articles"
	repo_article "go-grpc-micro/servers/articles/repository/articles"
	utils_database "go-grpc-micro/servers/articles/utils/databases"
)

func main() {

	config := configs.GetConfig()
	db := utils_database.InitDB(config)

	repo_article := repo_article.NewArticlesRepo(db)
	article_ctrl := article_ctrl.NewArticlesControllers(repo_article)

	routes.NewManagementServer(
		db,
		&routes.NewManagementServer(db,
			&pb_article.UnimplementedArticlesServer{},
			article_ctrl).
			UnimplementedArticlesServer,
		article_ctrl).
		Run()
}
