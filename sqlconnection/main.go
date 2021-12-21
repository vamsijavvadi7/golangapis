package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	

	_ "github.com/go-sql-driver/mysql"
)
var str string="vamsijavvadi@gmail.com"
func main() {
    fmt.Println(Fetch("select email,password,personid,role from person"))
   

    // if there is an error opening the connection, handle it
  

    // defer the close till after the main function has finished
    // executing

//insert, err := db.Query("INSERT INTO PERSON VALUES (1,'vamsi','krishna','vamsijavvadi','Ilovemydad7@','student',9,'vamsijavvadi@gmail.com')")
// insert, err := db.Query("INSERT INTO USER VALUES ('vamsijavvadi','vamsijavvadi@gmail.com')")
//  if err != nil {
//         panic(err.Error())
//   }
    // if there is an error inserting, handle it
   
    // be careful deferring Queries if you are using transactions



 //   defer insert.Close()
}
   
    func Fetch(query string) (string) {

    type User struct{
        Mail string  // <-- CHANGED THIS LINE
        Pass string
        Id string // <-- CHANGED THIS LINE
        Role string
    }
 db, err := sql.Open("mysql", "root:Ilovemydad7@@tcp(127.0.0.1:3306)/mydb")
   if err != nil {
        panic(err.Error())
    }   
 rows, err := db.Query(query)
 if err != nil {
  
        panic(err.Error())
       
    }  
    defer db.Close()
    users := make([]*User, 0)
    for rows.Next() {
      user := new(User)
      err := rows.Scan(&user.Mail, &user.Pass,&user.Id,&user.Role)
      if err != nil {
      panic(err)
    }
      users = append(users, user)
    }
defer rows.Close()
var res string


 for _, elem := range users {
        if elem.Mail==str{
        res=elem.Role
        }
 }
 fmt.Println(res)


    return(ToJSON(users)) // <-- CHANGED THIS LINE
}

func ToJSON(obj interface{}) (string) {

    res, err := json.Marshal(obj)

   
    if err != nil {
      panic("error with json serialization " + err.Error())
    }
    return string(res)
}