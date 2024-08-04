package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Entities struct {

EntityID int

EntityName string

EntityType string

CreatedAt string

UpdatedAt string

}
var tmplentities = template.Must(template.ParseGlob("entities/*")) 

func Indexentities(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM entities ORDER BY EntityID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Entities{}
    res := []Entities{}

    for selDB.Next() {
        
        var EntityID int
        
        var EntityName string
        
        var EntityType string
        
        var CreatedAt string
        
        var UpdatedAt string
        

        err = selDB.Scan(
        &EntityID ,&EntityName ,&EntityType ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.EntityID=EntityID
       
       emp.EntityName=EntityName
       
       emp.EntityType=EntityType
       
       emp.CreatedAt=CreatedAt
       
       emp.UpdatedAt=UpdatedAt
       

        res = append(res, emp)
    }

    tmplentities.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showentities(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM entities WHERE EntityID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Entities{}

    for selDB.Next() {
       
       var EntityID int
       var EntityName string
       var EntityType string
       var CreatedAt string
       var UpdatedAt string

        err = selDB.Scan(&EntityID ,&EntityName ,&EntityType ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.EntityID=EntityID
       emp.EntityName=EntityName
       emp.EntityType=EntityType
       emp.CreatedAt=CreatedAt
       emp.UpdatedAt=UpdatedAt
    }

    tmplentities.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newentities(w http.ResponseWriter, r *http.Request) {
    tmplentities.ExecuteTemplate(w, "New", nil)
}

func Editentities(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM entities WHERE EntityID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Entities{}

    for selDB.Next() {
      
      var EntityID int
      var EntityName string
      var EntityType string
      var CreatedAt string
      var UpdatedAt string
        err = selDB.Scan(
        &EntityID ,&EntityName ,&EntityType ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

      
      emp.EntityID=EntityID
      emp.EntityName=EntityName
      emp.EntityType=EntityType
      emp.CreatedAt=CreatedAt
      emp.UpdatedAt=UpdatedAt
    }

    tmplentities.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertentities(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        EntityName :=r.FormValue("EntityName")
        EntityType :=r.FormValue("EntityType")
        CreatedAt :=r.FormValue("CreatedAt")
        UpdatedAt :=r.FormValue("UpdatedAt")
        insForm, err := db.Prepare("INSERT INTO entities(EntityName ,EntityType ,CreatedAt ,UpdatedAt ) VALUE (? ,? ,? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        EntityName ,EntityType ,CreatedAt ,UpdatedAt )  
         
         log.Println("INSERT:EntityName:" +EntityName +"EntityType:" +EntityType +"CreatedAt:" +CreatedAt +"UpdatedAt:" +UpdatedAt )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/entities", 301)
}

func Updateentities(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       EntityName:=r.FormValue("EntityName")
       EntityType:=r.FormValue("EntityType")
       CreatedAt:=r.FormValue("CreatedAt")
       UpdatedAt:=r.FormValue("UpdatedAt")
       EntityID:=r.FormValue("uEntityID")
       
       insForm, err :=db.Prepare("UPDATE entities SET EntityName=? ,EntityType=? ,CreatedAt=? ,UpdatedAt=?  WHERE EntityID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        EntityName ,EntityType ,CreatedAt ,UpdatedAt ,EntityID)
        log.Println("UPDATE:EntityName: " + EntityName + "EntityType: " + EntityType + "CreatedAt: " + CreatedAt + "UpdatedAt: " + UpdatedAt ,EntityID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/entities", 301)
}

func Deleteentities(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM entities WHERE EntityID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/entities", 301)
}
