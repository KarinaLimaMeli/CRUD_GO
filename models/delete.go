package models

func Delete(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConncection()
	if err != nil {
		return 0, err
	}
	defer conn.close()

	res, err := conn.Exec(`DELETE FROM todos WHERE id=$4`, id)
	if err != nil{
		return 0, err
	}
	return res.RowsAffected()
}