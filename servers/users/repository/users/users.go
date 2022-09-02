package users

import (
	"errors"
	entities_user "go-grpc-micro/servers/users/entities/users"
	helpers "go-grpc-micro/servers/users/utils/generic"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

type Users_after_id struct {
	Updated_at string
}

// ===============       gRPC       ===============//
func (ur *UserRepository) Create(name, image_url string) (entities_user.User, error) {

	newUser := entities_user.User{
		Name:      name,
		Image_Url: image_url,
	}

	if err := ur.db.Save(&newUser).Error; err != nil {
		return newUser, errors.New("fail create user")
	}
	return newUser, nil
}

func (ur *UserRepository) Read(after_id, condition_where, condition_order_sort string, page_size, offset_page int) ([]entities_user.User, error) {
	users := []entities_user.User{}
	user_after_id := Users_after_id{}

	condition_after_id := time.Now().UnixNano() / int64(time.Millisecond)
	if after_id != "" {
		ur.db.
			Table("public.users u").
			Select("u.updated_at").
			Where("u.id = ?", after_id).
			Where("u.deleted_at is NULL").
			Limit(1).
			Scan(&user_after_id)
		ts, _ := time.Parse(time.RFC3339, user_after_id.Updated_at)
		if user_after_id.Updated_at != "" {
			condition_after_id = ts.Unix()
		}
	}

	query := `
		u.id,
		u.name,
		u.image_url,
		u.created_at,
		u.updated_at,
		u.deleted_at
	`

	if err := ur.db.
		Table("public.users u").
		Select(query).
		Where(condition_where).
		Where("extract(epoch FROM u.updated_at) < ?", condition_after_id).
		Where("u.deleted_at is NULL").
		Order(condition_order_sort).
		Limit(page_size).
		Offset(offset_page).
		Scan(&users).Error; err != nil || len(users) <= 0 {
		helpers.Errorf("user not found")
		log.Print(&helpers.Buf)
		return users, errors.New("user not found")
	}
	return users, nil
}

func (ur *UserRepository) Update(user_id, name, image_url string) (entities_user.User, error) {
	user := entities_user.User{
		Id: uuid.MustParse(user_id),
	}

	newUser := entities_user.User{
		Id:        uuid.MustParse(user_id),
		Name:      name,
		Image_Url: image_url,
	}

	if err := ur.db.Model(&user).Where(user).Updates(newUser).Error; err != nil {
		return newUser, errors.New("fail update users")
	}

	return newUser, nil
}

func (ur *UserRepository) Delete(user_id string) (entities_user.User, error) {
	user := entities_user.User{
		Id: uuid.MustParse(user_id),
	}
	if err := ur.db.First(&user, "id = ?", user_id).Error; err != nil {
		return user, errors.New("user not found")
	} else {
		if err := ur.db.Delete(&user).Error; err != nil {
			return user, errors.New("fail delete users")
		}
		return user, nil
	}
}

// =============== DATABASE GENERIC ===============//
func (ur *UserRepository) Get_one_by_user_id(user_id string) (entities_user.User, error) {
	user := entities_user.User{}

	query := `		
		u.id,		
		u.image_url,
		u.name,				
		u.created_at,
		u.updated_at,
		u.deleted_at
	`
	if err := ur.db.
		Table("public.users u").
		Select(query).
		Where("u.id = ?", user_id).
		Where("u.deleted_at is NULL").
		Limit(1).
		Offset(0).
		Scan(&user).Error; err != nil || user.Id == uuid.Nil {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (ur *UserRepository) Get_total_by_condition(condition_where string) (int16, error) {
	var total int16

	query := `	
		count(u.id) as "total"		
	`
	if err := ur.db.
		Table("public.users u").
		Select(query).
		Where(condition_where).
		Where("u.deleted_at is NULL").
		Scan(&total).Error; err != nil {
		return total, errors.New("user not found")
	}

	return total, nil
}
