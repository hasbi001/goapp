package entities
type Users struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Email       string 	`json:"email"`
	Address 	string  `json:"address"`
	NoHp 		string  `json:"nohp"`
}