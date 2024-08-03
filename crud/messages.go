package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Messages struct {

MessageID int

ConversationID int

Sender string

MessageText string

MessageTimestamp string

}
var tmplmessages = template.Must(template.ParseGlob("messages/*")) 

func Indexmessages(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM messages ORDER BY MessageID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Messages{}
    res := []Messages{}

    for selDB.Next() {
        
        var MessageID int
        
        var ConversationID int
        
        var Sender string
        
        var MessageText string
        
        var MessageTimestamp string
        

        err = selDB.Scan(
        &MessageID ,&ConversationID ,&Sender ,&MessageText ,&MessageTimestamp )
        if err != nil {
            panic(err.Error())
        }

       
       emp.MessageID=MessageID
       
       emp.ConversationID=ConversationID
       
       emp.Sender=Sender
       
       emp.MessageText=MessageText
       
       emp.MessageTimestamp=MessageTimestamp
       

        res = append(res, emp)
    }

    tmplmessages.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showmessages(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM messages WHERE MessageID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Messages{}

    for selDB.Next() {
       
       var MessageID int
       var ConversationID int
       var Sender string
       var MessageText string
       var MessageTimestamp string

        err = selDB.Scan(&MessageID ,&ConversationID ,&Sender ,&MessageText ,&MessageTimestamp )
        if err != nil {
            panic(err.Error())
        }

       
       emp.MessageID=MessageID
       emp.ConversationID=ConversationID
       emp.Sender=Sender
       emp.MessageText=MessageText
       emp.MessageTimestamp=MessageTimestamp
    }

    tmplmessages.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newmessages(w http.ResponseWriter, r *http.Request) {
    tmplmessages.ExecuteTemplate(w, "New", nil)
}

func Editmessages(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM messages WHERE MessageID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Messages{}

    for selDB.Next() {
      
      var MessageID int
      var ConversationID int
      var Sender string
      var MessageText string
      var MessageTimestamp string
        err = selDB.Scan(
        &MessageID ,&ConversationID ,&Sender ,&MessageText ,&MessageTimestamp )
        if err != nil {
            panic(err.Error())
        }

      
      emp.MessageID=MessageID
      emp.ConversationID=ConversationID
      emp.Sender=Sender
      emp.MessageText=MessageText
      emp.MessageTimestamp=MessageTimestamp
    }

    tmplmessages.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertmessages(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        ConversationID :=r.FormValue("ConversationID")
        Sender :=r.FormValue("Sender")
        MessageText :=r.FormValue("MessageText")
        MessageTimestamp :=r.FormValue("MessageTimestamp")
        insForm, err := db.Prepare("INSERT INTO messages(ConversationID ,Sender ,MessageText ,MessageTimestamp ) VALUE (? ,? ,? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        ConversationID ,Sender ,MessageText ,MessageTimestamp )  
         
         log.Println("INSERT:ConversationID:" +ConversationID +"Sender:" +Sender +"MessageText:" +MessageText +"MessageTimestamp:" +MessageTimestamp )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/messages", 301)
}

func Updatemessages(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       ConversationID:=r.FormValue("ConversationID")
       Sender:=r.FormValue("Sender")
       MessageText:=r.FormValue("MessageText")
       MessageTimestamp:=r.FormValue("MessageTimestamp")
       MessageID:=r.FormValue("uMessageID")
       
       insForm, err :=db.Prepare("UPDATE messages SET ConversationID=? ,Sender=? ,MessageText=? ,MessageTimestamp=?  WHERE MessageID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        ConversationID ,Sender ,MessageText ,MessageTimestamp ,MessageID)
        log.Println("UPDATE:ConversationID: " + ConversationID + "Sender: " + Sender + "MessageText: " + MessageText + "MessageTimestamp: " + MessageTimestamp ,MessageID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/messages", 301)
}

func Deletemessages(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM messages WHERE MessageID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/messages", 301)
}
