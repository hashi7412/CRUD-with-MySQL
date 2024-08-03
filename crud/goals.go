package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Goals struct {

GoalID int

UserID int

GoalTitle string

GoalDescription string

GoalDeadline string

IsAchieved int

CreatedAt string

UpdatedAt string

}
var tmplgoals = template.Must(template.ParseGlob("goals/*")) 

func Indexgoals(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM goals ORDER BY GoalID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Goals{}
    res := []Goals{}

    for selDB.Next() {
        
        var GoalID int
        
        var UserID int
        
        var GoalTitle string
        
        var GoalDescription string
        
        var GoalDeadline string
        
        var IsAchieved int
        
        var CreatedAt string
        
        var UpdatedAt string
        

        err = selDB.Scan(
        &GoalID ,&UserID ,&GoalTitle ,&GoalDescription ,&GoalDeadline ,&IsAchieved ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.GoalID=GoalID
       
       emp.UserID=UserID
       
       emp.GoalTitle=GoalTitle
       
       emp.GoalDescription=GoalDescription
       
       emp.GoalDeadline=GoalDeadline
       
       emp.IsAchieved=IsAchieved
       
       emp.CreatedAt=CreatedAt
       
       emp.UpdatedAt=UpdatedAt
       

        res = append(res, emp)
    }

    tmplgoals.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showgoals(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM goals WHERE GoalID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Goals{}

    for selDB.Next() {
       
       var GoalID int
       var UserID int
       var GoalTitle string
       var GoalDescription string
       var GoalDeadline string
       var IsAchieved int
       var CreatedAt string
       var UpdatedAt string

        err = selDB.Scan(&GoalID ,&UserID ,&GoalTitle ,&GoalDescription ,&GoalDeadline ,&IsAchieved ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.GoalID=GoalID
       emp.UserID=UserID
       emp.GoalTitle=GoalTitle
       emp.GoalDescription=GoalDescription
       emp.GoalDeadline=GoalDeadline
       emp.IsAchieved=IsAchieved
       emp.CreatedAt=CreatedAt
       emp.UpdatedAt=UpdatedAt
    }

    tmplgoals.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newgoals(w http.ResponseWriter, r *http.Request) {
    tmplgoals.ExecuteTemplate(w, "New", nil)
}

func Editgoals(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM goals WHERE GoalID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Goals{}

    for selDB.Next() {
      
      var GoalID int
      var UserID int
      var GoalTitle string
      var GoalDescription string
      var GoalDeadline string
      var IsAchieved int
      var CreatedAt string
      var UpdatedAt string
        err = selDB.Scan(
        &GoalID ,&UserID ,&GoalTitle ,&GoalDescription ,&GoalDeadline ,&IsAchieved ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

      
      emp.GoalID=GoalID
      emp.UserID=UserID
      emp.GoalTitle=GoalTitle
      emp.GoalDescription=GoalDescription
      emp.GoalDeadline=GoalDeadline
      emp.IsAchieved=IsAchieved
      emp.CreatedAt=CreatedAt
      emp.UpdatedAt=UpdatedAt
    }

    tmplgoals.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertgoals(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        UserID :=r.FormValue("UserID")
        GoalTitle :=r.FormValue("GoalTitle")
        GoalDescription :=r.FormValue("GoalDescription")
        GoalDeadline :=r.FormValue("GoalDeadline")
        IsAchieved :=r.FormValue("IsAchieved")
        CreatedAt :=r.FormValue("CreatedAt")
        UpdatedAt :=r.FormValue("UpdatedAt")
        insForm, err := db.Prepare("INSERT INTO goals(UserID ,GoalTitle ,GoalDescription ,GoalDeadline ,IsAchieved ,CreatedAt ,UpdatedAt ) VALUE (? ,? ,? ,? ,? ,? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        UserID ,GoalTitle ,GoalDescription ,GoalDeadline ,IsAchieved ,CreatedAt ,UpdatedAt )  
         
         log.Println("INSERT:UserID:" +UserID +"GoalTitle:" +GoalTitle +"GoalDescription:" +GoalDescription +"GoalDeadline:" +GoalDeadline +"IsAchieved:" +IsAchieved +"CreatedAt:" +CreatedAt +"UpdatedAt:" +UpdatedAt )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/goals", 301)
}

func Updategoals(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       UserID:=r.FormValue("UserID")
       GoalTitle:=r.FormValue("GoalTitle")
       GoalDescription:=r.FormValue("GoalDescription")
       GoalDeadline:=r.FormValue("GoalDeadline")
       IsAchieved:=r.FormValue("IsAchieved")
       CreatedAt:=r.FormValue("CreatedAt")
       UpdatedAt:=r.FormValue("UpdatedAt")
       GoalID:=r.FormValue("uGoalID")
       
       insForm, err :=db.Prepare("UPDATE goals SET UserID=? ,GoalTitle=? ,GoalDescription=? ,GoalDeadline=? ,IsAchieved=? ,CreatedAt=? ,UpdatedAt=?  WHERE GoalID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        UserID ,GoalTitle ,GoalDescription ,GoalDeadline ,IsAchieved ,CreatedAt ,UpdatedAt ,GoalID)
        log.Println("UPDATE:UserID: " + UserID + "GoalTitle: " + GoalTitle + "GoalDescription: " + GoalDescription + "GoalDeadline: " + GoalDeadline + "IsAchieved: " + IsAchieved + "CreatedAt: " + CreatedAt + "UpdatedAt: " + UpdatedAt ,GoalID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/goals", 301)
}

func Deletegoals(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM goals WHERE GoalID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/goals", 301)
}
