package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Knowledgebase struct {

KnowledgeID int

Topic string

Content string

CreatedAt string

UpdatedAt string

}
var tmplknowledgebase = template.Must(template.ParseGlob("knowledgebase/*")) 

func Indexknowledgebase(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM knowledgebase ORDER BY KnowledgeID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Knowledgebase{}
    res := []Knowledgebase{}

    for selDB.Next() {
        
        var KnowledgeID int
        
        var Topic string
        
        var Content string
        
        var CreatedAt string
        
        var UpdatedAt string
        

        err = selDB.Scan(
        &KnowledgeID ,&Topic ,&Content ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.KnowledgeID=KnowledgeID
       
       emp.Topic=Topic
       
       emp.Content=Content
       
       emp.CreatedAt=CreatedAt
       
       emp.UpdatedAt=UpdatedAt
       

        res = append(res, emp)
    }

    tmplknowledgebase.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showknowledgebase(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM knowledgebase WHERE KnowledgeID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Knowledgebase{}

    for selDB.Next() {
       
       var KnowledgeID int
       var Topic string
       var Content string
       var CreatedAt string
       var UpdatedAt string

        err = selDB.Scan(&KnowledgeID ,&Topic ,&Content ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.KnowledgeID=KnowledgeID
       emp.Topic=Topic
       emp.Content=Content
       emp.CreatedAt=CreatedAt
       emp.UpdatedAt=UpdatedAt
    }

    tmplknowledgebase.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newknowledgebase(w http.ResponseWriter, r *http.Request) {
    tmplknowledgebase.ExecuteTemplate(w, "New", nil)
}

func Editknowledgebase(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM knowledgebase WHERE KnowledgeID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Knowledgebase{}

    for selDB.Next() {
      
      var KnowledgeID int
      var Topic string
      var Content string
      var CreatedAt string
      var UpdatedAt string
        err = selDB.Scan(
        &KnowledgeID ,&Topic ,&Content ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

      
      emp.KnowledgeID=KnowledgeID
      emp.Topic=Topic
      emp.Content=Content
      emp.CreatedAt=CreatedAt
      emp.UpdatedAt=UpdatedAt
    }

    tmplknowledgebase.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertknowledgebase(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        Topic :=r.FormValue("Topic")
        Content :=r.FormValue("Content")
        CreatedAt :=r.FormValue("CreatedAt")
        UpdatedAt :=r.FormValue("UpdatedAt")
        insForm, err := db.Prepare("INSERT INTO knowledgebase(Topic ,Content ,CreatedAt ,UpdatedAt ) VALUE (? ,? ,? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        Topic ,Content ,CreatedAt ,UpdatedAt )  
         
         log.Println("INSERT:Topic:" +Topic +"Content:" +Content +"CreatedAt:" +CreatedAt +"UpdatedAt:" +UpdatedAt )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/knowledgebase", 301)
}

func Updateknowledgebase(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       Topic:=r.FormValue("Topic")
       Content:=r.FormValue("Content")
       CreatedAt:=r.FormValue("CreatedAt")
       UpdatedAt:=r.FormValue("UpdatedAt")
       KnowledgeID:=r.FormValue("uKnowledgeID")
       
       insForm, err :=db.Prepare("UPDATE knowledgebase SET Topic=? ,Content=? ,CreatedAt=? ,UpdatedAt=?  WHERE KnowledgeID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        Topic ,Content ,CreatedAt ,UpdatedAt ,KnowledgeID)
        log.Println("UPDATE:Topic: " + Topic + "Content: " + Content + "CreatedAt: " + CreatedAt + "UpdatedAt: " + UpdatedAt ,KnowledgeID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/knowledgebase", 301)
}

func Deleteknowledgebase(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM knowledgebase WHERE KnowledgeID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/knowledgebase", 301)
}
