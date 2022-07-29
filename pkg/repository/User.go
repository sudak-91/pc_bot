package repository

type User struct {
	TelegramID int64  `bson:"_id,omitempty"`
	Username   string `bson:"Username,omitempty"`
	Role       byte   `bson:"role,omitempty"`
}
