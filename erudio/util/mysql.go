
package util

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var sqldb *sql.DB

func InitializeMySql() { 
	// client = redis.NewClient(&redis.Options{
	// 	Addr:     "redis:6379",
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })

	// SetBookList(BookCollection)

	database, err := sql.Open("mysql", "user:user@(0.0.0.0:3306)/student_db?parseTime=true")
	if err != nil {
		fmt.Println("SQl Connection is not successful ^^^^^^^^^")
		log.Fatal(err)
	}
	// if err := database.Ping(); err != nil {
	// 	log.Fatal(err)
	// }
	sqldb = database
	fmt.Println("SQl Connection successful")
}

// func SetBookList(booklist *Collection) {
// 	err := client.Set("booklist",  booklist, 0).Err()
// 	if err != nil {
// 		log.Println("Couldn't save the booklist, Err: ", err)
// 	}
// }

// func GetBookList() *Collection {
// 	val, err := client.Get("booklist").Result()
// 	if err != nil {
// 		panic(err)
// 	}

// 	var books Collection
// 	err = json.Unmarshal([]byte(val), &books)
// 	if err != nil {
// 		panic(err)
// 	}



// 	return &books
// }


// type myJSON struct {
//     Array []string
// }

// func main() {
//     jsondat := &myJSON{Array: []string{"one", "two", "three"}}
//     encjson, _ := json.Marshal(jsondat)
//     fmt.Println(string(encjson))
// }

// func test()
// {
// 	jsondat := &Students{StudentList: []string{"one", "two", "three"}}
//     encjson, _ := json.Marshal(jsondat)
// }


// user := &User{Name: "Frank"}
//     b, err := json.Marshal(user)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     fmt.Println(string(b))




func GetStudentList() *StudentCollection {

	var StudentList []Student

	
	rows, err := sqldb.Query(`SELECT stu_id,name,father_name FROM student_db.std_profile`)
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}
	defer rows.Close()

	
	for rows.Next() {
		var u Student
		err := rows.Scan(&u.ID , &u.Name, &u.FatherName)
		if err != nil {
			log.Fatal(err)
		
		}

		// The standard algorithm to do this efficiently (which I assume append uses some variation of) is double the length of the underlying array 
		//  *every time you reallocate. This way, for n append operations, you perform only log(n) allocations, and you allocate approximately 2n memory.
		//  *Thus, the amount of space you allocate is linear in the number of append operations
		

		 StudentList = append(StudentList, u /* use the zero value of the element type */)

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Now we pass a slice of two pigeons
    data, _ := json.Marshal(StudentList)

	var students StudentCollection
	err = json.Unmarshal([]byte(data), &students)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%#v", users)
	return &students
}


func sqlConnect() (database *gorm.DB) {
	DBMS := "mysql"
	USER := "go_test"
	PASS := "password"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "go_database"
  
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	
	count := 0
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
	  for {
		if err == nil {
		  fmt.Println("")
		  break
		}
		fmt.Print(".")
		time.Sleep(time.Second)
		count++
		if count > 180 {
		  fmt.Println("")
		  panic(err)
		}
		db, err = gorm.Open(DBMS, CONNECT)
	  }
	}
  
	return db
  }


	







