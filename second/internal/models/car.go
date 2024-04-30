package models

// структура таблицы машин
type Test_Car struct {
	Car_id int          `json:"id" example:"12" gorm:"unique;primaryKey;autoIncrement"`
	RegNum string       `json:"regNum" example:""`
	Mark   string       `json:"mark" example:""`
	Model  string       `json:"model" example:"tesla"`
	Year   string       `json:"year" example:""`
	Owner  []Test_Owner `json:"owner" example:"" gorm:"many2many:test_bounds;"`
}
type Test_Owner struct {
	Owner_id   int    `json:"id" example:"" gorm:"unique;primaryKey;autoIncrement"`
	Name       string `json:"name" example:"jamson"`
	Surname    string `json:"surname" example:""`
	Patronymic string `json:"patronymic" example:""`
}

type Searcher struct {
	Limit  string
	Offset string
	Sort   string
}
