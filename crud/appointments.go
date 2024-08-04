package crud

import (
    "log"
    "net/http"
    "text/template"
    
)

type Appointments struct {

AppointmentID int

UserID int

AppointmentTitle string

AppointmentDetails string

AppointmentDate string

CreatedAt string

UpdatedAt string

}
var tmplappointments = template.Must(template.ParseGlob("appointments/*")) 

func Indexappointments(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM appointments ORDER BY AppointmentID DESC")

    if err != nil {
        panic(err.Error())
    }

    emp := Appointments{}
    res := []Appointments{}

    for selDB.Next() {
        
        var AppointmentID int
        
        var UserID int
        
        var AppointmentTitle string
        
        var AppointmentDetails string
        
        var AppointmentDate string
        
        var CreatedAt string
        
        var UpdatedAt string
        

        err = selDB.Scan(
        &AppointmentID ,&UserID ,&AppointmentTitle ,&AppointmentDetails ,&AppointmentDate ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.AppointmentID=AppointmentID
       
       emp.UserID=UserID
       
       emp.AppointmentTitle=AppointmentTitle
       
       emp.AppointmentDetails=AppointmentDetails
       
       emp.AppointmentDate=AppointmentDate
       
       emp.CreatedAt=CreatedAt
       
       emp.UpdatedAt=UpdatedAt
       

        res = append(res, emp)
    }

    tmplappointments.ExecuteTemplate(w, "Index", res)

    // defer db.Close()
}

func Showappointments(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM appointments WHERE AppointmentID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Appointments{}

    for selDB.Next() {
       
       var AppointmentID int
       var UserID int
       var AppointmentTitle string
       var AppointmentDetails string
       var AppointmentDate string
       var CreatedAt string
       var UpdatedAt string

        err = selDB.Scan(&AppointmentID ,&UserID ,&AppointmentTitle ,&AppointmentDetails ,&AppointmentDate ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

       
       emp.AppointmentID=AppointmentID
       emp.UserID=UserID
       emp.AppointmentTitle=AppointmentTitle
       emp.AppointmentDetails=AppointmentDetails
       emp.AppointmentDate=AppointmentDate
       emp.CreatedAt=CreatedAt
       emp.UpdatedAt=UpdatedAt
    }

    tmplappointments.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func Newappointments(w http.ResponseWriter, r *http.Request) {
    tmplappointments.ExecuteTemplate(w, "New", nil)
}

func Editappointments(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")

    selDB, err := db.Query("SELECT * FROM appointments WHERE AppointmentID=?", nId)

    if err != nil {
        panic(err.Error())
    }

    emp := Appointments{}

    for selDB.Next() {
      
      var AppointmentID int
      var UserID int
      var AppointmentTitle string
      var AppointmentDetails string
      var AppointmentDate string
      var CreatedAt string
      var UpdatedAt string
        err = selDB.Scan(
        &AppointmentID ,&UserID ,&AppointmentTitle ,&AppointmentDetails ,&AppointmentDate ,&CreatedAt ,&UpdatedAt )
        if err != nil {
            panic(err.Error())
        }

      
      emp.AppointmentID=AppointmentID
      emp.UserID=UserID
      emp.AppointmentTitle=AppointmentTitle
      emp.AppointmentDetails=AppointmentDetails
      emp.AppointmentDate=AppointmentDate
      emp.CreatedAt=CreatedAt
      emp.UpdatedAt=UpdatedAt
    }

    tmplappointments.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insertappointments(w http.ResponseWriter, r *http.Request) {
    db := dbConn()

    if r.Method == "POST" {
        
        UserID :=r.FormValue("UserID")
        AppointmentTitle :=r.FormValue("AppointmentTitle")
        AppointmentDetails :=r.FormValue("AppointmentDetails")
        AppointmentDate :=r.FormValue("AppointmentDate")
        CreatedAt :=r.FormValue("CreatedAt")
        UpdatedAt :=r.FormValue("UpdatedAt")
        insForm, err := db.Prepare("INSERT INTO appointments(UserID ,AppointmentTitle ,AppointmentDetails ,AppointmentDate ,CreatedAt ,UpdatedAt ) VALUE (? ,? ,? ,? ,? ,? )")
        
        if err != nil { 
            panic(err.Error()) 
        } 
        
        insForm.Exec(
        UserID ,AppointmentTitle ,AppointmentDetails ,AppointmentDate ,CreatedAt ,UpdatedAt )  
         
         log.Println("INSERT:UserID:" +UserID +"AppointmentTitle:" +AppointmentTitle +"AppointmentDetails:" +AppointmentDetails +"AppointmentDate:" +AppointmentDate +"CreatedAt:" +CreatedAt +"UpdatedAt:" +UpdatedAt )
       
    }

    defer db.Close()
    http.Redirect(w, r, "/appointments", 301)
}

func Updateappointments(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
       
       UserID:=r.FormValue("UserID")
       AppointmentTitle:=r.FormValue("AppointmentTitle")
       AppointmentDetails:=r.FormValue("AppointmentDetails")
       AppointmentDate:=r.FormValue("AppointmentDate")
       CreatedAt:=r.FormValue("CreatedAt")
       UpdatedAt:=r.FormValue("UpdatedAt")
       AppointmentID:=r.FormValue("uAppointmentID")
       
       insForm, err :=db.Prepare("UPDATE appointments SET UserID=? ,AppointmentTitle=? ,AppointmentDetails=? ,AppointmentDate=? ,CreatedAt=? ,UpdatedAt=?  WHERE AppointmentID=?")
       
       
        if err != nil {
            panic(err.Error())
        }

        insForm.Exec(
        UserID ,AppointmentTitle ,AppointmentDetails ,AppointmentDate ,CreatedAt ,UpdatedAt ,AppointmentID)
        log.Println("UPDATE:UserID: " + UserID + "AppointmentTitle: " + AppointmentTitle + "AppointmentDetails: " + AppointmentDetails + "AppointmentDate: " + AppointmentDate + "CreatedAt: " + CreatedAt + "UpdatedAt: " + UpdatedAt ,AppointmentID)
        
    }

    defer db.Close()
    http.Redirect(w, r, "/appointments", 301)
}

func Deleteappointments(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")

    delForm, err := db.Prepare("DELETE FROM appointments WHERE AppointmentID=?")

    if err != nil {
        panic(err.Error())
    }

    delForm.Exec(emp)
    log.Println("DELETE")

    defer db.Close()
    http.Redirect(w, r, "/appointments", 301)
}
