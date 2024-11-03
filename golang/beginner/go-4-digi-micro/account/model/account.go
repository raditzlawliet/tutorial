package model

type Account struct {
	AccountID int64  `json:"account_id" gorm:"primaryKey;autoIncrement;<-:false"`
	Name      string `json:"name"`
}

// func (Account) TableName() string {
// 	return "accounts"
// }
