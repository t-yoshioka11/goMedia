package base

import (
	"fmt"
	"log"
)

// GetAllArticles 記事全件取得
func GetAllArticles() ([]Article, error) {
	db, err := gormConnect()
	if err != nil {
		log.Panicln("Action=GetAllArticles error=", err.Error())
		return nil, err
	}
	defer db.Close()

	articles := []Article{}
	db.Order("created_at desc").Find(&articles)

	return articles, nil
}

// GetArticleByID 記事詳細取得
func (article *Article) GetArticleByID() error {
	db, err := gormConnect()
	if err != nil {
		log.Panicln("Action=GetArticleDetails error=", err.Error())
		return err
	}
	defer db.Close()
	db.Where(article.Id).First(&article)
	return nil
}

// UpdateArticleByID 記事更新
func (article *Article) UpdateArticleByID() error {
	db, err := gormConnect()
	if err != nil {
		log.Panicln("Action=UpdateArticle error=", err.Error())
		return err
	}
	defer db.Close()

	//　更新処理
	db.Where(article.Id).Update(&article)
	return nil
}

// CerateArticle 記事作成処理
func (article *Article) CerateArticle() error {
	db, err := gormConnect()
	if err != nil {
		log.Panicln("Action=CerateArticle error=", err.Error())
		return err
	}
	defer db.Close()
	db.Create(&article)
	return nil
}

// DeleteArticleByID 記事削除
func (article *Article) DeleteArticleByID() error {
	db, err := gormConnect()
	if err != nil {
		log.Panicln("Action=DeleteArticle error=", err.Error())
		return err
	}
	defer db.Close()
	db.Where(article.Id).Delete(&article)

	return nil
}

// CreateUser ユーザ登録
func (user *User) CreateUser() error {
	db, err := gormConnect()
	if err != nil {
		log.Panicln("Action=CreateUser error=", err.Error())
		return err
	}
	defer db.Close()
	db.Create(&user)
	return nil
}

// SignUpUser ユーザ検索
func (user *User) SignUpUser() error {
	db, err := gormConnect()
	if err != nil {
		log.Panicln("Action=SignUpUser error=", err.Error())
		return err
	}
	defer db.Close()
	db.Find(&user)
	return nil
}

// CreateSessions セッション生成
func (sess *Sessions) CreateSesstion() error {
	db, err := gormConnect()
	if err != nil {
		log.Panicln("Action=CreateSesstion error=", err.Error())
		return err
	}
	defer db.Close()
	db.Create(&sess)
	return nil
}

// FindSession
func (sess *Sessions) FindSession() error {
	db, err := gormConnect()
	if err != nil {
		log.Panicln("Action=FindSession error=", err.Error())
		return err
	}
	defer db.Close()
	db.Find(&sess)
	return nil
}

// GetIDForCookie
func (sess *Sessions) GetIDForCookie() string {
	return fmt.Sprint(sess.Id)
}
