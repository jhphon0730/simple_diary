package model
import (
	"gorm.io/gorm"
)

type CoupleInvitation struct {
	gorm.Model
	SenderID   uint   `json:"sender_id" gorm:"not null"` // 초대한 유저
	ReceiverID *uint  `json:"receiver_id,omitempty"`     // 초대받은 유저
	InviteCode string `json:"invite_code" gorm:"unique;not null"`
	Status     string `json:"status" gorm:"default:'pending'"` // "pending", "accepted"
}

func (coupleInvitation *CoupleInvitation) BeforeCreate(tx *gorm.DB) (err error) {
	coupleInvitation.Status = "pending"
	return
}

func FindMyCoupleByUserID(DB *gorm.DB, partnerID uint) (User, error) {
	partner, err := FindUserByID(DB, partnerID)
	if err != nil {
		return User{}, err
	}
	return partner, nil
}
