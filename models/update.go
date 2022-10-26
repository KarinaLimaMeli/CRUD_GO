package models

func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConncection()
	if err != nil {
		return 0, err
	}
	defer conn.close()
	res, err := conn.Exec(`UPDATE todos SET title=$1, description+$2, done=$3 WHERE id=$4`todo.Description, todo.Done,id)
	if err != nil{
		return 0, err
	}
	return res.RowsAffected()
}