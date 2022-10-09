package handlers

import (
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		io.WriteString(w, "Hello World\n")
	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

func PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		io.WriteString(w, "Posting Article...\n")
	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}
func ArticleListHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		io.WriteString(w, "Article List\n")
	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}
func ArticleShowHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		io.WriteString(w, "Article No.1\n")
	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)

	}
}

func PostNiceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		io.WriteString(w, "Posting Nice...\n")
	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

func PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		io.WriteString(w, "Posting Comment...\n")
	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}
