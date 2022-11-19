package entity

type EntityUgcTermsAndConditions struct {
	Id                      int    `gorm:"type:int(11), primaryKey, column:id,omitempty"`
	TermsAndConditionsWhite string `gorm:"type:varchar(1000), column:terms_and_conditions_white"`
	TermsAndConditionsBlack string `gorm:"type:varchar(1000), column:terms_and_conditions_black"`
	CreatedAt               string `gorm:"type:varchar(100), column:created_at"`
	UpdatedAt               string `gorm:"type:varchar(100), column:updated_at"`
	CreatedBy               int    `gorm:"type:int(30), column:created_by"`
	UpdatedBy               int    `gorm:"type:int(30), column:updated_by"`
	IsActive                int    `gorm:"type:tinyint(1), column:is_active"`
}

func (EntityUgcTermsAndConditions) TableName() string {
	return "ugc_terms_and_conditions"
}
