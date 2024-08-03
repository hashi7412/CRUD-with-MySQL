package crud

import "database/sql"

func dbConn() (db *sql.DB) {
	// dbDriver := viper.GetString("database.driver")
	// dbUser := viper.GetString("database.username")
	// dbPass := viper.GetString("database.password")
	// dbName := viper.GetString("database.dbName")

	// db, err := sql.Open(
	// 	dbDriver,
	// 	dbUser+":"+dbPass+"@/"+dbName,
	// )

	// if err != nil {
	// 	panic(err.Error())
	// }

	// return db
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "11109876543210Tolgahan."
	dbName := "ai_assistant"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
