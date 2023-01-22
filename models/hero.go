package models

type Hero struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	AlterEgo    string `json:"alterEgo"`
	Description string `json:"description"`
}

type CreateHeroInput struct {
	Name        string `json:"name" binding:"required"`
	AlterEgo    string `json:"alterEgo" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateHeroInput struct {
	Name        string `json:"name"`
	AlterEgo    string `json:"alterEgo"`
	Description string `json:"description"`
}
