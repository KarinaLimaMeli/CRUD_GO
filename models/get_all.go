package models

func GetAll() (todo []Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row,err := conn.Query(`SELECT * FROM todos`)
	if err != nil{
		return
	}
	for rows.Next() {
		var todo Todo
		err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}
		todos = append (todos, todo)
	}
	
	return
}