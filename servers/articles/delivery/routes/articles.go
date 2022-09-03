package routes

import (
	"context"
	article_ctrl "go-grpc-micro/servers/articles/delivery/controllers"
	client_user_ctrl "go-grpc-micro/servers/articles/delivery/controllers/external/users"
	pb_article "go-grpc-micro/servers/articles/protobuf_files/articles"
	helpers "go-grpc-micro/servers/articles/utils/generic"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (articleServer *ManagementServer) Create(ctx context.Context, in *pb_article.Request) (*pb_article.Response, error) {
	article_ctrl.Req_logger("Create", in.AfterId, in.Id, in.Content, in.Title, in.UserId, in.Order, in.Sort, in.Page)

	//reset
	article_ctrl.Response = &pb_article.Response{}
	helpers.UUID_Data = []helpers.UUID_Detail{}

	helpers.UUID_Data = append(helpers.UUID_Data, helpers.UUID_Detail{
		UUID: in.UserId,
		Name: "User",
	})

	err_uuid := helpers.IsValidUUID(helpers.UUID_Data)
	if err_uuid != nil {
		helpers.Errorf(err_uuid.Error())
		log.Print(&helpers.Buf)
		return article_ctrl.Response, status.Errorf(codes.InvalidArgument, err_uuid.Error())
	}

	//NOTE: Consume user data from user services
	res_user, err_user := client_user_ctrl.External_user(in.UserId)
	if err_user != nil {
		helpers.Errorf(err_user.Error())
		log.Print(&helpers.Buf)
		return article_ctrl.Response, err_user
	}

	err_title := helpers.IsValidContent(in.Title, "title")
	if err_title != nil {
		helpers.Errorf(err_title.Error())
		log.Print(&helpers.Buf)
		return article_ctrl.Response, status.Errorf(codes.InvalidArgument, err_title.Error())
	}

	err_content := helpers.IsValidContent(in.Content, "content")
	if err_content != nil {
		helpers.Errorf(err_content.Error())
		log.Print(&helpers.Buf)
		return article_ctrl.Response, status.Errorf(codes.InvalidArgument, err_content.Error())
	}

	res_article, err_article := articleServer.article_ctrl.Create(in.UserId, in.Title, in.Content)
	if err_article != nil {
		helpers.Errorf(err_article.Error())
		log.Print(&helpers.Buf)
		return article_ctrl.Response, status.Errorf(codes.Aborted, err_article.Error())
	}

	article_ctrl.Response.Meta = &pb_article.Meta{Total: 1, TotalPage: 1}
	article_ctrl.Response.Message = &pb_article.Message{Message: "Successfuly Operation"}
	article_ctrl.Response.Articles = []*pb_article.Data_Articles{{
		Id:      res_article.Id.String(),
		Title:   res_article.Title,
		Content: res_article.Content,
		User: &pb_article.Data_Users{
			Id:       res_user.Users[0].Id,
			Name:     res_user.Users[0].Name,
			ImageUrl: res_user.Users[0].ImageUrl,
		},
		CreatedAt: res_article.CreatedAt.Local().String(),
	}}

	return article_ctrl.Response, nil
}
