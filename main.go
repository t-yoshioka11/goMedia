package main

import (
	"goMedia/config"
	"goMedia/hundler"
	"goMedia/utils"
	"net/http"
)

// メイン処理
func main() {
	// 初期設定
	utils.LoggingSettings(config.Config.LogFile)

	http.HandleFunc("/test/", hundler.TopPageHundler)

	// ハンドラ追加
	http.HandleFunc("/top/", hundler.TopPageHundler)
	http.HandleFunc("/create/", hundler.CreateArticleHundler)
	http.HandleFunc("/detail/", hundler.GetArticleDetails)
	http.HandleFunc("/edit/", hundler.UpdateArticleHundler)
	http.HandleFunc("/delete/", hundler.DeleteArticleHundler)
	http.HandleFunc("/login/", hundler.LoginPageHundler)
	http.HandleFunc("/login/sign_in", hundler.SingInHundler)
	http.HandleFunc("/sign_up/", hundler.SignUpPageHundler)

	http.HandleFunc("/create/new/", hundler.NewArticleHundler)
	http.HandleFunc("/sign_up/signup_account/", hundler.SignUpHundler)

	// サーバ起動
	http.ListenAndServe(":8080", nil)
}
