package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Intentactions struct {

IntentActionID int

IntentID int

ActionID int

}
var tmplintentactions = template.Must(template.ParseGlob("intentactions/*")) 

func Indexintentactions(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM intentactions ORDER BY IntentActionID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Intentactions{}
    res := []Intentactions{}

    for selDB.Next() {
        
        var IntentActionID int
        
        var IntentID int
        
        var ActionID int
        

        err = selDB.Scan(
        &IntentActionID ,&IntentID ,&ActionID )
        if err != nil {
            panic(err.Error())
        }

       
       emp.IntentActionID=IntentActionID
       
       emp.IntentID=IntentID
       
       emp.ActionID=ActionID
       

        res = append(res, emp)
    }

    tmplintentactions.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showintentactions(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM intentactions WHERE IntentActionID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Intentactions{}

    for selDB.Next() {
       
       var IntentActionID int
       var IntentID int
       var ActionID int

        err = selDB.Scan(&IntentActionID ,&IntentID ,&ActionID )
        if err != nil {
            panic(err.Error())
        }

       
       emp.IntentActionID=IntentActionID
       emp.IntentID=IntentID
       emp.ActionID=ActionID
    }

    tmplintentactions.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newintentactions(w http.ResponseWriter, r *http.Request) {
    tmplintentactions.ExecuteTemplate(w, "New", nil)
}

func Editintentactions(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM intentactions WHERE IntentActionID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Intentactions{}

    for selDB.Next() {
      
      var IntentActionID int
      var IntentID int
      var ActionID int
        err = selDB.Scan(
        &IntentActionID ,&IntentID ,&ActionID )
        if err != nil {
            panic(err.Error())
        }

      
      emp.IntentActionID=IntentActionID
      emp.IntentID=IntentID
      emp.ActionID=ActionID
    }

    tmplintentactions.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertintentactions(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        IntentID :=r.FormValue("IntentID")
        ActionID :=r.FormValue("ActionID")
        insForm, err := db.Prepare("INSERT INTO intentactions(IntentID ,ActionID ) VALUE (? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        IntentID ,ActionID )  
         
         log.Println("INSERT:IntentID:" +IntentID +"ActionID:" +ActionID )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/intentactions", 301)
}

func Updateintentactions(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       IntentID:=r.FormValue("IntentID")
       ActionID:=r.FormValue("ActionID")
       IntentActionID:=r.FormValue("uIntentActionID")
       
       insForm, err :=db.Prepare("UPDATE intentactions SET IntentID=? ,ActionID=?  WHERE IntentActionID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        IntentID ,ActionID ,IntentActionID)
        log.Println("UPDATE:IntentID: " + IntentID + "ActionID: " + ActionID ,IntentActionID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/intentactions", 301)
}

func Deleteintentactions(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM intentactions WHERE IntentActionID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/intentactions", 301)
}
