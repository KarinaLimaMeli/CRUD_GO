package models

import(
	"github.com/aprendagolang/api-pgsql/db"
)

func Insert(todo Todo) (id int64, err error){
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.close()

	sql := `INSERT INTO todos(title, description, done)VALUE($1,$2,$3)RETURNING id`

	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)
	return
}