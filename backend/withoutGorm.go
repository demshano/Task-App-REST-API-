// package main

// import (
//     "database/sql"
//     "net/http"
//     "strconv"

//     "github.com/gin-gonic/gin"
//     _ "github.com/mattn/go-sqlite3"
// )

// type Todo struct {
// 	ID   int    `json:"id"`
// 	Text string `json:text`
// 	Done bool   `json:done`
// }

// var db *sql.DB


// func main() {
//     var err error
//     db, err = sql.Open("sqlite3", "./todos.db")
//     if err != nil {
//         panic(err)
//     }
//     defer db.Close()

//     // Create todos table
//     _, err = db.Exec(`CREATE TABLE IF NOT EXISTS todos (
//         id INTEGER PRIMARY KEY AUTOINCREMENT,
//         text TEXT,
//         done BOOLEAN
//     )`)
//     if err != nil {
//         panic(err)
//     }

//     router := gin.Default()

//     router.POST("/todos", createTodo)
//     router.GET("/todos", getTodos)
//     router.PUT("/todos/:id", updateTodo)
//     router.DELETE("/todos/:id", deleteTodo)

//     router.Run(":8080")
// }


// //craete todo
// func createTodo(c *gin.Context) {
//     var todo Todo
//     if err := c.BindJSON(&todo); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }

//     result, err := db.Exec("INSERT INTO todos (text, done) VALUES (?, ?)", todo.Text, todo.Done)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     id, _ := result.LastInsertId()
//     todo.ID = int(id)

//     c.JSON(http.StatusCreated, todo)
// }


// //display todo
// func getTodos(c *gin.Context) {
//     rows, err := db.Query("SELECT id, text, done FROM todos")
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }
//     defer rows.Close()

//     var todos []Todo
//     for rows.Next() {
//         var todo Todo
//         if err := rows.Scan(&todo.ID, &todo.Text, &todo.Done); err != nil {
//             c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//             return
//         }
//         todos = append(todos, todo)
//     }

//     c.JSON(http.StatusOK, todos)
// }


// //update todo
// func updateTodo(c *gin.Context) {
//     id, _ := strconv.Atoi(c.Param("id"))
//     var todo Todo
//     if err := c.BindJSON(&todo); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }

//     _, err := db.Exec("UPDATE todos SET text = ?, done = ? WHERE id = ?", todo.Text, todo.Done, id)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     c.Status(http.StatusNoContent)
// }


// //delete todo
// func deleteTodo(c *gin.Context) {
//     id, _ := strconv.Atoi(c.Param("id"))

//     _, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     c.Status(http.StatusNoContent)
// }
