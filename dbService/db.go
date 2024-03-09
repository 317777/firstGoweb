package dbservice

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
    ID   int
    Name string
    Age  string
}

func GetAllStudent()(students []Student){
	db, err := sql.Open("mysql", "root:123456@/db01?charset=utf8")
    checkErr(err)
	rows, err := db.Query("SELECT * FROM student")
	checkErr(err)
	defer rows.Close()
	checkErr(err)
	for rows.Next() {
        var student Student
        err := rows.Scan(&student.ID, &student.Name, &student.Age)
        checkErr(err)
		students = append(students, student)
    }
	return
}

func GetStudentByID(studentID int) (student Student) {
	db, err := sql.Open("mysql", "root:123456@/db01?charset=utf8")
	checkErr(err)
	defer db.Close()
	db.QueryRow("SELECT id, name, age FROM student WHERE id = ?", studentID).Scan(&student.ID, &student.Name, &student.Age)
	checkErr(err)
	return
}

func AddStudent(student Student) {
	db, err := sql.Open("mysql", "root:123456@/db01?charset=utf8")
	checkErr(err)
	defer db.Close()

	_, err = db.Exec("INSERT INTO student (name, age) VALUES (?, ?)", student.Name, student.Age)
	checkErr(err)
}

func UpdateStudent(student Student) {
	db, err := sql.Open("mysql", "root:123456@/db01?charset=utf8")
	checkErr(err)
	defer db.Close()

	_, err = db.Exec("UPDATE student SET name = ?, age = ? WHERE id = ?", student.Name, student.Age, student.ID)
	checkErr(err)
}

func DeleteStudent(studentID int) {
	db, err := sql.Open("mysql", "root:123456@/db01?charset=utf8")
	checkErr(err)
	defer db.Close()

	_, err = db.Exec("DELETE FROM student WHERE id = ?", studentID)
	checkErr(err)
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
