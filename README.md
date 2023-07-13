# Example of Golang CRUD using MySQL from scratch

This repository show how to do database CRUD (create, read, update, delete) operations using Golang and MySQL. We are going to create an interface as databse frontend to handle these operations

## Install app

1. Clone the app
```
git clone https://github.com/hashi7412/crud-with-mysql.git <project_name>
```

2. Create database
```
DROP TABLE IF EXISTS `employee`;
CREATE TABLE `employee` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `city` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;
```

3. Install Golang packages
```
cd <project_name>
go get -u github.com/go-sql-driver/mysql
```

4. Run the program
```
go run .
```

5. Load the following URL

```
http://localhost:8080
```

## Guide the repository

### [main.go](https://github.com/hashi7412/crud-with-mysql/blob/main/main.go)

- dbConn()

This function connect to database and return a handler

Here is sql connection example:
```
sql.Open(<db_driver>, "<db_driver>:<db_pass>@<db_name>")
```

- Index()

This function executes `forms/Index.tmpl` to show table of employee data

- Show()

`Show` function executes `forms/Show.tmpl` to show details of an employee

- New()

This function executes `forms/New.tmpl` to show the interface to insert an employee detail

- Edit()

This function executes `forms/Edit.tmpl` to show the inerface to edit an employee detail

- Insert()

This function handles to insert data from `Edit` page

- Update()

This function handles to update data from `Edit` page

- Delete()

This function handles to delete data

- main()

The main function that is executed first implementes a handler function for multiple URL paths that provide functionalities.


## Conclusion

This repository implemented CRUD operations with MySQL server

Here is some repositories for your guide:

- [Hands-on Go](https://github.com/hashi7412/handson-go)
- [Token-based Authentication with MySQL](https://github.com/hashi7412/tokenbased-authentication)

Thank you for looking at this repository. 👋
