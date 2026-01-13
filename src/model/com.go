package model

import "time"

type Com struct {
	ID int `gorm:"column:commentid;primaryKey;autoIncrement" json:"id"`

	CommentText    string     `gorm:"column:commenttext;type:varchar(1000);not null" json:"comment_text"`
	UserIdentity   string     `gorm:"column:useridentity;type:varchar(300);not null" json:"user_identity"`
	ApprovedStatus int16      `gorm:"column:approvedstatus;not null" json:"approved_status"` // 0=pending, 1=approved, 2=rejected
	ApprovedBy     *int       `gorm:"column:approvedby" json:"approved_by,omitempty"`
	ApprovedAt     *time.Time `gorm:"column:approveddatetime" json:"approved_at,omitempty"`
	NewsIdentity   string     `gorm:"column:newsidentity;type:varchar(500);not null;index" json:"news_identity"`
	PublisherID    *int       `gorm:"column:publisherid" json:"publisher_id,omitempty"`
	ParentComID    *int       `gorm:"column:referencecommentid;index" json:"parent_com_id,omitempty"`
	Parent         *Com       `gorm:"foreignKey:ParentComID;references:ID;constraint:OnDelete:CASCADE" json:"parent,omitempty"`
	Replies        []Com      `gorm:"foreignKey:ParentComID" json:"replies,omitempty"`
	CreatedAt      time.Time  `gorm:"column:createdtime;autoCreateTime" json:"created_at"`
}

func (Com) TableName() string {
	return "com"
}
