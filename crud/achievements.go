package crud

import (
	"log"
	"net/http"
	"text/template"
)

type Achievements struct {
	AchievementID int

	UserID int

	AchievementTitle string

	AchievementDetails string

	AchievementDate string

	CreatedAt string
}

var tmplachievements = template.Must(template.ParseGlob("achievements/*"))

func Indexachievements(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM achievements ORDER BY AchievementID DESC")

	if err != nil {
		panic(err.Error())
	}

	emp := Achievements{}
	res := []Achievements{}

	for selDB.Next() {

		var AchievementID int

		var UserID int

		var AchievementTitle string

		var AchievementDetails string

		var AchievementDate string

		var CreatedAt string

		err = selDB.Scan(
			&AchievementID, &UserID, &AchievementTitle, &AchievementDetails, &AchievementDate, &CreatedAt)
		if err != nil {
			panic(err.Error())
		}

		emp.AchievementID = AchievementID

		emp.UserID = UserID

		emp.AchievementTitle = AchievementTitle

		emp.AchievementDetails = AchievementDetails

		emp.AchievementDate = AchievementDate

		emp.CreatedAt = CreatedAt

		res = append(res, emp)
	}

	tmplachievements.ExecuteTemplate(w, "Index", res)
	defer db.Close()

	// defer db.Close()
}

func Showachievements(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")

	selDB, err := db.Query("SELECT * FROM achievements WHERE AchievementID=?", nId)

	if err != nil {
		panic(err.Error())
	}

	emp := Achievements{}

	for selDB.Next() {

		var AchievementID int
		var UserID int
		var AchievementTitle string
		var AchievementDetails string
		var AchievementDate string
		var CreatedAt string

		err = selDB.Scan(&AchievementID, &UserID, &AchievementTitle, &AchievementDetails, &AchievementDate, &CreatedAt)
		if err != nil {
			panic(err.Error())
		}

		emp.AchievementID = AchievementID
		emp.UserID = UserID
		emp.AchievementTitle = AchievementTitle
		emp.AchievementDetails = AchievementDetails
		emp.AchievementDate = AchievementDate
		emp.CreatedAt = CreatedAt
	}

	tmplachievements.ExecuteTemplate(w, "Show", emp)
	defer db.Close()
}

func Newachievements(w http.ResponseWriter, r *http.Request) {
	tmplachievements.ExecuteTemplate(w, "New", nil)
}

func Editachievements(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")

	selDB, err := db.Query("SELECT * FROM achievements WHERE AchievementID=?", nId)

	if err != nil {
		panic(err.Error())
	}

	emp := Achievements{}

	for selDB.Next() {

		var AchievementID int
		var UserID int
		var AchievementTitle string
		var AchievementDetails string
		var AchievementDate string
		var CreatedAt string
		err = selDB.Scan(
			&AchievementID, &UserID, &AchievementTitle, &AchievementDetails, &AchievementDate, &CreatedAt)
		if err != nil {
			panic(err.Error())
		}

		emp.AchievementID = AchievementID
		emp.UserID = UserID
		emp.AchievementTitle = AchievementTitle
		emp.AchievementDetails = AchievementDetails
		emp.AchievementDate = AchievementDate
		emp.CreatedAt = CreatedAt
	}

	tmplachievements.ExecuteTemplate(w, "Edit", emp)
	defer db.Close()
}

func Insertachievements(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	if r.Method == "POST" {

		UserID := r.FormValue("UserID")
		AchievementTitle := r.FormValue("AchievementTitle")
		AchievementDetails := r.FormValue("AchievementDetails")
		AchievementDate := r.FormValue("AchievementDate")
		CreatedAt := r.FormValue("CreatedAt")
		insForm, err := db.Prepare("INSERT INTO achievements(UserID ,AchievementTitle ,AchievementDetails ,AchievementDate ,CreatedAt ) VALUE (? ,? ,? ,? ,? )")

		if err != nil {
			panic(err.Error())
		}

		insForm.Exec(
			UserID, AchievementTitle, AchievementDetails, AchievementDate, CreatedAt)

		log.Println("INSERT:UserID:" + UserID + "AchievementTitle:" + AchievementTitle + "AchievementDetails:" + AchievementDetails + "AchievementDate:" + AchievementDate + "CreatedAt:" + CreatedAt)

	}

	defer db.Close()
	http.Redirect(w, r, "/achievements", 301)
}

func Updateachievements(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {

		UserID := r.FormValue("UserID")
		AchievementTitle := r.FormValue("AchievementTitle")
		AchievementDetails := r.FormValue("AchievementDetails")
		AchievementDate := r.FormValue("AchievementDate")
		CreatedAt := r.FormValue("CreatedAt")
		AchievementID := r.FormValue("uAchievementID")

		insForm, err := db.Prepare("UPDATE achievements SET UserID=? ,AchievementTitle=? ,AchievementDetails=? ,AchievementDate=? ,CreatedAt=?  WHERE AchievementID=?")

		if err != nil {
			panic(err.Error())
		}

		insForm.Exec(UserID, AchievementTitle, AchievementDetails, AchievementDate, CreatedAt, AchievementID)
		log.Println("UPDATE:UserID: "+UserID+"AchievementTitle: "+AchievementTitle+"AchievementDetails: "+AchievementDetails+"AchievementDate: "+AchievementDate+"CreatedAt: "+CreatedAt, AchievementID)

	}

	defer db.Close()
	http.Redirect(w, r, "/achievements", 301)
}

func Deleteachievements(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")

	delForm, err := db.Prepare("DELETE FROM achievements WHERE AchievementID=?")

	if err != nil {
		panic(err.Error())
	}

	delForm.Exec(emp)
	log.Println("DELETE")

	defer db.Close()
	http.Redirect(w, r, "/achievements", 301)
}
