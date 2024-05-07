package models

// структура таблицы машин
type Test_Car struct {
	Car_id int          `json:"id" example:"12" gorm:"unique;primaryKey;autoIncrement"`
	RegNum string       `json:"regNum" example:""`
	Mark   string       `json:"mark" example:""`
	Model  string       `json:"model" example:"tesla"`
	Year   string       `json:"year" example:""`
	Owner  []Test_Owner `json:"owner" example:"" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;many2many:test_bounds;joinForeignKey:car_uid;JoinReferences:owner_uid"`
}
type Test_Owner struct {
	// gorm.Model
	Owner_id   int    `json:"id" example:"" gorm:"unique;primaryKey;autoIncrement"`
	Name       string `json:"name" example:"jamson"`
	Surname    string `json:"surname" example:""`
	Patronymic string `gorm:"column:patron" json:"patronymic" example:""`
}

type Searcher struct {
	Limit  string
	Offset string
	Sort   string
}
