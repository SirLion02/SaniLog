package handler

import (
	"html/template"
	"net/http"
)

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/base.html", "templates/login.html")
	if err != nil {
		http.Error(w, "gagal load template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "base.html", nil); err != nil {
		http.Error(w, "gagal render: "+err.Error(), http.StatusInternalServerError)
	}
}

func LoginSubmitHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username_or_email")
	password := r.FormValue("password")

	ok, errMsg := ValidateLogin(username, password)
	if !ok {
		w.Write([]byte(`<p class="text-red-500 text-sm">` + errMsg + `</p>`))
		return
	}
	w.Write([]byte(`<p class="text-green-600 text-sm">Login berhasil (dummy)</p>`))
}

func ValidateLogin(username, password string) (bool, string) {
	if username == "" || password == "" {
		return false, "Username dan password wajib diisi"
	}
	if username == "admin" && password == "admin123" {
		return true, ""
	}
	return false, "Username atau password salah"
}
