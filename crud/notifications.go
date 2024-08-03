package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Notifications struct {

NotificationID int

UserID int

NotificationText string

NotificationType string

IsRead int

CreatedAt string

}
var tmplnotifications = template.Must(template.ParseGlob("notifications/*")) 

func Indexnotifications(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM notifications ORDER BY NotificationID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Notifications{}
    res := []Notifications{}

    for selDB.Next() {
        
        var NotificationID int
        
        var UserID int
        
        var NotificationText string
        
        var NotificationType string
        
        var IsRead int
        
        var CreatedAt string
        

        err = selDB.Scan(
        &NotificationID ,&UserID ,&NotificationText ,&NotificationType ,&IsRead ,&CreatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.NotificationID=NotificationID
       
       emp.UserID=UserID
       
       emp.NotificationText=NotificationText
       
       emp.NotificationType=NotificationType
       
       emp.IsRead=IsRead
       
       emp.CreatedAt=CreatedAt
       

        res = append(res, emp)
    }

    tmplnotifications.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Shownotifications(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM notifications WHERE NotificationID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Notifications{}

    for selDB.Next() {
       
       var NotificationID int
       var UserID int
       var NotificationText string
       var NotificationType string
       var IsRead int
       var CreatedAt string

        err = selDB.Scan(&NotificationID ,&UserID ,&NotificationText ,&NotificationType ,&IsRead ,&CreatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.NotificationID=NotificationID
       emp.UserID=UserID
       emp.NotificationText=NotificationText
       emp.NotificationType=NotificationType
       emp.IsRead=IsRead
       emp.CreatedAt=CreatedAt
    }

    tmplnotifications.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newnotifications(w http.ResponseWriter, r *http.Request) {
    tmplnotifications.ExecuteTemplate(w, "New", nil)
}

func Editnotifications(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM notifications WHERE NotificationID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Notifications{}

    for selDB.Next() {
      
      var NotificationID int
      var UserID int
      var NotificationText string
      var NotificationType string
      var IsRead int
      var CreatedAt string
        err = selDB.Scan(
        &NotificationID ,&UserID ,&NotificationText ,&NotificationType ,&IsRead ,&CreatedAt )
        if err != nil {
            panic(err.Error())
        }

      
      emp.NotificationID=NotificationID
      emp.UserID=UserID
      emp.NotificationText=NotificationText
      emp.NotificationType=NotificationType
      emp.IsRead=IsRead
      emp.CreatedAt=CreatedAt
    }

    tmplnotifications.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertnotifications(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        UserID :=r.FormValue("UserID")
        NotificationText :=r.FormValue("NotificationText")
        NotificationType :=r.FormValue("NotificationType")
        IsRead :=r.FormValue("IsRead")
        CreatedAt :=r.FormValue("CreatedAt")
        insForm, err := db.Prepare("INSERT INTO notifications(UserID ,NotificationText ,NotificationType ,IsRead ,CreatedAt ) VALUE (? ,? ,? ,? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        UserID ,NotificationText ,NotificationType ,IsRead ,CreatedAt )  
         
         log.Println("INSERT:UserID:" +UserID +"NotificationText:" +NotificationText +"NotificationType:" +NotificationType +"IsRead:" +IsRead +"CreatedAt:" +CreatedAt )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/notifications", 301)
}

func Updatenotifications(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       UserID:=r.FormValue("UserID")
       NotificationText:=r.FormValue("NotificationText")
       NotificationType:=r.FormValue("NotificationType")
       IsRead:=r.FormValue("IsRead")
       CreatedAt:=r.FormValue("CreatedAt")
       NotificationID:=r.FormValue("uNotificationID")
       
       insForm, err :=db.Prepare("UPDATE notifications SET UserID=? ,NotificationText=? ,NotificationType=? ,IsRead=? ,CreatedAt=?  WHERE NotificationID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        UserID ,NotificationText ,NotificationType ,IsRead ,CreatedAt ,NotificationID)
        log.Println("UPDATE:UserID: " + UserID + "NotificationText: " + NotificationText + "NotificationType: " + NotificationType + "IsRead: " + IsRead + "CreatedAt: " + CreatedAt ,NotificationID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/notifications", 301)
}

func Deletenotifications(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM notifications WHERE NotificationID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/notifications", 301)
}
