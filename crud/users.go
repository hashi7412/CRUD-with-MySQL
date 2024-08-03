package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Users struct {

UserID int

UserName string

UserEmail string

UserPassword string

UserRole string

CreatedAt string

UpdatedAt string

Name string

Email string

Username string

Password string

USER string

CURRENT_CONNECTIONS int

TOTAL_CONNECTIONS int

MAX_SESSION_CONTROLLED_MEMORY int

MAX_SESSION_TOTAL_MEMORY int

}
var tmplusers = template.Must(template.ParseGlob("users/*")) 

func Indexusers(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM users ORDER BY UserID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Users{}
    res := []Users{}

    for selDB.Next() {
        
        var UserID int
        
        var UserName string
        
        var UserEmail string
        
        var UserPassword string
        
        var UserRole string
        
        var CreatedAt string
        
        var UpdatedAt string
        
        var name string
        
        var email string
        
        var username string
        
        var password string
        
        var USER string
        
        var CURRENT_CONNECTIONS int
        
        var TOTAL_CONNECTIONS int
        
        var MAX_SESSION_CONTROLLED_MEMORY int
        
        var MAX_SESSION_TOTAL_MEMORY int
        

        err = selDB.Scan(
        &UserID ,&UserName ,&UserEmail ,&UserPassword ,&UserRole ,&CreatedAt ,&UpdatedAt ,&name ,&email ,&username ,&password ,&USER ,&CURRENT_CONNECTIONS ,&TOTAL_CONNECTIONS ,&MAX_SESSION_CONTROLLED_MEMORY ,&MAX_SESSION_TOTAL_MEMORY )
        if err != nil {
            panic(err.Error())
        }

       
       emp.UserID=UserID
       
       emp.UserName=UserName
       
       emp.UserEmail=UserEmail
       
       emp.UserPassword=UserPassword
       
       emp.UserRole=UserRole
       
       emp.CreatedAt=CreatedAt
       
       emp.UpdatedAt=UpdatedAt
       
       emp.Name=name
       
       emp.Email=email
       
       emp.Username=username
       
       emp.Password=password
       
       emp.USER=USER
       
       emp.CURRENT_CONNECTIONS=CURRENT_CONNECTIONS
       
       emp.TOTAL_CONNECTIONS=TOTAL_CONNECTIONS
       
       emp.MAX_SESSION_CONTROLLED_MEMORY=MAX_SESSION_CONTROLLED_MEMORY
       
       emp.MAX_SESSION_TOTAL_MEMORY=MAX_SESSION_TOTAL_MEMORY
       

        res = append(res, emp)
    }

    tmplusers.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showusers(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM users WHERE UserID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Users{}

    for selDB.Next() {
       
       var UserID int
       var UserName string
       var UserEmail string
       var UserPassword string
       var UserRole string
       var CreatedAt string
       var UpdatedAt string
       var name string
       var email string
       var username string
       var password string
       var USER string
       var CURRENT_CONNECTIONS int
       var TOTAL_CONNECTIONS int
       var MAX_SESSION_CONTROLLED_MEMORY int
       var MAX_SESSION_TOTAL_MEMORY int

        err = selDB.Scan(&UserID ,&UserName ,&UserEmail ,&UserPassword ,&UserRole ,&CreatedAt ,&UpdatedAt ,&name ,&email ,&username ,&password ,&USER ,&CURRENT_CONNECTIONS ,&TOTAL_CONNECTIONS ,&MAX_SESSION_CONTROLLED_MEMORY ,&MAX_SESSION_TOTAL_MEMORY )
        if err != nil {
            panic(err.Error())
        }

       
       emp.UserID=UserID
       emp.UserName=UserName
       emp.UserEmail=UserEmail
       emp.UserPassword=UserPassword
       emp.UserRole=UserRole
       emp.CreatedAt=CreatedAt
       emp.UpdatedAt=UpdatedAt
       emp.Name=name
       emp.Email=email
       emp.Username=username
       emp.Password=password
       emp.USER=USER
       emp.CURRENT_CONNECTIONS=CURRENT_CONNECTIONS
       emp.TOTAL_CONNECTIONS=TOTAL_CONNECTIONS
       emp.MAX_SESSION_CONTROLLED_MEMORY=MAX_SESSION_CONTROLLED_MEMORY
       emp.MAX_SESSION_TOTAL_MEMORY=MAX_SESSION_TOTAL_MEMORY
    }

    tmplusers.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newusers(w http.ResponseWriter, r *http.Request) {
    tmplusers.ExecuteTemplate(w, "New", nil)
}

func Editusers(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM users WHERE UserID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Users{}

    for selDB.Next() {
      
      var UserID int
      var UserName string
      var UserEmail string
      var UserPassword string
      var UserRole string
      var CreatedAt string
      var UpdatedAt string
      var name string
      var email string
      var username string
      var password string
      var USER string
      var CURRENT_CONNECTIONS int
      var TOTAL_CONNECTIONS int
      var MAX_SESSION_CONTROLLED_MEMORY int
      var MAX_SESSION_TOTAL_MEMORY int
        err = selDB.Scan(
        &UserID ,&UserName ,&UserEmail ,&UserPassword ,&UserRole ,&CreatedAt ,&UpdatedAt ,&name ,&email ,&username ,&password ,&USER ,&CURRENT_CONNECTIONS ,&TOTAL_CONNECTIONS ,&MAX_SESSION_CONTROLLED_MEMORY ,&MAX_SESSION_TOTAL_MEMORY )
        if err != nil {
            panic(err.Error())
        }

      
      emp.UserID=UserID
      emp.UserName=UserName
      emp.UserEmail=UserEmail
      emp.UserPassword=UserPassword
      emp.UserRole=UserRole
      emp.CreatedAt=CreatedAt
      emp.UpdatedAt=UpdatedAt
      emp.Name=name
      emp.Email=email
      emp.Username=username
      emp.Password=password
      emp.USER=USER
      emp.CURRENT_CONNECTIONS=CURRENT_CONNECTIONS
      emp.TOTAL_CONNECTIONS=TOTAL_CONNECTIONS
      emp.MAX_SESSION_CONTROLLED_MEMORY=MAX_SESSION_CONTROLLED_MEMORY
      emp.MAX_SESSION_TOTAL_MEMORY=MAX_SESSION_TOTAL_MEMORY
    }

    tmplusers.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertusers(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        UserName :=r.FormValue("UserName")
        UserEmail :=r.FormValue("UserEmail")
        UserPassword :=r.FormValue("UserPassword")
        UserRole :=r.FormValue("UserRole")
        CreatedAt :=r.FormValue("CreatedAt")
        UpdatedAt :=r.FormValue("UpdatedAt")
        name :=r.FormValue("name")
        email :=r.FormValue("email")
        username :=r.FormValue("username")
        password :=r.FormValue("password")
        USER :=r.FormValue("USER")
        CURRENT_CONNECTIONS :=r.FormValue("CURRENT_CONNECTIONS")
        TOTAL_CONNECTIONS :=r.FormValue("TOTAL_CONNECTIONS")
        MAX_SESSION_CONTROLLED_MEMORY :=r.FormValue("MAX_SESSION_CONTROLLED_MEMORY")
        MAX_SESSION_TOTAL_MEMORY :=r.FormValue("MAX_SESSION_TOTAL_MEMORY")
        insForm, err := db.Prepare("INSERT INTO users(UserName ,UserEmail ,UserPassword ,UserRole ,CreatedAt ,UpdatedAt ,name ,email ,username ,password ,USER ,CURRENT_CONNECTIONS ,TOTAL_CONNECTIONS ,MAX_SESSION_CONTROLLED_MEMORY ,MAX_SESSION_TOTAL_MEMORY ) VALUE (? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        UserName ,UserEmail ,UserPassword ,UserRole ,CreatedAt ,UpdatedAt ,name ,email ,username ,password ,USER ,CURRENT_CONNECTIONS ,TOTAL_CONNECTIONS ,MAX_SESSION_CONTROLLED_MEMORY ,MAX_SESSION_TOTAL_MEMORY )  
         
         log.Println("INSERT:UserName:" +UserName +"UserEmail:" +UserEmail +"UserPassword:" +UserPassword +"UserRole:" +UserRole +"CreatedAt:" +CreatedAt +"UpdatedAt:" +UpdatedAt +"Name:" +name +"Email:" +email +"Username:" +username +"Password:" +password +"USER:" +USER +"CURRENT_CONNECTIONS:" +CURRENT_CONNECTIONS +"TOTAL_CONNECTIONS:" +TOTAL_CONNECTIONS +"MAX_SESSION_CONTROLLED_MEMORY:" +MAX_SESSION_CONTROLLED_MEMORY +"MAX_SESSION_TOTAL_MEMORY:" +MAX_SESSION_TOTAL_MEMORY )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/users", 301)
}

func Updateusers(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       UserName:=r.FormValue("UserName")
       UserEmail:=r.FormValue("UserEmail")
       UserPassword:=r.FormValue("UserPassword")
       UserRole:=r.FormValue("UserRole")
       CreatedAt:=r.FormValue("CreatedAt")
       UpdatedAt:=r.FormValue("UpdatedAt")
       name:=r.FormValue("name")
       email:=r.FormValue("email")
       username:=r.FormValue("username")
       password:=r.FormValue("password")
       USER:=r.FormValue("USER")
       CURRENT_CONNECTIONS:=r.FormValue("CURRENT_CONNECTIONS")
       TOTAL_CONNECTIONS:=r.FormValue("TOTAL_CONNECTIONS")
       MAX_SESSION_CONTROLLED_MEMORY:=r.FormValue("MAX_SESSION_CONTROLLED_MEMORY")
       MAX_SESSION_TOTAL_MEMORY:=r.FormValue("MAX_SESSION_TOTAL_MEMORY")
       UserID:=r.FormValue("uUserID")
       
       insForm, err :=db.Prepare("UPDATE users SET UserName=? ,UserEmail=? ,UserPassword=? ,UserRole=? ,CreatedAt=? ,UpdatedAt=? ,name=? ,email=? ,username=? ,password=? ,USER=? ,CURRENT_CONNECTIONS=? ,TOTAL_CONNECTIONS=? ,MAX_SESSION_CONTROLLED_MEMORY=? ,MAX_SESSION_TOTAL_MEMORY=?  WHERE UserID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        UserName ,UserEmail ,UserPassword ,UserRole ,CreatedAt ,UpdatedAt ,name ,email ,username ,password ,USER ,CURRENT_CONNECTIONS ,TOTAL_CONNECTIONS ,MAX_SESSION_CONTROLLED_MEMORY ,MAX_SESSION_TOTAL_MEMORY ,UserID)
        log.Println("UPDATE:UserName: " + UserName + "UserEmail: " + UserEmail + "UserPassword: " + UserPassword + "UserRole: " + UserRole + "CreatedAt: " + CreatedAt + "UpdatedAt: " + UpdatedAt + "Name: " + name + "Email: " + email + "Username: " + username + "Password: " + password + "USER: " + USER + "CURRENT_CONNECTIONS: " + CURRENT_CONNECTIONS + "TOTAL_CONNECTIONS: " + TOTAL_CONNECTIONS + "MAX_SESSION_CONTROLLED_MEMORY: " + MAX_SESSION_CONTROLLED_MEMORY + "MAX_SESSION_TOTAL_MEMORY: " + MAX_SESSION_TOTAL_MEMORY ,UserID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/users", 301)
}

func Deleteusers(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM users WHERE UserID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/users", 301)
}
