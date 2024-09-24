package main

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Id    int    `json:"id, omitempty"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score, omitEmpty"`
}

var students map[int]Student
var lastId int

func SetupHandlers(g *gin.Engine) {
	g.GET("./students", GetStudentsHandler)
	g.GET("./student/:id", GetStudentHandler)
	g.POST("./student", PostStudentHandler)
	g.DELETE("./student/:id", DeleteStudentHandler)

	students = make(map[int]Student)
	students[0] = Student{1, "aaa", 16, 87}
	students[1] = Student{2, "bbb", 18, 98}

	lastId = 2
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

func GetStudentsHandler(c *gin.Context) {
	list := make(Students, 0)
	for _, student := range students {
		list = append(list, student)
	}

	sort.Sort(list)
	c.JSON(http.StatusOK, list)
}

func GetStudentHandler(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	student, ok := students[id]

	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, student)
}

func PostStudentHandler(c *gin.Context) {
	var student Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	lastId++

	student.Id = lastId
	students[lastId] = student
	c.String(http.StatusCreated, "Success to add id: %d", lastId)
}

func DeleteStudentHandler(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	delete(students, id)
	c.String(http.StatusOK, "Success to delete id: %d", id)
}

func main() {
	r := gin.Default()
	SetupHandlers(r)
	r.Run(":8080")
}
