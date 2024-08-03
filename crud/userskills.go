package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Userskills struct {

UserSkillID int

UserID int

SkillID int

ProficiencyLevel string

CreatedAt string

}
var tmpluserskills = template.Must(template.ParseGlob("userskills/*")) 

func Indexuserskills(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM userskills ORDER BY UserSkillID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Userskills{}
    res := []Userskills{}

    for selDB.Next() {
        
        var UserSkillID int
        
        var UserID int
        
        var SkillID int
        
        var ProficiencyLevel string
        
        var CreatedAt string
        

        err = selDB.Scan(
        &UserSkillID ,&UserID ,&SkillID ,&ProficiencyLevel ,&CreatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.UserSkillID=UserSkillID
       
       emp.UserID=UserID
       
       emp.SkillID=SkillID
       
       emp.ProficiencyLevel=ProficiencyLevel
       
       emp.CreatedAt=CreatedAt
       

        res = append(res, emp)
    }

    tmpluserskills.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showuserskills(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM userskills WHERE UserSkillID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Userskills{}

    for selDB.Next() {
       
       var UserSkillID int
       var UserID int
       var SkillID int
       var ProficiencyLevel string
       var CreatedAt string

        err = selDB.Scan(&UserSkillID ,&UserID ,&SkillID ,&ProficiencyLevel ,&CreatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.UserSkillID=UserSkillID
       emp.UserID=UserID
       emp.SkillID=SkillID
       emp.ProficiencyLevel=ProficiencyLevel
       emp.CreatedAt=CreatedAt
    }

    tmpluserskills.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newuserskills(w http.ResponseWriter, r *http.Request) {
    tmpluserskills.ExecuteTemplate(w, "New", nil)
}

func Edituserskills(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM userskills WHERE UserSkillID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Userskills{}

    for selDB.Next() {
      
      var UserSkillID int
      var UserID int
      var SkillID int
      var ProficiencyLevel string
      var CreatedAt string
        err = selDB.Scan(
        &UserSkillID ,&UserID ,&SkillID ,&ProficiencyLevel ,&CreatedAt )
        if err != nil {
            panic(err.Error())
        }

      
      emp.UserSkillID=UserSkillID
      emp.UserID=UserID
      emp.SkillID=SkillID
      emp.ProficiencyLevel=ProficiencyLevel
      emp.CreatedAt=CreatedAt
    }

    tmpluserskills.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertuserskills(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        UserID :=r.FormValue("UserID")
        SkillID :=r.FormValue("SkillID")
        ProficiencyLevel :=r.FormValue("ProficiencyLevel")
        CreatedAt :=r.FormValue("CreatedAt")
        insForm, err := db.Prepare("INSERT INTO userskills(UserID ,SkillID ,ProficiencyLevel ,CreatedAt ) VALUE (? ,? ,? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        UserID ,SkillID ,ProficiencyLevel ,CreatedAt )  
         
         log.Println("INSERT:UserID:" +UserID +"SkillID:" +SkillID +"ProficiencyLevel:" +ProficiencyLevel +"CreatedAt:" +CreatedAt )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/userskills", 301)
}

func Updateuserskills(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       UserID:=r.FormValue("UserID")
       SkillID:=r.FormValue("SkillID")
       ProficiencyLevel:=r.FormValue("ProficiencyLevel")
       CreatedAt:=r.FormValue("CreatedAt")
       UserSkillID:=r.FormValue("uUserSkillID")
       
       insForm, err :=db.Prepare("UPDATE userskills SET UserID=? ,SkillID=? ,ProficiencyLevel=? ,CreatedAt=?  WHERE UserSkillID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        UserID ,SkillID ,ProficiencyLevel ,CreatedAt ,UserSkillID)
        log.Println("UPDATE:UserID: " + UserID + "SkillID: " + SkillID + "ProficiencyLevel: " + ProficiencyLevel + "CreatedAt: " + CreatedAt ,UserSkillID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/userskills", 301)
}

func Deleteuserskills(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM userskills WHERE UserSkillID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/userskills", 301)
}
