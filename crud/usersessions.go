package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Usersessions struct {

SessionID int

UserID int

SessionToken string

SessionStart string

SessionEnd string

}
var tmplusersessions = template.Must(template.ParseGlob("usersessions/*")) 

func Indexusersessions(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM usersessions ORDER BY SessionID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Usersessions{}
    res := []Usersessions{}

    for selDB.Next() {
        
        var SessionID int
        
        var UserID int
        
        var SessionToken string
        
        var SessionStart string
        
        var SessionEnd string
        

        err = selDB.Scan(
        &SessionID ,&UserID ,&SessionToken ,&SessionStart ,&SessionEnd )
        if err != nil {
            panic(err.Error())
        }

       
       emp.SessionID=SessionID
       
       emp.UserID=UserID
       
       emp.SessionToken=SessionToken
       
       emp.SessionStart=SessionStart
       
       emp.SessionEnd=SessionEnd
       

        res = append(res, emp)
    }

    tmplusersessions.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showusersessions(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM usersessions WHERE SessionID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Usersessions{}

    for selDB.Next() {
       
       var SessionID int
       var UserID int
       var SessionToken string
       var SessionStart string
       var SessionEnd string

        err = selDB.Scan(&SessionID ,&UserID ,&SessionToken ,&SessionStart ,&SessionEnd )
        if err != nil {
            panic(err.Error())
        }

       
       emp.SessionID=SessionID
       emp.UserID=UserID
       emp.SessionToken=SessionToken
       emp.SessionStart=SessionStart
       emp.SessionEnd=SessionEnd
    }

    tmplusersessions.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newusersessions(w http.ResponseWriter, r *http.Request) {
    tmplusersessions.ExecuteTemplate(w, "New", nil)
}

func Editusersessions(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM usersessions WHERE SessionID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Usersessions{}

    for selDB.Next() {
      
      var SessionID int
      var UserID int
      var SessionToken string
      var SessionStart string
      var SessionEnd string
        err = selDB.Scan(
        &SessionID ,&UserID ,&SessionToken ,&SessionStart ,&SessionEnd )
        if err != nil {
            panic(err.Error())
        }

      
      emp.SessionID=SessionID
      emp.UserID=UserID
      emp.SessionToken=SessionToken
      emp.SessionStart=SessionStart
      emp.SessionEnd=SessionEnd
    }

    tmplusersessions.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertusersessions(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        UserID :=r.FormValue("UserID")
        SessionToken :=r.FormValue("SessionToken")
        SessionStart :=r.FormValue("SessionStart")
        SessionEnd :=r.FormValue("SessionEnd")
        insForm, err := db.Prepare("INSERT INTO usersessions(UserID ,SessionToken ,SessionStart ,SessionEnd ) VALUE (? ,? ,? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        UserID ,SessionToken ,SessionStart ,SessionEnd )  
         
         log.Println("INSERT:UserID:" +UserID +"SessionToken:" +SessionToken +"SessionStart:" +SessionStart +"SessionEnd:" +SessionEnd )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/usersessions", 301)
}

func Updateusersessions(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       UserID:=r.FormValue("UserID")
       SessionToken:=r.FormValue("SessionToken")
       SessionStart:=r.FormValue("SessionStart")
       SessionEnd:=r.FormValue("SessionEnd")
       SessionID:=r.FormValue("uSessionID")
       
       insForm, err :=db.Prepare("UPDATE usersessions SET UserID=? ,SessionToken=? ,SessionStart=? ,SessionEnd=?  WHERE SessionID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        UserID ,SessionToken ,SessionStart ,SessionEnd ,SessionID)
        log.Println("UPDATE:UserID: " + UserID + "SessionToken: " + SessionToken + "SessionStart: " + SessionStart + "SessionEnd: " + SessionEnd ,SessionID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/usersessions", 301)
}

func Deleteusersessions(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM usersessions WHERE SessionID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/usersessions", 301)
}
