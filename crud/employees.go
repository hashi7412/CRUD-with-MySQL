package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Employees struct {

Id int

Name string

City string

}
var tmplemployees = template.Must(template.ParseGlob("employees/*")) 

func Indexemployees(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM employees ORDER BY id DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Employees{}
    res := []Employees{}

    for selDB.Next() {
        
        var id int
        
        var name string
        
        var city string
        

        err = selDB.Scan(
        &id ,&name ,&city )
        if err != nil {
            panic(err.Error())
        }

       
       emp.Id=id
       
       emp.Name=name
       
       emp.City=city
       

        res = append(res, emp)
    }

    tmplemployees.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showemployees(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM employees WHERE id=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Employees{}

    for selDB.Next() {
       
       var id int
       var name string
       var city string

        err = selDB.Scan(&id ,&name ,&city )
        if err != nil {
            panic(err.Error())
        }

       
       emp.Id=id
       emp.Name=name
       emp.City=city
    }

    tmplemployees.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newemployees(w http.ResponseWriter, r *http.Request) {
    tmplemployees.ExecuteTemplate(w, "New", nil)
}

func Editemployees(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM employees WHERE id=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Employees{}

    for selDB.Next() {
      
      var id int
      var name string
      var city string
        err = selDB.Scan(
        &id ,&name ,&city )
        if err != nil {
            panic(err.Error())
        }

      
      emp.Id=id
      emp.Name=name
      emp.City=city
    }

    tmplemployees.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertemployees(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        name :=r.FormValue("name")
        city :=r.FormValue("city")
        insForm, err := db.Prepare("INSERT INTO employees(name ,city ) VALUE (? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        name ,city )  
         
         log.Println("INSERT:Name:" +name +"City:" +city )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/employees", 301)
}

func Updateemployees(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       name:=r.FormValue("name")
       city:=r.FormValue("city")
       id:=r.FormValue("uid")
       
       insForm, err :=db.Prepare("UPDATE employees SET name=? ,city=?  WHERE id=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        name ,city ,id)
        log.Println("UPDATE:Name: " + name + "City: " + city ,id)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/employees", 301)
}

func Deleteemployees(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM employees WHERE id=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/employees", 301)
}
