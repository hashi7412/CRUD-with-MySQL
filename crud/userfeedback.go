package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Userfeedback struct {

FeedbackID int

UserID int

FeedbackText string

FeedbackRating int

CreatedAt string

}
var tmpluserfeedback = template.Must(template.ParseGlob("userfeedback/*")) 

func Indexuserfeedback(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM userfeedback ORDER BY FeedbackID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Userfeedback{}
    res := []Userfeedback{}

    for selDB.Next() {
        
        var FeedbackID int
        
        var UserID int
        
        var FeedbackText string
        
        var FeedbackRating int
        
        var CreatedAt string
        

        err = selDB.Scan(
        &FeedbackID ,&UserID ,&FeedbackText ,&FeedbackRating ,&CreatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.FeedbackID=FeedbackID
       
       emp.UserID=UserID
       
       emp.FeedbackText=FeedbackText
       
       emp.FeedbackRating=FeedbackRating
       
       emp.CreatedAt=CreatedAt
       

        res = append(res, emp)
    }

    tmpluserfeedback.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showuserfeedback(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM userfeedback WHERE FeedbackID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Userfeedback{}

    for selDB.Next() {
       
       var FeedbackID int
       var UserID int
       var FeedbackText string
       var FeedbackRating int
       var CreatedAt string

        err = selDB.Scan(&FeedbackID ,&UserID ,&FeedbackText ,&FeedbackRating ,&CreatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.FeedbackID=FeedbackID
       emp.UserID=UserID
       emp.FeedbackText=FeedbackText
       emp.FeedbackRating=FeedbackRating
       emp.CreatedAt=CreatedAt
    }

    tmpluserfeedback.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newuserfeedback(w http.ResponseWriter, r *http.Request) {
    tmpluserfeedback.ExecuteTemplate(w, "New", nil)
}

func Edituserfeedback(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM userfeedback WHERE FeedbackID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Userfeedback{}

    for selDB.Next() {
      
      var FeedbackID int
      var UserID int
      var FeedbackText string
      var FeedbackRating int
      var CreatedAt string
        err = selDB.Scan(
        &FeedbackID ,&UserID ,&FeedbackText ,&FeedbackRating ,&CreatedAt )
        if err != nil {
            panic(err.Error())
        }

      
      emp.FeedbackID=FeedbackID
      emp.UserID=UserID
      emp.FeedbackText=FeedbackText
      emp.FeedbackRating=FeedbackRating
      emp.CreatedAt=CreatedAt
    }

    tmpluserfeedback.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertuserfeedback(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        UserID :=r.FormValue("UserID")
        FeedbackText :=r.FormValue("FeedbackText")
        FeedbackRating :=r.FormValue("FeedbackRating")
        CreatedAt :=r.FormValue("CreatedAt")
        insForm, err := db.Prepare("INSERT INTO userfeedback(UserID ,FeedbackText ,FeedbackRating ,CreatedAt ) VALUE (? ,? ,? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        UserID ,FeedbackText ,FeedbackRating ,CreatedAt )  
         
         log.Println("INSERT:UserID:" +UserID +"FeedbackText:" +FeedbackText +"FeedbackRating:" +FeedbackRating +"CreatedAt:" +CreatedAt )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/userfeedback", 301)
}

func Updateuserfeedback(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       UserID:=r.FormValue("UserID")
       FeedbackText:=r.FormValue("FeedbackText")
       FeedbackRating:=r.FormValue("FeedbackRating")
       CreatedAt:=r.FormValue("CreatedAt")
       FeedbackID:=r.FormValue("uFeedbackID")
       
       insForm, err :=db.Prepare("UPDATE userfeedback SET UserID=? ,FeedbackText=? ,FeedbackRating=? ,CreatedAt=?  WHERE FeedbackID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        UserID ,FeedbackText ,FeedbackRating ,CreatedAt ,FeedbackID)
        log.Println("UPDATE:UserID: " + UserID + "FeedbackText: " + FeedbackText + "FeedbackRating: " + FeedbackRating + "CreatedAt: " + CreatedAt ,FeedbackID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/userfeedback", 301)
}

func Deleteuserfeedback(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM userfeedback WHERE FeedbackID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/userfeedback", 301)
}
