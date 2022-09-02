package users

import (
	entities_user "go-grpc-micro/servers/users/entities/users"
	repo_user "go-grpc-micro/servers/users/repository/users"
	helpers "go-grpc-micro/servers/users/utils/generic"
	"log"
)

type UsersController struct {
	Repo repo_user.UsersInterface
}

func NewUsersControllers(usrep repo_user.UsersInterface) *UsersController {
	return &UsersController{Repo: usrep}
}

// ===============       gRPC       ===============//
func (uscon UsersController) Create(name, image_url string) (entities_user.User, error) {

	res, err := uscon.Repo.Create(name, image_url)
	if err != nil {
		helpers.Errorf(err.Error())
		log.Print(&helpers.Buf)
		return res, err
	}
	return res, nil
}

func (uscon UsersController) Read(after_id, condition_where, condition_order_sort string, page_size, offset_page int) ([]entities_user.User, error) {

	res, err := uscon.Repo.Read(after_id, condition_where, condition_order_sort, page_size, offset_page)
	if err != nil {
		helpers.Errorf(err.Error())
		log.Print(&helpers.Buf)
		return res, err
	}
	return res, nil
}

func (uscon UsersController) Update(user_id, name, image_url string) (entities_user.User, error) {

	res, err := uscon.Repo.Update(user_id, name, image_url)
	if err != nil {
		helpers.Errorf(err.Error())
		log.Print(&helpers.Buf)
		return res, err
	}
	return res, nil
}

func (uscon UsersController) Delete(user_id string) (entities_user.User, error) {

	res, err := uscon.Repo.Delete(user_id)
	if err != nil {
		helpers.Errorf(err.Error())
		log.Print(&helpers.Buf)
		return res, err
	}
	return res, nil
}

// ===============     GENERIC     ===============//
func (uscon UsersController) Get_one_by_user_id_Controller(user_id string) (entities_user.User, error) {

	res, err := uscon.Repo.Get_one_by_user_id(user_id)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (uscon UsersController) Get_total_by_condition_Controller(condition_where string) (int16, error) {

	res, err := uscon.Repo.Get_total_by_condition(condition_where)
	if err != nil {
		return res, err
	}
	return res, nil
}
