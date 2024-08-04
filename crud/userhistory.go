package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Userhistory struct {

HistoryID int

UserID int

HistoryData string

CreatedAt string

}
var tmpluserhistory = template.Must(template.ParseGlob("userhistory/*")) 

func Indexuserhistory(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM userhistory ORDER BY HistoryID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Userhistory{}
    res := []Userhistory{}

    for selDB.Next() {
        
        var HistoryID int
        
        var UserID int
        
        var HistoryData string
        
        var CreatedAt string
        

        err = selDB.Scan(
        &HistoryID ,&UserID ,&HistoryData ,&CreatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.HistoryID=HistoryID
       
       emp.UserID=UserID
       
       emp.HistoryData=HistoryData
       
       emp.CreatedAt=CreatedAt
       

        res = append(res, emp)
    }

    tmpluserhistory.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showuserhistory(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM userhistory WHERE HistoryID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Userhistory{}

    for selDB.Next() {
       
       var HistoryID int
       var UserID int
       var HistoryData string
       var CreatedAt string

        err = selDB.Scan(&HistoryID ,&UserID ,&HistoryData ,&CreatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.HistoryID=HistoryID
       emp.UserID=UserID
       emp.HistoryData=HistoryData
       emp.CreatedAt=CreatedAt
    }

    tmpluserhistory.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newuserhistory(w http.ResponseWriter, r *http.Request) {
    tmpluserhistory.ExecuteTemplate(w, "New", nil)
}

func Edituserhistory(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM userhistory WHERE HistoryID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Userhistory{}

    for selDB.Next() {
      
      var HistoryID int
      var UserID int
      var HistoryData string
      var CreatedAt string
        err = selDB.Scan(
        &HistoryID ,&UserID ,&HistoryData ,&CreatedAt )
        if err != nil {
            panic(err.Error())
        }

      
      emp.HistoryID=HistoryID
      emp.UserID=UserID
      emp.HistoryData=HistoryData
      emp.CreatedAt=CreatedAt
    }

    tmpluserhistory.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertuserhistory(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        UserID :=r.FormValue("UserID")
        HistoryData :=r.FormValue("HistoryData")
        CreatedAt :=r.FormValue("CreatedAt")
        insForm, err := db.Prepare("INSERT INTO userhistory(UserID ,HistoryData ,CreatedAt ) VALUE (? ,? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        UserID ,HistoryData ,CreatedAt )  
         
         log.Println("INSERT:UserID:" +UserID +"HistoryData:" +HistoryData +"CreatedAt:" +CreatedAt )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/userhistory", 301)
}

func Updateuserhistory(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       UserID:=r.FormValue("UserID")
       HistoryData:=r.FormValue("HistoryData")
       CreatedAt:=r.FormValue("CreatedAt")
       HistoryID:=r.FormValue("uHistoryID")
       
       insForm, err :=db.Prepare("UPDATE userhistory SET UserID=? ,HistoryData=? ,CreatedAt=?  WHERE HistoryID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        UserID ,HistoryData ,CreatedAt ,HistoryID)
        log.Println("UPDATE:UserID: " + UserID + "HistoryData: " + HistoryData + "CreatedAt: " + CreatedAt ,HistoryID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/userhistory", 301)
}

func Deleteuserhistory(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM userhistory WHERE HistoryID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/userhistory", 301)
}
