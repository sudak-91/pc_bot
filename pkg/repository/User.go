package repository

type User struct {
	UserID     string `bson:"userid"`
	TelegramID int64  `bson:"telegram_id"`
	Username   string `bson:"Username"`
}
