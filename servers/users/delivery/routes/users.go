package routes

import (
	"context"
	"fmt"
	user_ctrl "go-grpc-micro/servers/users/delivery/controllers"
	pb_user "go-grpc-micro/servers/users/protobuf_files/users"
	helpers "go-grpc-micro/servers/users/utils/generic"
	"log"
	"math"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (userServer *ManagementServer) Create(ctx context.Context, in *pb_user.Request) (*pb_user.Response, error) {
	user_ctrl.Req_logger("Create", in.AfterId, in.Id, in.ImageUrl, in.Name, in.Order, in.Sort, in.Page)

	//reset
	user_ctrl.Response = &pb_user.Response{}
	helpers.UUID_Data = []helpers.UUID_Detail{}

	err_name := helpers.IsValidContent(in.Name, "name")
	if err_name != nil {
		helpers.Errorf(err_name.Error())
		log.Print(&helpers.Buf)
		return user_ctrl.Response, status.Errorf(codes.InvalidArgument, err_name.Error())
	}

	res_user, err_user := userServer.user_ctrl.Create(in.Name, in.ImageUrl)
	if err_user != nil {
		helpers.Errorf(err_user.Error())
		log.Print(&helpers.Buf)
		return user_ctrl.Response, status.Errorf(codes.Aborted, err_user.Error())
	}

	user_ctrl.Response.Meta = &pb_user.Meta{Total: 1, TotalPage: 1}
	user_ctrl.Response.Message = &pb_user.Message{Message: "Successfuly Operation"}
	user_ctrl.Response.Users = []*pb_user.Data_Users{{
		Id:        res_user.Id.String(),
		Name:      res_user.Name,
		ImageUrl:  res_user.Image_Url,
		CreatedAt: res_user.CreatedAt.Local().String(),
	}}

	return user_ctrl.Response, nil
}

func (userServer *ManagementServer) Read(ctx context.Context, in *pb_user.Request) (*pb_user.Response, error) {
	user_ctrl.Req_logger("Read", in.AfterId, in.Id, in.ImageUrl, in.Name, in.Order, in.Sort, in.Page)

	//reset
	user_ctrl.Response = &pb_user.Response{}
	helpers.UUID_Data = []helpers.UUID_Detail{}

	page_size := 20
	offset_page := 0
	default_sort := "\"updated_at\""
	default_order := "DESC"
	condition_where, condition_order_sort := "", "u."+default_sort+" "+default_order

	if in.Page > 1 {
		offset_page = int((in.Page - 1) * int32(page_size))
	}

	if in.AfterId != "" {
		helpers.UUID_Data = append(helpers.UUID_Data, helpers.UUID_Detail{
			UUID: in.AfterId,
			Name: "After_id",
		})

		err_uuid := helpers.IsValidUUID(helpers.UUID_Data)
		if err_uuid != nil {
			helpers.Errorf(err_uuid.Error())
			log.Print(&helpers.Buf)
			return user_ctrl.Response, status.Errorf(codes.InvalidArgument, err_uuid.Error())
		}
	}

	if in.ImageUrl != "" {
		helpers.Errorf("image_url not support for read")
		log.Print(&helpers.Buf)
		return user_ctrl.Response, status.Errorf(codes.Canceled, "image_url not support for read")
	}

	// condition_order_sort
	if in.Sort != "" {
		condition_order_sort = strings.Replace(condition_order_sort, "u.\"updated_at\"", "r.\""+in.Sort+"\"", -1)
	}
	if in.Order != "" {
		condition_order_sort = strings.Replace(condition_order_sort, "DESC", strings.ToUpper(in.Order), -1)
	}

	// condition_where
	if in.Id != "" {
		helpers.UUID_Data = append(helpers.UUID_Data, helpers.UUID_Detail{
			UUID: in.Id,
			Name: "User",
		})

		err_uuid := helpers.IsValidUUID(helpers.UUID_Data)
		if err_uuid != nil {
			helpers.Errorf(err_uuid.Error())
			log.Print(&helpers.Buf)
			return user_ctrl.Response, status.Errorf(codes.InvalidArgument, err_uuid.Error())
		}
		condition_where = fmt.Sprint("u.id = '", in.Id, "'")
	}

	if in.Name != "" {
		condition_where = fmt.Sprint("u.name ilike '", in.Name, "'")
	}

	res_total_user, _ := userServer.user_ctrl.Get_total_by_condition_Controller(condition_where)
	res_user, err_user := userServer.user_ctrl.Read(in.AfterId, condition_where, condition_order_sort, page_size, offset_page)
	if err_user != nil {
		helpers.Errorf(err_user.Error())
		log.Print(&helpers.Buf)
		return user_ctrl.Response, status.Errorf(codes.NotFound, err_user.Error())
	}

	user_ctrl.Response.Message = &pb_user.Message{Message: "Successfully Operation"}
	user_ctrl.Response.Meta = &pb_user.Meta{Total: int32(res_total_user), TotalPage: int32(math.Ceil(float64(res_total_user) / float64(page_size)))}
	for i := 0; i < len(res_user); i++ {
		user_ctrl.Response.Users = append(user_ctrl.Response.Users, &pb_user.Data_Users{
			Id:        res_user[i].Id.String(),
			Name:      res_user[i].Name,
			ImageUrl:  res_user[i].Image_Url,
			CreatedAt: res_user[i].CreatedAt.Local().String(),
			UpdatedAt: res_user[i].UpdatedAt.Local().String(),
		})
	}

	return user_ctrl.Response, nil
}

func (userServer *ManagementServer) Update(ctx context.Context, in *pb_user.Request) (*pb_user.Response, error) {
	user_ctrl.Req_logger("Update", in.AfterId, in.Id, in.ImageUrl, in.Name, in.Order, in.Sort, in.Page)

	//reset
	user_ctrl.Response = &pb_user.Response{}
	helpers.UUID_Data = []helpers.UUID_Detail{}
	new_name, new_image_url := "", ""

	helpers.UUID_Data = append(helpers.UUID_Data, helpers.UUID_Detail{
		UUID: in.Id,
		Name: "User",
	})

	err_uuid := helpers.IsValidUUID(helpers.UUID_Data)
	if err_uuid != nil {
		helpers.Errorf(err_uuid.Error())
		log.Print(&helpers.Buf)
		return user_ctrl.Response, status.Errorf(codes.InvalidArgument, err_uuid.Error())
	}
	res_user, err_user := userServer.user_ctrl.Get_one_by_user_id_Controller(in.Id)
	if err_user != nil {
		helpers.Errorf(err_user.Error())
		log.Print(&helpers.Buf)
		return user_ctrl.Response, status.Errorf(codes.NotFound, err_user.Error())
	}
	new_name = res_user.Name
	new_image_url = res_user.Image_Url

	if in.Name != "" {
		new_name = in.Name
	}
	if in.ImageUrl != "" {
		new_image_url = in.ImageUrl
	}

	res_user_update, err_user := userServer.user_ctrl.Update(in.Id, new_name, new_image_url)
	if err_user != nil {
		helpers.Errorf(err_user.Error())
		log.Print(&helpers.Buf)
		return user_ctrl.Response, status.Errorf(codes.Aborted, err_user.Error())
	}

	user_ctrl.Response.Message = &pb_user.Message{Message: "Successfully Operation"}
	user_ctrl.Response.Meta = &pb_user.Meta{Total: 1, TotalPage: 1}
	user_ctrl.Response.Users = []*pb_user.Data_Users{{
		Id:        res_user_update.Id.String(),
		Name:      res_user_update.Name,
		ImageUrl:  res_user_update.Image_Url,
		CreatedAt: res_user_update.CreatedAt.Local().String(),
		UpdatedAt: res_user_update.UpdatedAt.Local().String(),
	}}

	return user_ctrl.Response, nil
}

func (userServer *ManagementServer) Delete(ctx context.Context, in *pb_user.Request) (*pb_user.Response, error) {
	user_ctrl.Req_logger("Delete", in.AfterId, in.Id, in.ImageUrl, in.Name, in.Order, in.Sort, in.Page)

	//reset
	user_ctrl.Response = &pb_user.Response{}
	helpers.UUID_Data = []helpers.UUID_Detail{}

	helpers.UUID_Data = append(helpers.UUID_Data, helpers.UUID_Detail{
		UUID: in.Id,
		Name: "User",
	})

	err_uuid := helpers.IsValidUUID(helpers.UUID_Data)
	if err_uuid != nil {
		helpers.Errorf(err_uuid.Error())
		log.Print(&helpers.Buf)
		return user_ctrl.Response, status.Errorf(codes.InvalidArgument, err_uuid.Error())
	}

	res_user, err_user := userServer.user_ctrl.Delete(in.Id)
	if err_user != nil {
		helpers.Errorf(err_user.Error())
		log.Print(&helpers.Buf)
		return user_ctrl.Response, status.Errorf(codes.Aborted, err_user.Error())
	}

	user_ctrl.Response.Message = &pb_user.Message{Message: "Successfully Operation"}
	user_ctrl.Response.Meta = &pb_user.Meta{Total: 1, TotalPage: 1}
	user_ctrl.Response.Users = []*pb_user.Data_Users{{
		Id:        res_user.Id.String(),
		Name:      res_user.Name,
		ImageUrl:  res_user.Image_Url,
		CreatedAt: res_user.CreatedAt.Local().String(),
		UpdatedAt: res_user.UpdatedAt.Local().String(),
		DeletedAt: res_user.DeletedAt.Time.Local().String(),
	}}

	return user_ctrl.Response, nil
}
