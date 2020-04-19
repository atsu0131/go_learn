package main

import (
      //  "net/http"
			 "github.com/labstack/echo"
			 "github.com/labstack/echo/middleware"
			 "app/handler"
)

func main () {
		e := echo.New()
		 
		// Middleware
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		 
		//  // ルーティング
		e.GET("/", handler.MainPage())
    // Routes
    // e.GET("/users", handler.ListUser)
    // e.GET("/users/:id", handler.GetUser)
    e.POST("/users", handler.CreateUser)
    // e.PUT("/users/:id", handler.UpdateUser)
    // e.DELETE("/users/:id", handler.DeleteUser)

	//  // サーバー起動
	// e.Startの中はdocker-composeのgoコンテナで設定したportsを指定してください。
		e.Logger.Fatal(e.Start(":8082"))
}

// package main

// import (
//     "database/sql"
//     "fmt"
//     "net/http"

//     _ "github.com/go-sql-driver/mysql"
//     "github.com/labstack/echo"
// )

// // JSON用の構造体を定義
// type Task struct {
//     ID   int    `json:"id"`
//     Name string `json:"name"`
// }

// func main() {
//     e := echo.New()
//     // e.GET("/tasks", getTasks)
// 		// e.Start(":8082")
//     //  e.GET("/", func(c echo.Context) error {
//     //     // return c.String(http.StatusOK, "Hello, World!")

//     //  })
// 		//  e.Logger.Fatal(e.Start(":8082"))
		
// 		e.GET("/", func(c echo.Context) error {
// 			return c.String(http.StatusOK, "Hello, World!")
// 		})
// 		e.Logger.Fatal(e.Start(":8082"))

    
// }

// func dbConnect() *sql.DB {
//     // DB接続処理
//     db, err := sql.Open("mysql", "root:password@tcp(test_db)/test_todo")
//     if err != nil {
//         panic(err.Error())
//     }
//     return db
// }

// func getRows(db *sql.DB) *sql.Rows {
//     rows, err := db.Query("SELECT id,name FROM tasks")
//     if err != nil {
//         panic(err.Error())
//     }
//     return rows
// }

// func getTasks(c echo.Context) error {
//     // DB接続
//     db := dbConnect()
//     defer db.Close()
//     rows := getRows(db)
//     defer rows.Close()

//     task := Task{}
//     var results []Task
//     for rows.Next() {
//         err := rows.Scan(&task.ID, &task.Name)
//         if err != nil {
//             panic(err.Error())
//         } else {
//             results = append(results, task)
//         }
//     }

//     return c.JSON(http.StatusOK, results)
// }