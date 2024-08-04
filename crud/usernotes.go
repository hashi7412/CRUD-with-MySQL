package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Usernotes struct {

NoteID int

UserID int

NoteTitle string

NoteContent string

CreatedAt string

UpdatedAt string

}
var tmplusernotes = template.Must(template.ParseGlob("usernotes/*")) 

func Indexusernotes(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM usernotes ORDER BY NoteID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Usernotes{}
    res := []Usernotes{}

    for selDB.Next() {
        
        var NoteID int
        
        var UserID int
        
        var NoteTitle string
        
        var NoteContent string
        
        var CreatedAt string
        
        var UpdatedAt string
        

        err = selDB.Scan(
        &NoteID ,&UserID ,&NoteTitle ,&NoteContent ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.NoteID=NoteID
       
       emp.UserID=UserID
       
       emp.NoteTitle=NoteTitle
       
       emp.NoteContent=NoteContent
       
       emp.CreatedAt=CreatedAt
       
       emp.UpdatedAt=UpdatedAt
       

        res = append(res, emp)
    }

    tmplusernotes.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showusernotes(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM usernotes WHERE NoteID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Usernotes{}

    for selDB.Next() {
       
       var NoteID int
       var UserID int
       var NoteTitle string
       var NoteContent string
       var CreatedAt string
       var UpdatedAt string

        err = selDB.Scan(&NoteID ,&UserID ,&NoteTitle ,&NoteContent ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.NoteID=NoteID
       emp.UserID=UserID
       emp.NoteTitle=NoteTitle
       emp.NoteContent=NoteContent
       emp.CreatedAt=CreatedAt
       emp.UpdatedAt=UpdatedAt
    }

    tmplusernotes.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newusernotes(w http.ResponseWriter, r *http.Request) {
    tmplusernotes.ExecuteTemplate(w, "New", nil)
}

func Editusernotes(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM usernotes WHERE NoteID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Usernotes{}

    for selDB.Next() {
      
      var NoteID int
      var UserID int
      var NoteTitle string
      var NoteContent string
      var CreatedAt string
      var UpdatedAt string
        err = selDB.Scan(
        &NoteID ,&UserID ,&NoteTitle ,&NoteContent ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

      
      emp.NoteID=NoteID
      emp.UserID=UserID
      emp.NoteTitle=NoteTitle
      emp.NoteContent=NoteContent
      emp.CreatedAt=CreatedAt
      emp.UpdatedAt=UpdatedAt
    }

    tmplusernotes.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertusernotes(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        UserID :=r.FormValue("UserID")
        NoteTitle :=r.FormValue("NoteTitle")
        NoteContent :=r.FormValue("NoteContent")
        CreatedAt :=r.FormValue("CreatedAt")
        UpdatedAt :=r.FormValue("UpdatedAt")
        insForm, err := db.Prepare("INSERT INTO usernotes(UserID ,NoteTitle ,NoteContent ,CreatedAt ,UpdatedAt ) VALUE (? ,? ,? ,? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        UserID ,NoteTitle ,NoteContent ,CreatedAt ,UpdatedAt )  
         
         log.Println("INSERT:UserID:" +UserID +"NoteTitle:" +NoteTitle +"NoteContent:" +NoteContent +"CreatedAt:" +CreatedAt +"UpdatedAt:" +UpdatedAt )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/usernotes", 301)
}

func Updateusernotes(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       UserID:=r.FormValue("UserID")
       NoteTitle:=r.FormValue("NoteTitle")
       NoteContent:=r.FormValue("NoteContent")
       CreatedAt:=r.FormValue("CreatedAt")
       UpdatedAt:=r.FormValue("UpdatedAt")
       NoteID:=r.FormValue("uNoteID")
       
       insForm, err :=db.Prepare("UPDATE usernotes SET UserID=? ,NoteTitle=? ,NoteContent=? ,CreatedAt=? ,UpdatedAt=?  WHERE NoteID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        UserID ,NoteTitle ,NoteContent ,CreatedAt ,UpdatedAt ,NoteID)
        log.Println("UPDATE:UserID: " + UserID + "NoteTitle: " + NoteTitle + "NoteContent: " + NoteContent + "CreatedAt: " + CreatedAt + "UpdatedAt: " + UpdatedAt ,NoteID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/usernotes", 301)
}

func Deleteusernotes(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM usernotes WHERE NoteID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/usernotes", 301)
}
