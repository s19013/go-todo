package service

import (
	"context"
	"database/sql"
	"go-todo/model"
)

type TodoService struct {
	db *sql.DB
}

// なんか毎度毎度新しく作らないとメモリ関係?とかでバグるんだってさ｡
// なんかエディターがポインター使ったほうがいいよっていうからポインターを使う
func NewTODOService(db *sql.DB) *TodoService {
	return &TodoService{
		db: db,
	}
}

// なんかserviceはポインターを指定するのが通常なんだってさ
// なんか返り値ポインタにしたら短くかけるんだってさ､後パフォーマンスが良いんだとか
func (service *TodoService) Create(context context.Context, request model.CreateTodoRequest) (*model.CreateTodoResponse, error) {
	const query string = `INSERT INTO todos(subject, description) VALUES(?, ?)`

	// prepareを使ってinjectionとかを回避する
	preparedQuery, err := service.db.PrepareContext(context, query)
	if err != nil {
		return nil, err
	}

	// なんかしらないけどPrepareContext使ったらcloseしないといけないらしい
	defer preparedQuery.Close()

	// データベースに登録
	result, err := preparedQuery.ExecContext(context, request.Subject, request.Description)
	if err != nil {
		return nil, err
	}

	// 今保存したデータのIDを取り出す
	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	response := model.CreateTodoResponse{
		ID:      int(lastID),
		Message: "登録しました",
	}

	return &response, nil
}

func (service *TodoService) Get(context context.Context, request model.GetTodoRequest) (*model.GetTodoResponse, error) {

	const query string = `SELECT * FROM todos WHERE id = ?`

	preparedQuery, err := service.db.PrepareContext(context, query)
	if err != nil {
		return nil, err
	}

	defer preparedQuery.Close()

	// 1つだけ決め打ちで返してほしいならQueryRowContextらしい
	row := preparedQuery.QueryRowContext(context, request.ID)

	var todo model.Todo

	err2 := row.Scan(&todo.ID, &todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return nil, err2
	}

	response := model.GetTodoResponse{
		Todo: todo,
	}

	return &response, nil

}
