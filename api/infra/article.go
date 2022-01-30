package infra

import (
	"fmt"
	"time"
	"yuki0920/go-notes/domain/model"
	"yuki0920/go-notes/domain/repository"
)

type ArticleRepository struct {
	SqlHandler
}

// repository.ArticleRepository のインターフェースを返すことのみを保証している
// つまり、インターフェースさえあれば、sqlHandler を使わず、DB に依存しないでモックすることが可能になる
func NewArticleRepository(sqlHandler SqlHandler) repository.ArticleRepository {
	return &ArticleRepository{
		SqlHandler: sqlHandler,
	}
}

func (articleRepository *ArticleRepository) GetById(id int) (*model.Article, error) {
	var article model.Article
	articleQuery := `SELECT *
	FROM articles
	WHERE id = ?;`
	if err := articleRepository.SqlHandler.Conn.Get(&article, articleQuery, id); err != nil {
		err = fmt.Errorf("failed to select article: %w", err)
		return nil, err
	}

	var categories []model.Category
	categoryQuery := `SELECT categories.*
	FROM article_categories
	INNER JOIN categories
	ON article_categories.category_id = categories.id
	WHERE article_id = ?;`
	if err := articleRepository.SqlHandler.Conn.Select(&categories, categoryQuery, id); err != nil {
		err = fmt.Errorf("failed to select categories: %w", err)
		return nil, err
	}
	article.Categories = categories

	return &article, nil
}

func (articleRepository *ArticleRepository) ListByPage(page int) ([]*model.Article, int, error) {
	selectQuery := `SELECT *
	FROM articles
	ORDER BY id desc
	LIMIT 5
	OFFSET ?`

	// クエリ結果を格納するスライスを初期化 5 件取得のため、サイズとキャパシティを指定
	articles := make([]*model.Article, 0, 5)
	offset := 5*(page-1) + 1
	if err := articleRepository.SqlHandler.Conn.Select(&articles, selectQuery, offset); err != nil {
		err = fmt.Errorf("failed to select articles: %w", err)
		return nil, 0, err
	}

	// TODO: N+1の解消
	for _, article := range articles {
		var categories []model.Category
		categoryQuery := `SELECT categories.*
		FROM article_categories
		INNER JOIN categories
		ON article_categories.category_id = categories.id
		WHERE article_id = ?;`
		if err := articleRepository.SqlHandler.Conn.Select(&categories, categoryQuery, article.ID); err != nil {
			err = fmt.Errorf("failed to select categories: %w", err)
			return nil, 0, err
		}
		article.Categories = categories
	}

	var count int
	countQuery := `SELECT
	COUNT(*)
	FROM articles
	`

	rows := articleRepository.SqlHandler.Conn.QueryRow(countQuery)
	err := rows.Scan(&count)
	if err != nil {
		err = fmt.Errorf("failed to count articles: %w", err)
		return nil, 0, err
	}

	totalPage := count/5 + 1

	return articles, totalPage, nil
}

func (articleRepository *ArticleRepository) Create(article *model.Article) (int64, error) {
	now := time.Now()

	article.Created = now
	article.Updated = now

	query := `INSERT INTO articles (title, body, created, updated)
	VALUES (:title, :body, :created, :updated);`

	// トランザクションを開始
	tx := articleRepository.SqlHandler.Conn.MustBegin()

	// クエリ文字列内の「:title」「:body」「:created」「:updated」は構造体の値で置換される
	// 構造体タグで指定してあるフィールドが対象となる ex)`db:"title"`
	res, err := tx.NamedExec(query, article)
	if err != nil {
		err = fmt.Errorf("failed to create article: %w", err)
		tx.Rollback()

		return 0, err
	}

	tx.Commit()

	id, err := res.LastInsertId()
	if err != nil {
		err = fmt.Errorf("failed to get last insert id: %w", err)
		return 0, err
	}

	return id, nil
}

func (articleRepository *ArticleRepository) Update(article *model.Article) error {
	now := time.Now()
	article.Updated = now

	query := `UPDATE articles
	SET title = :title,
			body = :body,
			updated = :updated
	WHERE id = :id;`

	tx := articleRepository.SqlHandler.Conn.MustBegin()

	// クエリ文字列内の :title, :body, :id には、第 2 引数の Article 構造体の Title, Body, ID が bind される
	if _, err := tx.NamedExec(query, article); err != nil {
		err = fmt.Errorf("failed to update article: %w", err)
		tx.Rollback()

		return err
	}

	tx.Commit()

	return nil
}

func (articleRepository *ArticleRepository) Delete(id int) error {
	query := `DELETE FROM articles WHERE id = ?;`

	tx := articleRepository.SqlHandler.Conn.MustBegin()

	if _, err := tx.Exec(query, id); err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit()
}

func (articleRepository *ArticleRepository) CreateCategories(articleID int, article *model.Article) error {
	for _, categoryId := range article.CategoryIDs {
		q := `INSERT INTO article_categories (article_id, category_id) VALUES (?, ?)`
		if _, err := articleRepository.SqlHandler.Conn.Exec(q, articleID, categoryId); err != nil {
			err = fmt.Errorf("failed to create article_category: %w", err)
			return err
		}
	}
	return nil
}

func (articleRepository *ArticleRepository) DeleteCategories(articleID int) error {
	fmt.Println("articleID:", articleID)
	q := `DELETE FROM article_categories WHERE article_id = ?;`
	if _, err := articleRepository.SqlHandler.Conn.Exec(q, articleID); err != nil {
		err = fmt.Errorf("failed to delete article_category: %w", err)
		return err
	}

	return nil
}
