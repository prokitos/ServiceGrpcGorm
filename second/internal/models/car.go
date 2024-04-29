package models

// структура таблицы машин
type Carer struct {
	Id     int    `json:"id" example:"12" gorm:"unique;primaryKey;autoIncrement"`
	RegNum string `json:"regNum" example:""`
	Mark   string `json:"mark" example:""`
	Model  string `json:"model" example:"tesla"`
	Year   string `json:"year" example:""`
}
