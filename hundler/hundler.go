package hundler

import (
	"fmt"
	"goMedia/models/base"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

// TopPageHundler 記事一覧表示処理
func TopPageHundler(w http.ResponseWriter, r *http.Request) {
	user, err := authSession(r)
	if err != nil {
		log.Println(err.Error())
		//return
	}
	log.Println(user)

	articles, _ := base.GetAllArticles()
	generateHTML(w, articles, "layout", "public.navbar", "index")
}

// CreateArticleHundler 記事追加画面
func CreateArticleHundler(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public.navbar", "create")
}

// NewArticleHundler 記事投稿処理
func NewArticleHundler(w http.ResponseWriter, r *http.Request) {
	article := base.Article{Title: r.PostFormValue("title"), Body: r.PostFormValue("body")}
	err := article.CerateArticle()
	if err != nil {
		log.Println(err.Error())
		return
	}
	http.Redirect(w, r, "/top/", 302)
}

// GetArticleDetails 記事詳細取得ハンドラ
func GetArticleDetails(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	id, err := strconv.Atoi(vals.Get("id"))
	if err != nil {
		log.Println(err.Error())
		return
	}
	article := base.Article{Id: id}
	_ = article.GetArticleByID()
	log.Println(article)
	generateHTML(w, article, "layout", "public.navbar", "detail")
}

// UpdateArticleHundler 記事更新ハンドラ
func UpdateArticleHundler(w http.ResponseWriter, r *http.Request) {
	id := 4
	title := "更新タイトル"
	body := "テスト記事を更新しました。"

	article := base.Article{Id: id, Title: title, Body: body}
	article.UpdateArticleByID()
	log.Println("update", id)
}

// DeleteArticleHundler 記事削除ハンドラ
func DeleteArticleHundler(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	id, err := strconv.Atoi(vals.Get("id"))
	if err != nil {
		log.Println(err.Error())
		return
	}
	article := base.Article{Id: id}
	err = article.DeleteArticleByID()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("delete", id)
	http.Redirect(w, r, "/top/", 302)
}

// LoginPageHundler ログイン画面
func LoginPageHundler(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public.navbar", "login")
}

// SignUpPageHundler SignUp画面
func SignUpPageHundler(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public.navbar", "signup")
}

// SignUpHundler SignUP処理
func SignUpHundler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err, "Cannot parse form")
	}

	email := r.PostFormValue("email")
	pass := r.PostFormValue("password")

	//登録前チェック
	//  メールアドレス重複チェック
	//  パスワード文字数/ハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Failed to generate: %s", err)
		return
	}
	fmt.Println(hash)
	user := base.User{Email: email, Password: string(hash)}
	user.CreateUser()

	//実行後処理
	//  登録チェック 成功or失敗
	//  ログイン処理

	http.Redirect(w, r, "/top/", 302)
}

// SingInHundler サインイン処理
func SingInHundler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err, "Cannot parse form")
		return
	}

	email := r.PostFormValue("email")
	pass := r.PostFormValue("password")
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Failed to generate: %s", err)
		return
	}

	user := base.User{Email: email, Password: string(hash)}
	err = user.SignUpUser()
	if err != nil {
		// 	認証エラー
		log.Println(err.Error())
		return
	}
	// サインイン処理（cookiに情報を登録）
	sess := base.Sessions{Email: user.Email, UserId: user.UserId}
	err = sess.CreateSesstion()
	if err != nil {
		// err
		log.Println(err.Error())
		return
	}
	cookie := http.Cookie{Name: "goMedia_cookie",
		Value:    sess.GetIDForCookie(),
		Path:     "/",
		HttpOnly: true}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/top/", 302)
}

//////////////////
// privateメソッド
//////////////////

// generateHTML
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("view/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// parse HTML templates
// pass in a list of file names, and get a template
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("view/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

// authSession
func authSession(r *http.Request) (user base.User, err error) {
	cookie, err := r.Cookie("goMedia_cookie")
	if err != nil {
		// 	認証エラー
		log.Println(err.Error())
		return
	}
	sessID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Println(err.Error())
		return
	}
	sess := base.Sessions{Id: sessID}
	err = sess.FindSession()
	if err != nil {
		log.Println(err.Error())
		return
	}
	user.UserId = sess.UserId
	user.SignUpUser()
	return
}
