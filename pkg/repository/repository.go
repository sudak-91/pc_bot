package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type Repository interface {
	Users
	News
	Questions
}

//User from telegram
type Users interface {
	CreateUser(TelegramID int64, username string) error
	UpdateUser(NewData User) error
	GetUser(TelegramID int64) ([]User, error)
	GetUsers() ([]User, error)
	DeleteUser(TelegramID int64) error
	DeleteAll() error
}

//News from users
type Newser interface {
	CreateNews(Text string, ContributerID int64) error
	GetAllNews() ([]News, error)
	GetNotAsReadNews() ([]News, error)
	GetAsReadNews() ([]News, error)
	GetNewsWithDate(time time.Time) ([]News, error)
	GetNewsFromConsumer(ConsumerID int64) ([]News, error)
	UpdateNews(NewNews News) error
	DeleteNews(NewsID uuid.UUID) error
}

//Question from Telgram bot
type Questions interface {
	CreateQuestion(Text string, ContributerID int64) error
	GetAllQuestions() ([]Question, error)
	GetNotAnswerQuestion() ([]Question, error)
	GetAsAnswerQuestion() ([]Question, error)
	GetQuestionFromConsumer(ConsumerID int64) ([]Question, error)
	UpdateQuestion(NewQuestion Question) error
	DeleteQuestion(QuestionID string) error
	MarkAsAnswer(QuestionID uuid.UUID) error
}
