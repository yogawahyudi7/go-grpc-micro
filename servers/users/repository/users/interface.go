package users

import entities_user "go-grpc-micro/servers/users/entities/users"

type UsersInterface interface {
	// ===============       gRPC       ===============//
	Create(name, image_url string) (entities_user.User, error)
	Read(after_id, condition_where, condition_order_sort string, page_size, offset_page int) ([]entities_user.User, error)
	Update(user_id, name, image_url string) (entities_user.User, error)
	Delete(user_id string) (entities_user.User, error)

	// =============== DATABASE GENERIC ===============//
	Get_one_by_user_id(user_id string) (entities_user.User, error)
	Get_total_by_condition(condition_where string) (int16, error)
}
