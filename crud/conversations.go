package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Conversations struct {

ConversationID int

UserID int

StartTime string

EndTime string

}
var tmplconversations = template.Must(template.ParseGlob("conversations/*")) 

func Indexconversations(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM conversations ORDER BY ConversationID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Conversations{}
    res := []Conversations{}

    for selDB.Next() {
        
        var ConversationID int
        
        var UserID int
        
        var StartTime string
        
        var EndTime string
        

        err = selDB.Scan(
        &ConversationID ,&UserID ,&StartTime ,&EndTime )
        if err != nil {
            panic(err.Error())
        }

       
       emp.ConversationID=ConversationID
       
       emp.UserID=UserID
       
       emp.StartTime=StartTime
       
       emp.EndTime=EndTime
       

        res = append(res, emp)
    }

    tmplconversations.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showconversations(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM conversations WHERE ConversationID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Conversations{}

    for selDB.Next() {
       
       var ConversationID int
       var UserID int
       var StartTime string
       var EndTime string

        err = selDB.Scan(&ConversationID ,&UserID ,&StartTime ,&EndTime )
        if err != nil {
            panic(err.Error())
        }

       
       emp.ConversationID=ConversationID
       emp.UserID=UserID
       emp.StartTime=StartTime
       emp.EndTime=EndTime
    }

    tmplconversations.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newconversations(w http.ResponseWriter, r *http.Request) {
    tmplconversations.ExecuteTemplate(w, "New", nil)
}

func Editconversations(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM conversations WHERE ConversationID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Conversations{}

    for selDB.Next() {
      
      var ConversationID int
      var UserID int
      var StartTime string
      var EndTime string
        err = selDB.Scan(
        &ConversationID ,&UserID ,&StartTime ,&EndTime )
        if err != nil {
            panic(err.Error())
        }

      
      emp.ConversationID=ConversationID
      emp.UserID=UserID
      emp.StartTime=StartTime
      emp.EndTime=EndTime
    }

    tmplconversations.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertconversations(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        UserID :=r.FormValue("UserID")
        StartTime :=r.FormValue("StartTime")
        EndTime :=r.FormValue("EndTime")
        insForm, err := db.Prepare("INSERT INTO conversations(UserID ,StartTime ,EndTime ) VALUE (? ,? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        UserID ,StartTime ,EndTime )  
         
         log.Println("INSERT:UserID:" +UserID +"StartTime:" +StartTime +"EndTime:" +EndTime )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/conversations", 301)
}

func Updateconversations(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       UserID:=r.FormValue("UserID")
       StartTime:=r.FormValue("StartTime")
       EndTime:=r.FormValue("EndTime")
       ConversationID:=r.FormValue("uConversationID")
       
       insForm, err :=db.Prepare("UPDATE conversations SET UserID=? ,StartTime=? ,EndTime=?  WHERE ConversationID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        UserID ,StartTime ,EndTime ,ConversationID)
        log.Println("UPDATE:UserID: " + UserID + "StartTime: " + StartTime + "EndTime: " + EndTime ,ConversationID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/conversations", 301)
}

func Deleteconversations(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM conversations WHERE ConversationID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/conversations", 301)
}
