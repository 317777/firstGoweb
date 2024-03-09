package router

import (
	db "stuService/dbService"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Router() (router *gin.Engine){
	router = gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/getAllStudent", func(c *gin.Context) {
		students := db.GetAllStudent()
		c.JSON(200, students)	
	})

	router.POST("/addStudent", addStudent)

	router.DELETE("/deleteStudent/:id", deleteStudent)

	router.PUT("/updateStudent/:id", updateStudent)

	router.GET("/getStudentByID/:id", getStudentByID)
	
	return
}

// 添加学生
func addStudent(c *gin.Context) {
	var student db.Student
	if err := c.BindJSON(&student); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.AddStudent(student)
	c.JSON(200, gin.H{"message": "学生信息添加成功"})
}

// 删除学生
func deleteStudent(c *gin.Context) {
	studentID := c.Param("id")
	id, _ := strconv.Atoi(studentID)
	student := db.GetStudentByID(id)
	if student.ID == 0{
		c.JSON(404, gin.H{"error": "学生未找到"})
	}else{
		db.DeleteStudent(id)
		c.JSON(200, gin.H{"message": "删除成功"})
	}
}

// 修改学生信息
func updateStudent(c *gin.Context) {
	var student db.Student
	if err := c.BindJSON(&student); err != nil {
		c.JSON(400, gin.H{"error": "无效的请求体"})
		return
	}
	studentID := c.Param("id")
	id, _ := strconv.Atoi(studentID)
	hasStudent := db.GetStudentByID(id)
	if hasStudent.ID == 0{
		c.JSON(404, gin.H{"error": "学生未找到"})
	}else{
		student.ID = id
		db.UpdateStudent(student)
		c.JSON(200, gin.H{"message": "修改成功"})
	}
	
}

// 根据学生 ID 查找学生
func getStudentByID(c *gin.Context) {
	studentID := c.Param("id")
	id, _ := strconv.Atoi(studentID)
	student := db.GetStudentByID(id)
	if student.ID == 0 {
		c.JSON(404, gin.H{"error": "学生未找到"})
	}else{
		c.JSON(200, student)
	}
}
