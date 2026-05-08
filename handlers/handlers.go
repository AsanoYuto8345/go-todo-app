package handlers

import (
	"io"
	"net/http"
)

func TodoListHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Get All Todos!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetTodoByIdHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "do hogehoge")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func CreateTodoHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "create New Todo!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func PutTodoByIdHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPut {
		io.WriteString(w, "update Todo!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func DeleteTodoByIdHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodDelete {
		io.WriteString(w, "delete 3Todo!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
