package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Actions struct {

ActionID int

ActionName string

ActionType string

Description string

CreatedAt string

UpdatedAt string

}
var tmplactions = template.Must(template.ParseGlob("actions/*")) 

func Indexactions(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM actions ORDER BY ActionID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Actions{}
    res := []Actions{}

    for selDB.Next() {
        
        var ActionID int
        
        var ActionName string
        
        var ActionType string
        
        var Description string
        
        var CreatedAt string
        
        var UpdatedAt string
        

        err = selDB.Scan(
        &ActionID ,&ActionName ,&ActionType ,&Description ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.ActionID=ActionID
       
       emp.ActionName=ActionName
       
       emp.ActionType=ActionType
       
       emp.Description=Description
       
       emp.CreatedAt=CreatedAt
       
       emp.UpdatedAt=UpdatedAt
       

        res = append(res, emp)
    }

    tmplactions.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showactions(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM actions WHERE ActionID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Actions{}

    for selDB.Next() {
       
       var ActionID int
       var ActionName string
       var ActionType string
       var Description string
       var CreatedAt string
       var UpdatedAt string

        err = selDB.Scan(&ActionID ,&ActionName ,&ActionType ,&Description ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.ActionID=ActionID
       emp.ActionName=ActionName
       emp.ActionType=ActionType
       emp.Description=Description
       emp.CreatedAt=CreatedAt
       emp.UpdatedAt=UpdatedAt
    }

    tmplactions.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newactions(w http.ResponseWriter, r *http.Request) {
    tmplactions.ExecuteTemplate(w, "New", nil)
}

func Editactions(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM actions WHERE ActionID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Actions{}

    for selDB.Next() {
      
      var ActionID int
      var ActionName string
      var ActionType string
      var Description string
      var CreatedAt string
      var UpdatedAt string
        err = selDB.Scan(
        &ActionID ,&ActionName ,&ActionType ,&Description ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

      
      emp.ActionID=ActionID
      emp.ActionName=ActionName
      emp.ActionType=ActionType
      emp.Description=Description
      emp.CreatedAt=CreatedAt
      emp.UpdatedAt=UpdatedAt
    }

    tmplactions.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertactions(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        ActionName :=r.FormValue("ActionName")
        ActionType :=r.FormValue("ActionType")
        Description :=r.FormValue("Description")
        CreatedAt :=r.FormValue("CreatedAt")
        UpdatedAt :=r.FormValue("UpdatedAt")
        insForm, err := db.Prepare("INSERT INTO actions(ActionName ,ActionType ,Description ,CreatedAt ,UpdatedAt ) VALUE (? ,? ,? ,? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        ActionName ,ActionType ,Description ,CreatedAt ,UpdatedAt )  
         
         log.Println("INSERT:ActionName:" +ActionName +"ActionType:" +ActionType +"Description:" +Description +"CreatedAt:" +CreatedAt +"UpdatedAt:" +UpdatedAt )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/actions", 301)
}

func Updateactions(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       ActionName:=r.FormValue("ActionName")
       ActionType:=r.FormValue("ActionType")
       Description:=r.FormValue("Description")
       CreatedAt:=r.FormValue("CreatedAt")
       UpdatedAt:=r.FormValue("UpdatedAt")
       ActionID:=r.FormValue("uActionID")
       
       insForm, err :=db.Prepare("UPDATE actions SET ActionName=? ,ActionType=? ,Description=? ,CreatedAt=? ,UpdatedAt=?  WHERE ActionID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        ActionName ,ActionType ,Description ,CreatedAt ,UpdatedAt ,ActionID)
        log.Println("UPDATE:ActionName: " + ActionName + "ActionType: " + ActionType + "Description: " + Description + "CreatedAt: " + CreatedAt + "UpdatedAt: " + UpdatedAt ,ActionID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/actions", 301)
}

func Deleteactions(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM actions WHERE ActionID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/actions", 301)
}
