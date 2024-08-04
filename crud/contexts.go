package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Contexts struct {

ContextID int

ContextName string

Description string

CreatedAt string

UpdatedAt string

}
var tmplcontexts = template.Must(template.ParseGlob("contexts/*")) 

func Indexcontexts(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM contexts ORDER BY ContextID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Contexts{}
    res := []Contexts{}

    for selDB.Next() {
        
        var ContextID int
        
        var ContextName string
        
        var Description string
        
        var CreatedAt string
        
        var UpdatedAt string
        

        err = selDB.Scan(
        &ContextID ,&ContextName ,&Description ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.ContextID=ContextID
       
       emp.ContextName=ContextName
       
       emp.Description=Description
       
       emp.CreatedAt=CreatedAt
       
       emp.UpdatedAt=UpdatedAt
       

        res = append(res, emp)
    }

    tmplcontexts.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showcontexts(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM contexts WHERE ContextID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Contexts{}

    for selDB.Next() {
       
       var ContextID int
       var ContextName string
       var Description string
       var CreatedAt string
       var UpdatedAt string

        err = selDB.Scan(&ContextID ,&ContextName ,&Description ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.ContextID=ContextID
       emp.ContextName=ContextName
       emp.Description=Description
       emp.CreatedAt=CreatedAt
       emp.UpdatedAt=UpdatedAt
    }

    tmplcontexts.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newcontexts(w http.ResponseWriter, r *http.Request) {
    tmplcontexts.ExecuteTemplate(w, "New", nil)
}

func Editcontexts(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM contexts WHERE ContextID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Contexts{}

    for selDB.Next() {
      
      var ContextID int
      var ContextName string
      var Description string
      var CreatedAt string
      var UpdatedAt string
        err = selDB.Scan(
        &ContextID ,&ContextName ,&Description ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

      
      emp.ContextID=ContextID
      emp.ContextName=ContextName
      emp.Description=Description
      emp.CreatedAt=CreatedAt
      emp.UpdatedAt=UpdatedAt
    }

    tmplcontexts.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertcontexts(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        ContextName :=r.FormValue("ContextName")
        Description :=r.FormValue("Description")
        CreatedAt :=r.FormValue("CreatedAt")
        UpdatedAt :=r.FormValue("UpdatedAt")
        insForm, err := db.Prepare("INSERT INTO contexts(ContextName ,Description ,CreatedAt ,UpdatedAt ) VALUE (? ,? ,? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        ContextName ,Description ,CreatedAt ,UpdatedAt )  
         
         log.Println("INSERT:ContextName:" +ContextName +"Description:" +Description +"CreatedAt:" +CreatedAt +"UpdatedAt:" +UpdatedAt )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/contexts", 301)
}

func Updatecontexts(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       ContextName:=r.FormValue("ContextName")
       Description:=r.FormValue("Description")
       CreatedAt:=r.FormValue("CreatedAt")
       UpdatedAt:=r.FormValue("UpdatedAt")
       ContextID:=r.FormValue("uContextID")
       
       insForm, err :=db.Prepare("UPDATE contexts SET ContextName=? ,Description=? ,CreatedAt=? ,UpdatedAt=?  WHERE ContextID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        ContextName ,Description ,CreatedAt ,UpdatedAt ,ContextID)
        log.Println("UPDATE:ContextName: " + ContextName + "Description: " + Description + "CreatedAt: " + CreatedAt + "UpdatedAt: " + UpdatedAt ,ContextID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/contexts", 301)
}

func Deletecontexts(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM contexts WHERE ContextID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/contexts", 301)
}
