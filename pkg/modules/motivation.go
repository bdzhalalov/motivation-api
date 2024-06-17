package modules

type Motivation struct {
	ID         uint   `gorm:"primary_key;auto_increment"`
	Nickname   string `gorm:"type:varchar(64);not null"`
	Motivation string `gorm:"type:text;not null"`
}

func (Motivation) TableName() string {
	return "motivations"
}
