package handler

import (
    "net/http"
		"github.com/labstack/echo"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// JSON用の構造体を定義
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func MainPage() echo.HandlerFunc {
    return func(c echo.Context) error {     //c をいじって Request, Responseを色々する 
				// return c.String(http.StatusOK, "Hello World")
				
				jsonMap := map[string]string{
					"foo": "bar",
					"hoge": "fuga",
			}

			return c.JSON(http.StatusOK, jsonMap)

    }
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

func GetTasks() echo.HandlerFunc {
	return func(c echo.Context) error {
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


}