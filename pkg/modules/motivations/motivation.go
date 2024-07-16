package motivations

type Motivation struct {
	ID         uint   `gorm:"primary_key;auto_increment" json:"id"`
	Nickname   string `gorm:"type:varchar(64);not null" json:"nickname"`
	Motivation string `gorm:"type:text;not null" json:"motivation"`
}

func (Motivation) TableName() string {
	return "motivations"
}
