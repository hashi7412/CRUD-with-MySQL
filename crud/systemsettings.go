package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Systemsettings struct {

SettingID int

SettingKey string

SettingValue string

CreatedAt string

UpdatedAt string

}
var tmplsystemsettings = template.Must(template.ParseGlob("systemsettings/*")) 

func Indexsystemsettings(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM systemsettings ORDER BY SettingID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Systemsettings{}
    res := []Systemsettings{}

    for selDB.Next() {
        
        var SettingID int
        
        var SettingKey string
        
        var SettingValue string
        
        var CreatedAt string
        
        var UpdatedAt string
        

        err = selDB.Scan(
        &SettingID ,&SettingKey ,&SettingValue ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.SettingID=SettingID
       
       emp.SettingKey=SettingKey
       
       emp.SettingValue=SettingValue
       
       emp.CreatedAt=CreatedAt
       
       emp.UpdatedAt=UpdatedAt
       

        res = append(res, emp)
    }

    tmplsystemsettings.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showsystemsettings(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM systemsettings WHERE SettingID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Systemsettings{}

    for selDB.Next() {
       
       var SettingID int
       var SettingKey string
       var SettingValue string
       var CreatedAt string
       var UpdatedAt string

        err = selDB.Scan(&SettingID ,&SettingKey ,&SettingValue ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.SettingID=SettingID
       emp.SettingKey=SettingKey
       emp.SettingValue=SettingValue
       emp.CreatedAt=CreatedAt
       emp.UpdatedAt=UpdatedAt
    }

    tmplsystemsettings.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newsystemsettings(w http.ResponseWriter, r *http.Request) {
    tmplsystemsettings.ExecuteTemplate(w, "New", nil)
}

func Editsystemsettings(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM systemsettings WHERE SettingID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Systemsettings{}

    for selDB.Next() {
      
      var SettingID int
      var SettingKey string
      var SettingValue string
      var CreatedAt string
      var UpdatedAt string
        err = selDB.Scan(
        &SettingID ,&SettingKey ,&SettingValue ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

      
      emp.SettingID=SettingID
      emp.SettingKey=SettingKey
      emp.SettingValue=SettingValue
      emp.CreatedAt=CreatedAt
      emp.UpdatedAt=UpdatedAt
    }

    tmplsystemsettings.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertsystemsettings(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        SettingKey :=r.FormValue("SettingKey")
        SettingValue :=r.FormValue("SettingValue")
        CreatedAt :=r.FormValue("CreatedAt")
        UpdatedAt :=r.FormValue("UpdatedAt")
        insForm, err := db.Prepare("INSERT INTO systemsettings(SettingKey ,SettingValue ,CreatedAt ,UpdatedAt ) VALUE (? ,? ,? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        SettingKey ,SettingValue ,CreatedAt ,UpdatedAt )  
         
         log.Println("INSERT:SettingKey:" +SettingKey +"SettingValue:" +SettingValue +"CreatedAt:" +CreatedAt +"UpdatedAt:" +UpdatedAt )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/systemsettings", 301)
}

func Updatesystemsettings(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       SettingKey:=r.FormValue("SettingKey")
       SettingValue:=r.FormValue("SettingValue")
       CreatedAt:=r.FormValue("CreatedAt")
       UpdatedAt:=r.FormValue("UpdatedAt")
       SettingID:=r.FormValue("uSettingID")
       
       insForm, err :=db.Prepare("UPDATE systemsettings SET SettingKey=? ,SettingValue=? ,CreatedAt=? ,UpdatedAt=?  WHERE SettingID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        SettingKey ,SettingValue ,CreatedAt ,UpdatedAt ,SettingID)
        log.Println("UPDATE:SettingKey: " + SettingKey + "SettingValue: " + SettingValue + "CreatedAt: " + CreatedAt + "UpdatedAt: " + UpdatedAt ,SettingID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/systemsettings", 301)
}

func Deletesystemsettings(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM systemsettings WHERE SettingID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/systemsettings", 301)
}
