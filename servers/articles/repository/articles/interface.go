package articles

import entities_article "go-grpc-micro/servers/articles/entities/articles"

type ArticlesInterface interface {
	// ===============       gRPC       ===============//
	Create(user_id, title, content string) (entities_article.Article, error)
	Read(after_id, condition_where, condition_order_sort string, page_size, offset_page int) ([]entities_article.Article, error)
	Update(article_id, title, content string) (entities_article.Article, error)
	Delete(article_id string) (entities_article.Article, error)

	// =============== DATABASE GENERIC ===============//
	Get_one_by_article_id(article_id string) (entities_article.Article, error)
	Get_total_by_condition(condition_where string) (int16, error)
}
