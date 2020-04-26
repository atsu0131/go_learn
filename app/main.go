package main

import (
    "database/sql"
    // "fmt"
    "net/http"
		"app/handler"
    _ "github.com/go-sql-driver/mysql"
    "github.com/labstack/echo"
)

// JSON用の構造体を定義
type Task struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func main() {
    e := echo.New()
		e.GET("/tasks", getTasks)
		// e.GET("/", handler.GetTasks)
		
    // e.GET("/users", handler.ListUser)
    // e.GET("/users/:id", handler.GetUser)
    // e.POST("/users", handler.CreateUser)
    // e.PUT("/users/:id", handler.UpdateUser)
    // e.DELETE("/users/:id", handler.DeleteUser)

    e.Start(":8082")
}

func dbConnect() *sql.DB {
    // DB接続処理
    db, err := sql.Open("mysql", "root:password@tcp(test_db)/test_todo")
    if err != nil {
        panic(err.Error())
    }
    return db
}

func getRows(db *sql.DB) *sql.Rows {
    rows, err := db.Query("SELECT id,name FROM tasks")
    if err != nil {
        panic(err.Error())
    }
    return rows
}

func getTasks(c echo.Context) error {
    // DB接続
    db := dbConnect()
    defer db.Close()
    rows := getRows(db)
    defer rows.Close()

    task := Task{}
    var results []Task
    for rows.Next() {
        err := rows.Scan(&task.ID, &task.Name)
        if err != nil {
            panic(err.Error())
        } else {
            results = append(results, task)
        }
    }

    return c.JSON(http.StatusOK, results)
}