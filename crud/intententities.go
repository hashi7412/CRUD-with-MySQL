package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Intententities struct {

IntentEntityID int

IntentID int

EntityID int

}
var tmplintententities = template.Must(template.ParseGlob("intententities/*")) 

func Indexintententities(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM intententities ORDER BY IntentEntityID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Intententities{}
    res := []Intententities{}

    for selDB.Next() {
        
        var IntentEntityID int
        
        var IntentID int
        
        var EntityID int
        

        err = selDB.Scan(
        &IntentEntityID ,&IntentID ,&EntityID )
        if err != nil {
            panic(err.Error())
        }

       
       emp.IntentEntityID=IntentEntityID
       
       emp.IntentID=IntentID
       
       emp.EntityID=EntityID
       

        res = append(res, emp)
    }

    tmplintententities.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showintententities(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM intententities WHERE IntentEntityID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Intententities{}

    for selDB.Next() {
       
       var IntentEntityID int
       var IntentID int
       var EntityID int

        err = selDB.Scan(&IntentEntityID ,&IntentID ,&EntityID )
        if err != nil {
            panic(err.Error())
        }

       
       emp.IntentEntityID=IntentEntityID
       emp.IntentID=IntentID
       emp.EntityID=EntityID
    }

    tmplintententities.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newintententities(w http.ResponseWriter, r *http.Request) {
    tmplintententities.ExecuteTemplate(w, "New", nil)
}

func Editintententities(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM intententities WHERE IntentEntityID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Intententities{}

    for selDB.Next() {
      
      var IntentEntityID int
      var IntentID int
      var EntityID int
        err = selDB.Scan(
        &IntentEntityID ,&IntentID ,&EntityID )
        if err != nil {
            panic(err.Error())
        }

      
      emp.IntentEntityID=IntentEntityID
      emp.IntentID=IntentID
      emp.EntityID=EntityID
    }

    tmplintententities.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertintententities(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        IntentID :=r.FormValue("IntentID")
        EntityID :=r.FormValue("EntityID")
        insForm, err := db.Prepare("INSERT INTO intententities(IntentID ,EntityID ) VALUE (? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        IntentID ,EntityID )  
         
         log.Println("INSERT:IntentID:" +IntentID +"EntityID:" +EntityID )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/intententities", 301)
}

func Updateintententities(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       IntentID:=r.FormValue("IntentID")
       EntityID:=r.FormValue("EntityID")
       IntentEntityID:=r.FormValue("uIntentEntityID")
       
       insForm, err :=db.Prepare("UPDATE intententities SET IntentID=? ,EntityID=?  WHERE IntentEntityID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        IntentID ,EntityID ,IntentEntityID)
        log.Println("UPDATE:IntentID: " + IntentID + "EntityID: " + EntityID ,IntentEntityID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/intententities", 301)
}

func Deleteintententities(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM intententities WHERE IntentEntityID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/intententities", 301)
}
