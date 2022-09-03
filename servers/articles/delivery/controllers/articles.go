package articles

import (
	entities_article "go-grpc-micro/servers/articles/entities/articles"
	repo_article "go-grpc-micro/servers/articles/repository/articles"
	helpers "go-grpc-micro/servers/articles/utils/generic"
	"log"
)

type ArticlesController struct {
	Repo repo_article.ArticlesInterface
}

func NewArticlesControllers(arrep repo_article.ArticlesInterface) *ArticlesController {
	return &ArticlesController{Repo: arrep}
}

// ===============       gRPC       ===============//
func (arcon ArticlesController) Create(user_id, title, content string) (entities_article.Article, error) {

	res, err := arcon.Repo.Create(user_id, title, content)
	if err != nil {
		helpers.Errorf(err.Error())
		log.Print(&helpers.Buf)
		return res, err
	}
	return res, nil
}

func (arcon ArticlesController) Read(after_id, condition_where, condition_order_sort string, page_size, offset_page int) ([]entities_article.Article, error) {

	res, err := arcon.Repo.Read(after_id, condition_where, condition_order_sort, page_size, offset_page)
	if err != nil {
		helpers.Errorf(err.Error())
		log.Print(&helpers.Buf)
		return res, err
	}
	return res, nil
}

func (arcon ArticlesController) Update(article_id, title, content string) (entities_article.Article, error) {

	res, err := arcon.Repo.Update(article_id, title, content)
	if err != nil {
		helpers.Errorf(err.Error())
		log.Print(&helpers.Buf)
		return res, err
	}
	return res, nil
}

func (arcon ArticlesController) Delete(article_id string) (entities_article.Article, error) {

	res, err := arcon.Repo.Delete(article_id)
	if err != nil {
		helpers.Errorf(err.Error())
		log.Print(&helpers.Buf)
		return res, err
	}
	return res, nil
}

// ===============     GENERIC     ===============//
func (arcon ArticlesController) Get_one_by_article_id_Controller(article_id string) (entities_article.Article, error) {

	res, err := arcon.Repo.Get_one_by_article_id(article_id)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (arcon ArticlesController) Get_total_by_condition_Controller(condition_where string) (int16, error) {

	res, err := arcon.Repo.Get_total_by_condition(condition_where)
	if err != nil {
		return res, err
	}
	return res, nil
}
