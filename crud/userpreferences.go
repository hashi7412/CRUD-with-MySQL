package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Userpreferences struct {

PreferenceID int

UserID int

PreferenceKey string

PreferenceValue string

CreatedAt string

UpdatedAt string

}
var tmpluserpreferences = template.Must(template.ParseGlob("userpreferences/*")) 

func Indexuserpreferences(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM userpreferences ORDER BY PreferenceID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Userpreferences{}
    res := []Userpreferences{}

    for selDB.Next() {
        
        var PreferenceID int
        
        var UserID int
        
        var PreferenceKey string
        
        var PreferenceValue string
        
        var CreatedAt string
        
        var UpdatedAt string
        

        err = selDB.Scan(
        &PreferenceID ,&UserID ,&PreferenceKey ,&PreferenceValue ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.PreferenceID=PreferenceID
       
       emp.UserID=UserID
       
       emp.PreferenceKey=PreferenceKey
       
       emp.PreferenceValue=PreferenceValue
       
       emp.CreatedAt=CreatedAt
       
       emp.UpdatedAt=UpdatedAt
       

        res = append(res, emp)
    }

    tmpluserpreferences.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showuserpreferences(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM userpreferences WHERE PreferenceID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Userpreferences{}

    for selDB.Next() {
       
       var PreferenceID int
       var UserID int
       var PreferenceKey string
       var PreferenceValue string
       var CreatedAt string
       var UpdatedAt string

        err = selDB.Scan(&PreferenceID ,&UserID ,&PreferenceKey ,&PreferenceValue ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.PreferenceID=PreferenceID
       emp.UserID=UserID
       emp.PreferenceKey=PreferenceKey
       emp.PreferenceValue=PreferenceValue
       emp.CreatedAt=CreatedAt
       emp.UpdatedAt=UpdatedAt
    }

    tmpluserpreferences.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newuserpreferences(w http.ResponseWriter, r *http.Request) {
    tmpluserpreferences.ExecuteTemplate(w, "New", nil)
}

func Edituserpreferences(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM userpreferences WHERE PreferenceID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Userpreferences{}

    for selDB.Next() {
      
      var PreferenceID int
      var UserID int
      var PreferenceKey string
      var PreferenceValue string
      var CreatedAt string
      var UpdatedAt string
        err = selDB.Scan(
        &PreferenceID ,&UserID ,&PreferenceKey ,&PreferenceValue ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

      
      emp.PreferenceID=PreferenceID
      emp.UserID=UserID
      emp.PreferenceKey=PreferenceKey
      emp.PreferenceValue=PreferenceValue
      emp.CreatedAt=CreatedAt
      emp.UpdatedAt=UpdatedAt
    }

    tmpluserpreferences.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertuserpreferences(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        UserID :=r.FormValue("UserID")
        PreferenceKey :=r.FormValue("PreferenceKey")
        PreferenceValue :=r.FormValue("PreferenceValue")
        CreatedAt :=r.FormValue("CreatedAt")
        UpdatedAt :=r.FormValue("UpdatedAt")
        insForm, err := db.Prepare("INSERT INTO userpreferences(UserID ,PreferenceKey ,PreferenceValue ,CreatedAt ,UpdatedAt ) VALUE (? ,? ,? ,? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        UserID ,PreferenceKey ,PreferenceValue ,CreatedAt ,UpdatedAt )  
         
         log.Println("INSERT:UserID:" +UserID +"PreferenceKey:" +PreferenceKey +"PreferenceValue:" +PreferenceValue +"CreatedAt:" +CreatedAt +"UpdatedAt:" +UpdatedAt )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/userpreferences", 301)
}

func Updateuserpreferences(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       UserID:=r.FormValue("UserID")
       PreferenceKey:=r.FormValue("PreferenceKey")
       PreferenceValue:=r.FormValue("PreferenceValue")
       CreatedAt:=r.FormValue("CreatedAt")
       UpdatedAt:=r.FormValue("UpdatedAt")
       PreferenceID:=r.FormValue("uPreferenceID")
       
       insForm, err :=db.Prepare("UPDATE userpreferences SET UserID=? ,PreferenceKey=? ,PreferenceValue=? ,CreatedAt=? ,UpdatedAt=?  WHERE PreferenceID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        UserID ,PreferenceKey ,PreferenceValue ,CreatedAt ,UpdatedAt ,PreferenceID)
        log.Println("UPDATE:UserID: " + UserID + "PreferenceKey: " + PreferenceKey + "PreferenceValue: " + PreferenceValue + "CreatedAt: " + CreatedAt + "UpdatedAt: " + UpdatedAt ,PreferenceID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/userpreferences", 301)
}

func Deleteuserpreferences(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM userpreferences WHERE PreferenceID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/userpreferences", 301)
}
