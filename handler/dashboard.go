package handler

import (
	"html/template"
	"net/http"
)

type DashboardData struct {
	Nama      string
	Level     int
	TotalXP   int
	TotalSesi int

	TotalKg     float64
	TotalDurasi int
	HVP         float64
	LVP         float64
	Organik     float64
	Residue     float64
	B3          float64

	SesiTerbaru []Sesi

	JumlahAchievement int
	Achievements      []Achievement
}

type Sesi struct {
	Lokasi  string
	Tanggal string
	Kg      float64
	XP      int
}

type Achievement struct {
	Nama string
	Icon string
}

func DashboardPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/base.html", "templates/dashboard.html")
	if err != nil {
		http.Error(w, "gagal load template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := DashboardData{
		Nama:      "Layen",
		Level:     5,
		TotalXP:   1240,
		TotalSesi: 12,

		TotalKg:           45.8,
		TotalDurasi:       8,
		JumlahAchievement: 5,
		HVP:               12,
		LVP:               10,
		Organik:           20,
		Residue:           3.8,
		B3:                0,

		SesiTerbaru: []Sesi{
			{Lokasi: "Depok, Jawa Barat", Tanggal: "12 Okt 2023", Kg: 9.5, XP: 112},
			{Lokasi: "Jakarta Selatan", Tanggal: "05 Okt 2023", Kg: 4.2, XP: 50},
			{Lokasi: "Sungai Ciliwung, Depok", Tanggal: "28 Sep 2023", Kg: 7.8, XP: 88},
		},

		Achievements: []Achievement{
			{Nama: "1kg Club", Icon: "🏆"},
			{Nama: "Eco Hero", Icon: "⭐"},
		},
	}

	if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
		http.Error(w, "gagal render: "+err.Error(), http.StatusInternalServerError)
	}
}
