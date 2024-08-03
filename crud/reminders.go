package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Reminders struct {

ReminderID int

UserID int

ReminderText string

ReminderTime string

IsCompleted int

CreatedAt string

UpdatedAt string

}
var tmplreminders = template.Must(template.ParseGlob("reminders/*")) 

func Indexreminders(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM reminders ORDER BY ReminderID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Reminders{}
    res := []Reminders{}

    for selDB.Next() {
        
        var ReminderID int
        
        var UserID int
        
        var ReminderText string
        
        var ReminderTime string
        
        var IsCompleted int
        
        var CreatedAt string
        
        var UpdatedAt string
        

        err = selDB.Scan(
        &ReminderID ,&UserID ,&ReminderText ,&ReminderTime ,&IsCompleted ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.ReminderID=ReminderID
       
       emp.UserID=UserID
       
       emp.ReminderText=ReminderText
       
       emp.ReminderTime=ReminderTime
       
       emp.IsCompleted=IsCompleted
       
       emp.CreatedAt=CreatedAt
       
       emp.UpdatedAt=UpdatedAt
       

        res = append(res, emp)
    }

    tmplreminders.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showreminders(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM reminders WHERE ReminderID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Reminders{}

    for selDB.Next() {
       
       var ReminderID int
       var UserID int
       var ReminderText string
       var ReminderTime string
       var IsCompleted int
       var CreatedAt string
       var UpdatedAt string

        err = selDB.Scan(&ReminderID ,&UserID ,&ReminderText ,&ReminderTime ,&IsCompleted ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.ReminderID=ReminderID
       emp.UserID=UserID
       emp.ReminderText=ReminderText
       emp.ReminderTime=ReminderTime
       emp.IsCompleted=IsCompleted
       emp.CreatedAt=CreatedAt
       emp.UpdatedAt=UpdatedAt
    }

    tmplreminders.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newreminders(w http.ResponseWriter, r *http.Request) {
    tmplreminders.ExecuteTemplate(w, "New", nil)
}

func Editreminders(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM reminders WHERE ReminderID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Reminders{}

    for selDB.Next() {
      
      var ReminderID int
      var UserID int
      var ReminderText string
      var ReminderTime string
      var IsCompleted int
      var CreatedAt string
      var UpdatedAt string
        err = selDB.Scan(
        &ReminderID ,&UserID ,&ReminderText ,&ReminderTime ,&IsCompleted ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

      
      emp.ReminderID=ReminderID
      emp.UserID=UserID
      emp.ReminderText=ReminderText
      emp.ReminderTime=ReminderTime
      emp.IsCompleted=IsCompleted
      emp.CreatedAt=CreatedAt
      emp.UpdatedAt=UpdatedAt
    }

    tmplreminders.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertreminders(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        UserID :=r.FormValue("UserID")
        ReminderText :=r.FormValue("ReminderText")
        ReminderTime :=r.FormValue("ReminderTime")
        IsCompleted :=r.FormValue("IsCompleted")
        CreatedAt :=r.FormValue("CreatedAt")
        UpdatedAt :=r.FormValue("UpdatedAt")
        insForm, err := db.Prepare("INSERT INTO reminders(UserID ,ReminderText ,ReminderTime ,IsCompleted ,CreatedAt ,UpdatedAt ) VALUE (? ,? ,? ,? ,? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        UserID ,ReminderText ,ReminderTime ,IsCompleted ,CreatedAt ,UpdatedAt )  
         
         log.Println("INSERT:UserID:" +UserID +"ReminderText:" +ReminderText +"ReminderTime:" +ReminderTime +"IsCompleted:" +IsCompleted +"CreatedAt:" +CreatedAt +"UpdatedAt:" +UpdatedAt )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/reminders", 301)
}

func Updatereminders(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       UserID:=r.FormValue("UserID")
       ReminderText:=r.FormValue("ReminderText")
       ReminderTime:=r.FormValue("ReminderTime")
       IsCompleted:=r.FormValue("IsCompleted")
       CreatedAt:=r.FormValue("CreatedAt")
       UpdatedAt:=r.FormValue("UpdatedAt")
       ReminderID:=r.FormValue("uReminderID")
       
       insForm, err :=db.Prepare("UPDATE reminders SET UserID=? ,ReminderText=? ,ReminderTime=? ,IsCompleted=? ,CreatedAt=? ,UpdatedAt=?  WHERE ReminderID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        UserID ,ReminderText ,ReminderTime ,IsCompleted ,CreatedAt ,UpdatedAt ,ReminderID)
        log.Println("UPDATE:UserID: " + UserID + "ReminderText: " + ReminderText + "ReminderTime: " + ReminderTime + "IsCompleted: " + IsCompleted + "CreatedAt: " + CreatedAt + "UpdatedAt: " + UpdatedAt ,ReminderID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/reminders", 301)
}

func Deletereminders(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM reminders WHERE ReminderID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/reminders", 301)
}
