package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Intents struct {

IntentID int

IntentName string

Description string

CreatedAt string

UpdatedAt string

}
var tmplintents = template.Must(template.ParseGlob("intents/*")) 

func Indexintents(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM intents ORDER BY IntentID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Intents{}
    res := []Intents{}

    for selDB.Next() {
        
        var IntentID int
        
        var IntentName string
        
        var Description string
        
        var CreatedAt string
        
        var UpdatedAt string
        

        err = selDB.Scan(
        &IntentID ,&IntentName ,&Description ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.IntentID=IntentID
       
       emp.IntentName=IntentName
       
       emp.Description=Description
       
       emp.CreatedAt=CreatedAt
       
       emp.UpdatedAt=UpdatedAt
       

        res = append(res, emp)
    }

    tmplintents.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showintents(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM intents WHERE IntentID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Intents{}

    for selDB.Next() {
       
       var IntentID int
       var IntentName string
       var Description string
       var CreatedAt string
       var UpdatedAt string

        err = selDB.Scan(&IntentID ,&IntentName ,&Description ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.IntentID=IntentID
       emp.IntentName=IntentName
       emp.Description=Description
       emp.CreatedAt=CreatedAt
       emp.UpdatedAt=UpdatedAt
    }

    tmplintents.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newintents(w http.ResponseWriter, r *http.Request) {
    tmplintents.ExecuteTemplate(w, "New", nil)
}

func Editintents(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM intents WHERE IntentID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Intents{}

    for selDB.Next() {
      
      var IntentID int
      var IntentName string
      var Description string
      var CreatedAt string
      var UpdatedAt string
        err = selDB.Scan(
        &IntentID ,&IntentName ,&Description ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

      
      emp.IntentID=IntentID
      emp.IntentName=IntentName
      emp.Description=Description
      emp.CreatedAt=CreatedAt
      emp.UpdatedAt=UpdatedAt
    }

    tmplintents.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertintents(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        IntentName :=r.FormValue("IntentName")
        Description :=r.FormValue("Description")
        CreatedAt :=r.FormValue("CreatedAt")
        UpdatedAt :=r.FormValue("UpdatedAt")
        insForm, err := db.Prepare("INSERT INTO intents(IntentName ,Description ,CreatedAt ,UpdatedAt ) VALUE (? ,? ,? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        IntentName ,Description ,CreatedAt ,UpdatedAt )  
         
         log.Println("INSERT:IntentName:" +IntentName +"Description:" +Description +"CreatedAt:" +CreatedAt +"UpdatedAt:" +UpdatedAt )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/intents", 301)
}

func Updateintents(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       IntentName:=r.FormValue("IntentName")
       Description:=r.FormValue("Description")
       CreatedAt:=r.FormValue("CreatedAt")
       UpdatedAt:=r.FormValue("UpdatedAt")
       IntentID:=r.FormValue("uIntentID")
       
       insForm, err :=db.Prepare("UPDATE intents SET IntentName=? ,Description=? ,CreatedAt=? ,UpdatedAt=?  WHERE IntentID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        IntentName ,Description ,CreatedAt ,UpdatedAt ,IntentID)
        log.Println("UPDATE:IntentName: " + IntentName + "Description: " + Description + "CreatedAt: " + CreatedAt + "UpdatedAt: " + UpdatedAt ,IntentID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/intents", 301)
}

func Deleteintents(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM intents WHERE IntentID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/intents", 301)
}
