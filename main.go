package main

import (
	"CRUD-WITH-MYSQL-MAIN/crud"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		panic(err.Error())
	}

	log.Println("Server started on PORT 8000")

	http.HandleFunc("/achievements", crud.Indexachievements)
	http.HandleFunc("/showachievements", crud.Showachievements)

	http.HandleFunc("/newachievements", crud.Newachievements)
	http.HandleFunc("/editachievements", crud.Editachievements)
	http.HandleFunc("/insertachievements", crud.Insertachievements)
	http.HandleFunc("/updateachievements", crud.Updateachievements)
	http.HandleFunc("/deleteachievements", crud.Deleteachievements)

	http.HandleFunc("/employees", crud.Indexemployees)
	http.HandleFunc("/showemployees", crud.Showemployees)

	http.HandleFunc("/newemployees", crud.Newemployees)
	http.HandleFunc("/editemployees", crud.Editemployees)
	http.HandleFunc("/insertemployees", crud.Insertemployees)
	http.HandleFunc("/updateemployees", crud.Updateemployees)
	http.HandleFunc("/deleteemployees", crud.Deleteemployees)

	http.HandleFunc("/actions", crud.Indexactions)
	http.HandleFunc("/showactions", crud.Showactions)

	http.HandleFunc("/newactions", crud.Newactions)
	http.HandleFunc("/editactions", crud.Editactions)
	http.HandleFunc("/insertactions", crud.Insertactions)
	http.HandleFunc("/updateactions", crud.Updateactions)
	http.HandleFunc("/deleteactions", crud.Deleteactions)

	http.HandleFunc("/appointments", crud.Indexappointments)
	http.HandleFunc("/showappointments", crud.Showappointments)

	http.HandleFunc("/newappointments", crud.Newappointments)
	http.HandleFunc("/editappointments", crud.Editappointments)
	http.HandleFunc("/insertappointments", crud.Insertappointments)
	http.HandleFunc("/updateappointments", crud.Updateappointments)
	http.HandleFunc("/deleteappointments", crud.Deleteappointments)

	http.HandleFunc("/contexts", crud.Indexcontexts)
	http.HandleFunc("/showcontexts", crud.Showcontexts)

	http.HandleFunc("/newcontexts", crud.Newcontexts)
	http.HandleFunc("/editcontexts", crud.Editcontexts)
	http.HandleFunc("/insertcontexts", crud.Insertcontexts)
	http.HandleFunc("/updatecontexts", crud.Updatecontexts)
	http.HandleFunc("/deletecontexts", crud.Deletecontexts)

	http.HandleFunc("/conversations", crud.Indexconversations)
	http.HandleFunc("/showconversations", crud.Showconversations)

	http.HandleFunc("/newconversations", crud.Newconversations)
	http.HandleFunc("/editconversations", crud.Editconversations)
	http.HandleFunc("/insertconversations", crud.Insertconversations)
	http.HandleFunc("/updateconversations", crud.Updateconversations)
	http.HandleFunc("/deleteconversations", crud.Deleteconversations)

	http.HandleFunc("/entities", crud.Indexentities)
	http.HandleFunc("/showentities", crud.Showentities)

	http.HandleFunc("/newentities", crud.Newentities)
	http.HandleFunc("/editentities", crud.Editentities)
	http.HandleFunc("/insertentities", crud.Insertentities)
	http.HandleFunc("/updateentities", crud.Updateentities)
	http.HandleFunc("/deleteentities", crud.Deleteentities)

	http.HandleFunc("/goals", crud.Indexgoals)
	http.HandleFunc("/showgoals", crud.Showgoals)

	http.HandleFunc("/newgoals", crud.Newgoals)
	http.HandleFunc("/editgoals", crud.Editgoals)
	http.HandleFunc("/insertgoals", crud.Insertgoals)
	http.HandleFunc("/updategoals", crud.Updategoals)
	http.HandleFunc("/deletegoals", crud.Deletegoals)

	http.HandleFunc("/intentactions", crud.Indexintentactions)
	http.HandleFunc("/showintentactions", crud.Showintentactions)

	http.HandleFunc("/newintentactions", crud.Newintentactions)
	http.HandleFunc("/editintentactions", crud.Editintentactions)
	http.HandleFunc("/insertintentactions", crud.Insertintentactions)
	http.HandleFunc("/updateintentactions", crud.Updateintentactions)
	http.HandleFunc("/deleteintentactions", crud.Deleteintentactions)

	http.HandleFunc("/intententities", crud.Indexintententities)
	http.HandleFunc("/showintententities", crud.Showintententities)

	http.HandleFunc("/newintententities", crud.Newintententities)
	http.HandleFunc("/editintententities", crud.Editintententities)
	http.HandleFunc("/insertintententities", crud.Insertintententities)
	http.HandleFunc("/updateintententities", crud.Updateintententities)
	http.HandleFunc("/deleteintententities", crud.Deleteintententities)

	http.HandleFunc("/intents", crud.Indexintents)
	http.HandleFunc("/showintents", crud.Showintents)

	http.HandleFunc("/newintents", crud.Newintents)
	http.HandleFunc("/editintents", crud.Editintents)
	http.HandleFunc("/insertintents", crud.Insertintents)
	http.HandleFunc("/updateintents", crud.Updateintents)
	http.HandleFunc("/deleteintents", crud.Deleteintents)

	http.HandleFunc("/knowledgebase", crud.Indexknowledgebase)
	http.HandleFunc("/showknowledgebase", crud.Showknowledgebase)

	http.HandleFunc("/newknowledgebase", crud.Newknowledgebase)
	http.HandleFunc("/editknowledgebase", crud.Editknowledgebase)
	http.HandleFunc("/insertknowledgebase", crud.Insertknowledgebase)
	http.HandleFunc("/updateknowledgebase", crud.Updateknowledgebase)
	http.HandleFunc("/deleteknowledgebase", crud.Deleteknowledgebase)

	http.HandleFunc("/logs", crud.Indexlogs)
	http.HandleFunc("/showlogs", crud.Showlogs)

	http.HandleFunc("/newlogs", crud.Newlogs)
	http.HandleFunc("/editlogs", crud.Editlogs)
	http.HandleFunc("/insertlogs", crud.Insertlogs)
	http.HandleFunc("/updatelogs", crud.Updatelogs)
	http.HandleFunc("/deletelogs", crud.Deletelogs)

	http.HandleFunc("/messageentities", crud.Indexmessageentities)
	http.HandleFunc("/showmessageentities", crud.Showmessageentities)

	http.HandleFunc("/newmessageentities", crud.Newmessageentities)
	http.HandleFunc("/editmessageentities", crud.Editmessageentities)
	http.HandleFunc("/insertmessageentities", crud.Insertmessageentities)
	http.HandleFunc("/updatemessageentities", crud.Updatemessageentities)
	http.HandleFunc("/deletemessageentities", crud.Deletemessageentities)

	http.HandleFunc("/messages", crud.Indexmessages)
	http.HandleFunc("/showmessages", crud.Showmessages)

	http.HandleFunc("/newmessages", crud.Newmessages)
	http.HandleFunc("/editmessages", crud.Editmessages)
	http.HandleFunc("/insertmessages", crud.Insertmessages)
	http.HandleFunc("/updatemessages", crud.Updatemessages)
	http.HandleFunc("/deletemessages", crud.Deletemessages)

	http.HandleFunc("/notifications", crud.Indexnotifications)
	http.HandleFunc("/shownotifications", crud.Shownotifications)

	http.HandleFunc("/newnotifications", crud.Newnotifications)
	http.HandleFunc("/editnotifications", crud.Editnotifications)
	http.HandleFunc("/insertnotifications", crud.Insertnotifications)
	http.HandleFunc("/updatenotifications", crud.Updatenotifications)
	http.HandleFunc("/deletenotifications", crud.Deletenotifications)

	http.HandleFunc("/reminders", crud.Indexreminders)
	http.HandleFunc("/showreminders", crud.Showreminders)

	http.HandleFunc("/newreminders", crud.Newreminders)
	http.HandleFunc("/editreminders", crud.Editreminders)
	http.HandleFunc("/insertreminders", crud.Insertreminders)
	http.HandleFunc("/updatereminders", crud.Updatereminders)
	http.HandleFunc("/deletereminders", crud.Deletereminders)

	http.HandleFunc("/skills", crud.Indexskills)
	http.HandleFunc("/showskills", crud.Showskills)

	http.HandleFunc("/newskills", crud.Newskills)
	http.HandleFunc("/editskills", crud.Editskills)
	http.HandleFunc("/insertskills", crud.Insertskills)
	http.HandleFunc("/updateskills", crud.Updateskills)
	http.HandleFunc("/deleteskills", crud.Deleteskills)

	http.HandleFunc("/systemsettings", crud.Indexsystemsettings)
	http.HandleFunc("/showsystemsettings", crud.Showsystemsettings)

	http.HandleFunc("/newsystemsettings", crud.Newsystemsettings)
	http.HandleFunc("/editsystemsettings", crud.Editsystemsettings)
	http.HandleFunc("/insertsystemsettings", crud.Insertsystemsettings)
	http.HandleFunc("/updatesystemsettings", crud.Updatesystemsettings)
	http.HandleFunc("/deletesystemsettings", crud.Deletesystemsettings)

	http.HandleFunc("/userfeedback", crud.Indexuserfeedback)
	http.HandleFunc("/showuserfeedback", crud.Showuserfeedback)

	http.HandleFunc("/newuserfeedback", crud.Newuserfeedback)
	http.HandleFunc("/edituserfeedback", crud.Edituserfeedback)
	http.HandleFunc("/insertuserfeedback", crud.Insertuserfeedback)
	http.HandleFunc("/updateuserfeedback", crud.Updateuserfeedback)
	http.HandleFunc("/deleteuserfeedback", crud.Deleteuserfeedback)

	http.HandleFunc("/userhistory", crud.Indexuserhistory)
	http.HandleFunc("/showuserhistory", crud.Showuserhistory)

	http.HandleFunc("/newuserhistory", crud.Newuserhistory)
	http.HandleFunc("/edituserhistory", crud.Edituserhistory)
	http.HandleFunc("/insertuserhistory", crud.Insertuserhistory)
	http.HandleFunc("/updateuserhistory", crud.Updateuserhistory)
	http.HandleFunc("/deleteuserhistory", crud.Deleteuserhistory)

	http.HandleFunc("/usernotes", crud.Indexusernotes)
	http.HandleFunc("/showusernotes", crud.Showusernotes)

	http.HandleFunc("/newusernotes", crud.Newusernotes)
	http.HandleFunc("/editusernotes", crud.Editusernotes)
	http.HandleFunc("/insertusernotes", crud.Insertusernotes)
	http.HandleFunc("/updateusernotes", crud.Updateusernotes)
	http.HandleFunc("/deleteusernotes", crud.Deleteusernotes)

	http.HandleFunc("/userpreferences", crud.Indexuserpreferences)
	http.HandleFunc("/showuserpreferences", crud.Showuserpreferences)

	http.HandleFunc("/newuserpreferences", crud.Newuserpreferences)
	http.HandleFunc("/edituserpreferences", crud.Edituserpreferences)
	http.HandleFunc("/insertuserpreferences", crud.Insertuserpreferences)
	http.HandleFunc("/updateuserpreferences", crud.Updateuserpreferences)
	http.HandleFunc("/deleteuserpreferences", crud.Deleteuserpreferences)

	http.HandleFunc("/users", crud.Indexusers)
	http.HandleFunc("/showusers", crud.Showusers)

	http.HandleFunc("/newusers", crud.Newusers)
	http.HandleFunc("/editusers", crud.Editusers)
	http.HandleFunc("/insertusers", crud.Insertusers)
	http.HandleFunc("/updateusers", crud.Updateusers)
	http.HandleFunc("/deleteusers", crud.Deleteusers)

	http.HandleFunc("/usersessions", crud.Indexusersessions)
	http.HandleFunc("/showusersessions", crud.Showusersessions)

	http.HandleFunc("/newusersessions", crud.Newusersessions)
	http.HandleFunc("/editusersessions", crud.Editusersessions)
	http.HandleFunc("/insertusersessions", crud.Insertusersessions)
	http.HandleFunc("/updateusersessions", crud.Updateusersessions)
	http.HandleFunc("/deleteusersessions", crud.Deleteusersessions)

	http.HandleFunc("/userskills", crud.Indexuserskills)
	http.HandleFunc("/showuserskills", crud.Showuserskills)

	http.HandleFunc("/newuserskills", crud.Newuserskills)
	http.HandleFunc("/edituserskills", crud.Edituserskills)
	http.HandleFunc("/insertuserskills", crud.Insertuserskills)
	http.HandleFunc("/updateuserskills", crud.Updateuserskills)
	http.HandleFunc("/deleteuserskills", crud.Deleteuserskills)

	http.Handle(
		"/public/",
		http.StripPrefix(
			"/public/",
			http.FileServer(http.Dir("./public/")),
		),
	)

	http.ListenAndServe(":8080", nil)
}
