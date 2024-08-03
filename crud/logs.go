package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Logs struct {

LogID int

UserID int

ActionID int

LogDetails string

LogTimestamp string

}
var tmpllogs = template.Must(template.ParseGlob("logs/*")) 

func Indexlogs(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM logs ORDER BY LogID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Logs{}
    res := []Logs{}

    for selDB.Next() {
        
        var LogID int
        
        var UserID int
        
        var ActionID int
        
        var LogDetails string
        
        var LogTimestamp string
        

        err = selDB.Scan(
        &LogID ,&UserID ,&ActionID ,&LogDetails ,&LogTimestamp )
        if err != nil {
            panic(err.Error())
        }

       
       emp.LogID=LogID
       
       emp.UserID=UserID
       
       emp.ActionID=ActionID
       
       emp.LogDetails=LogDetails
       
       emp.LogTimestamp=LogTimestamp
       

        res = append(res, emp)
    }

    tmpllogs.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showlogs(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM logs WHERE LogID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Logs{}

    for selDB.Next() {
       
       var LogID int
       var UserID int
       var ActionID int
       var LogDetails string
       var LogTimestamp string

        err = selDB.Scan(&LogID ,&UserID ,&ActionID ,&LogDetails ,&LogTimestamp )
        if err != nil {
            panic(err.Error())
        }

       
       emp.LogID=LogID
       emp.UserID=UserID
       emp.ActionID=ActionID
       emp.LogDetails=LogDetails
       emp.LogTimestamp=LogTimestamp
    }

    tmpllogs.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newlogs(w http.ResponseWriter, r *http.Request) {
    tmpllogs.ExecuteTemplate(w, "New", nil)
}

func Editlogs(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM logs WHERE LogID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Logs{}

    for selDB.Next() {
      
      var LogID int
      var UserID int
      var ActionID int
      var LogDetails string
      var LogTimestamp string
        err = selDB.Scan(
        &LogID ,&UserID ,&ActionID ,&LogDetails ,&LogTimestamp )
        if err != nil {
            panic(err.Error())
        }

      
      emp.LogID=LogID
      emp.UserID=UserID
      emp.ActionID=ActionID
      emp.LogDetails=LogDetails
      emp.LogTimestamp=LogTimestamp
    }

    tmpllogs.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertlogs(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        UserID :=r.FormValue("UserID")
        ActionID :=r.FormValue("ActionID")
        LogDetails :=r.FormValue("LogDetails")
        LogTimestamp :=r.FormValue("LogTimestamp")
        insForm, err := db.Prepare("INSERT INTO logs(UserID ,ActionID ,LogDetails ,LogTimestamp ) VALUE (? ,? ,? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        UserID ,ActionID ,LogDetails ,LogTimestamp )  
         
         log.Println("INSERT:UserID:" +UserID +"ActionID:" +ActionID +"LogDetails:" +LogDetails +"LogTimestamp:" +LogTimestamp )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/logs", 301)
}

func Updatelogs(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       UserID:=r.FormValue("UserID")
       ActionID:=r.FormValue("ActionID")
       LogDetails:=r.FormValue("LogDetails")
       LogTimestamp:=r.FormValue("LogTimestamp")
       LogID:=r.FormValue("uLogID")
       
       insForm, err :=db.Prepare("UPDATE logs SET UserID=? ,ActionID=? ,LogDetails=? ,LogTimestamp=?  WHERE LogID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        UserID ,ActionID ,LogDetails ,LogTimestamp ,LogID)
        log.Println("UPDATE:UserID: " + UserID + "ActionID: " + ActionID + "LogDetails: " + LogDetails + "LogTimestamp: " + LogTimestamp ,LogID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/logs", 301)
}

func Deletelogs(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM logs WHERE LogID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/logs", 301)
}
