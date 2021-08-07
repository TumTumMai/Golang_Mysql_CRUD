package model

type Item struct {
	Id    uint   `gorm:"primary_key" json:"id"`
	Name  string `json:"name" validate:"required"` ////เเท็กเอาไว้ระรับค่า เเละส่งค่าตามที่ระบุ "เป็นkeyของjson"เเละ"เป็นชื่อคอลั่มในdatabase"
	Price int    `json:"price" validate:"required"`
}
