package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Messageentities struct {

MessageEntityID int

MessageID int

EntityID int

EntityValue string

}
var tmplmessageentities = template.Must(template.ParseGlob("messageentities/*")) 

func Indexmessageentities(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM messageentities ORDER BY MessageEntityID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Messageentities{}
    res := []Messageentities{}

    for selDB.Next() {
        
        var MessageEntityID int
        
        var MessageID int
        
        var EntityID int
        
        var EntityValue string
        

        err = selDB.Scan(
        &MessageEntityID ,&MessageID ,&EntityID ,&EntityValue )
        if err != nil {
            panic(err.Error())
        }

       
       emp.MessageEntityID=MessageEntityID
       
       emp.MessageID=MessageID
       
       emp.EntityID=EntityID
       
       emp.EntityValue=EntityValue
       

        res = append(res, emp)
    }

    tmplmessageentities.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showmessageentities(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM messageentities WHERE MessageEntityID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Messageentities{}

    for selDB.Next() {
       
       var MessageEntityID int
       var MessageID int
       var EntityID int
       var EntityValue string

        err = selDB.Scan(&MessageEntityID ,&MessageID ,&EntityID ,&EntityValue )
        if err != nil {
            panic(err.Error())
        }

       
       emp.MessageEntityID=MessageEntityID
       emp.MessageID=MessageID
       emp.EntityID=EntityID
       emp.EntityValue=EntityValue
    }

    tmplmessageentities.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newmessageentities(w http.ResponseWriter, r *http.Request) {
    tmplmessageentities.ExecuteTemplate(w, "New", nil)
}

func Editmessageentities(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM messageentities WHERE MessageEntityID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Messageentities{}

    for selDB.Next() {
      
      var MessageEntityID int
      var MessageID int
      var EntityID int
      var EntityValue string
        err = selDB.Scan(
        &MessageEntityID ,&MessageID ,&EntityID ,&EntityValue )
        if err != nil {
            panic(err.Error())
        }

      
      emp.MessageEntityID=MessageEntityID
      emp.MessageID=MessageID
      emp.EntityID=EntityID
      emp.EntityValue=EntityValue
    }

    tmplmessageentities.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertmessageentities(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        MessageID :=r.FormValue("MessageID")
        EntityID :=r.FormValue("EntityID")
        EntityValue :=r.FormValue("EntityValue")
        insForm, err := db.Prepare("INSERT INTO messageentities(MessageID ,EntityID ,EntityValue ) VALUE (? ,? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        MessageID ,EntityID ,EntityValue )  
         
         log.Println("INSERT:MessageID:" +MessageID +"EntityID:" +EntityID +"EntityValue:" +EntityValue )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/messageentities", 301)
}

func Updatemessageentities(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       MessageID:=r.FormValue("MessageID")
       EntityID:=r.FormValue("EntityID")
       EntityValue:=r.FormValue("EntityValue")
       MessageEntityID:=r.FormValue("uMessageEntityID")
       
       insForm, err :=db.Prepare("UPDATE messageentities SET MessageID=? ,EntityID=? ,EntityValue=?  WHERE MessageEntityID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        MessageID ,EntityID ,EntityValue ,MessageEntityID)
        log.Println("UPDATE:MessageID: " + MessageID + "EntityID: " + EntityID + "EntityValue: " + EntityValue ,MessageEntityID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/messageentities", 301)
}

func Deletemessageentities(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM messageentities WHERE MessageEntityID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/messageentities", 301)
}
