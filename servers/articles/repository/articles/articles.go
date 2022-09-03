package articles

import (
	"errors"
	entities_article "go-grpc-micro/servers/articles/entities/articles"
	helpers "go-grpc-micro/servers/articles/utils/generic"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticlesRepo(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

type Articles_after_id struct {
	Updated_at string
}

// ===============       gRPC       ===============//
func (ar *ArticleRepository) Create(user_id, title, content string) (entities_article.Article, error) {

	newArticle := entities_article.Article{
		UserId:  uuid.MustParse(user_id),
		Title:   title,
		Content: content,
	}

	if err := ar.db.Save(&newArticle).Error; err != nil {
		return newArticle, errors.New("fail create article")
	}
	return newArticle, nil
}

func (ar *ArticleRepository) Read(after_id, condition_where, condition_order_sort string, page_size, offset_page int) ([]entities_article.Article, error) {
	articles := []entities_article.Article{}
	article_after_id := Articles_after_id{}

	condition_after_id := time.Now().UnixNano() / int64(time.Millisecond)
	if after_id != "" {
		ar.db.
			Table("public.articles a").
			Select("a.updated_at").
			Where("a.id = ?", after_id).
			Where("a.deleted_at is NULL").
			Limit(1).
			Scan(&article_after_id)
		ts, _ := time.Parse(time.RFC3339, article_after_id.Updated_at)
		if article_after_id.Updated_at != "" {
			condition_after_id = ts.Unix()
		}
	}

	query := `
		a.id,
		a.title,
		a.content,
		a.created_at,
		a.updated_at,
		a.deleted_at
	`

	if err := ar.db.
		Table("public.articles a").
		Select(query).
		Where(condition_where).
		Where("extract(epoch FROM a.updated_at) < ?", condition_after_id).
		Where("a.deleted_at is NULL").
		Order(condition_order_sort).
		Limit(page_size).
		Offset(offset_page).
		Scan(&articles).Error; err != nil || len(articles) <= 0 {
		helpers.Errorf("article not found")
		log.Print(&helpers.Buf)
		return articles, errors.New("articles not found")
	}
	return articles, nil
}

func (ar *ArticleRepository) Update(article_id, title, content string) (entities_article.Article, error) {
	article := entities_article.Article{
		Id: uuid.MustParse(article_id),
	}

	newArticle := entities_article.Article{
		Id:      uuid.MustParse(article_id),
		Title:   title,
		Content: content,
	}

	if err := ar.db.Model(&article).Where(article).Updates(newArticle).Error; err != nil {
		return newArticle, errors.New("fail update articles")
	}

	return newArticle, nil
}

func (ar *ArticleRepository) Delete(article_id string) (entities_article.Article, error) {
	article := entities_article.Article{
		Id: uuid.MustParse(article_id),
	}
	if err := ar.db.First(&article, "id = ?", article_id).Error; err != nil {
		return article, errors.New("article not found")
	} else {
		if err := ar.db.Delete(&article).Error; err != nil {
			return article, errors.New("fail delete articles")
		}
		return article, nil
	}
}

// =============== DATABASE GENERIC ===============//
func (ar *ArticleRepository) Get_one_by_article_id(article_id string) (entities_article.Article, error) {
	article := entities_article.Article{}

	query := `
		a.id,
		a.title,
		a.content,
		a.created_at,
		a.updated_at,
		a.deleted_at
	`
	if err := ar.db.
		Table("public.articles a").
		Select(query).
		Where("a.id = ?", article_id).
		Where("a.deleted_at is NULL").
		Limit(1).
		Offset(0).
		Scan(&article).Error; err != nil || article.Id == uuid.Nil {
		return article, errors.New("article not found")
	}

	return article, nil
}

func (ar *ArticleRepository) Get_total_by_condition(condition_where string) (int16, error) {
	var total int16

	query := `
		count(a.id) as "total"
	`
	if err := ar.db.
		Table("public.articles a").
		Select(query).
		Where(condition_where).
		Where("a.deleted_at is NULL").
		Scan(&total).Error; err != nil {
		return total, errors.New("article not found")
	}

	return total, nil
}
