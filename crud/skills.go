package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Skills struct {

SkillID int

SkillName string

Description string

CreatedAt string

UpdatedAt string

}
var tmplskills = template.Must(template.ParseGlob("skills/*")) 

func Indexskills(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM skills ORDER BY SkillID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Skills{}
    res := []Skills{}

    for selDB.Next() {
        
        var SkillID int
        
        var SkillName string
        
        var Description string
        
        var CreatedAt string
        
        var UpdatedAt string
        

        err = selDB.Scan(
        &SkillID ,&SkillName ,&Description ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.SkillID=SkillID
       
       emp.SkillName=SkillName
       
       emp.Description=Description
       
       emp.CreatedAt=CreatedAt
       
       emp.UpdatedAt=UpdatedAt
       

        res = append(res, emp)
    }

    tmplskills.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showskills(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM skills WHERE SkillID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Skills{}

    for selDB.Next() {
       
       var SkillID int
       var SkillName string
       var Description string
       var CreatedAt string
       var UpdatedAt string

        err = selDB.Scan(&SkillID ,&SkillName ,&Description ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.SkillID=SkillID
       emp.SkillName=SkillName
       emp.Description=Description
       emp.CreatedAt=CreatedAt
       emp.UpdatedAt=UpdatedAt
    }

    tmplskills.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newskills(w http.ResponseWriter, r *http.Request) {
    tmplskills.ExecuteTemplate(w, "New", nil)
}

func Editskills(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM skills WHERE SkillID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Skills{}

    for selDB.Next() {
      
      var SkillID int
      var SkillName string
      var Description string
      var CreatedAt string
      var UpdatedAt string
        err = selDB.Scan(
        &SkillID ,&SkillName ,&Description ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

      
      emp.SkillID=SkillID
      emp.SkillName=SkillName
      emp.Description=Description
      emp.CreatedAt=CreatedAt
      emp.UpdatedAt=UpdatedAt
    }

    tmplskills.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertskills(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        SkillName :=r.FormValue("SkillName")
        Description :=r.FormValue("Description")
        CreatedAt :=r.FormValue("CreatedAt")
        UpdatedAt :=r.FormValue("UpdatedAt")
        insForm, err := db.Prepare("INSERT INTO skills(SkillName ,Description ,CreatedAt ,UpdatedAt ) VALUE (? ,? ,? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        SkillName ,Description ,CreatedAt ,UpdatedAt )  
         
         log.Println("INSERT:SkillName:" +SkillName +"Description:" +Description +"CreatedAt:" +CreatedAt +"UpdatedAt:" +UpdatedAt )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/skills", 301)
}

func Updateskills(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       SkillName:=r.FormValue("SkillName")
       Description:=r.FormValue("Description")
       CreatedAt:=r.FormValue("CreatedAt")
       UpdatedAt:=r.FormValue("UpdatedAt")
       SkillID:=r.FormValue("uSkillID")
       
       insForm, err :=db.Prepare("UPDATE skills SET SkillName=? ,Description=? ,CreatedAt=? ,UpdatedAt=?  WHERE SkillID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        SkillName ,Description ,CreatedAt ,UpdatedAt ,SkillID)
        log.Println("UPDATE:SkillName: " + SkillName + "Description: " + Description + "CreatedAt: " + CreatedAt + "UpdatedAt: " + UpdatedAt ,SkillID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/skills", 301)
}

func Deleteskills(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM skills WHERE SkillID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/skills", 301)
}
