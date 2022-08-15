package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type Repository interface {
	Users
	News
	Questions
	Firms
}

//User from telegram
type Users interface {
	CreateUser(TelegramID int64, username string) error
	UpdateUser(NewData User) error
	GetUser(TelegramID int64) ([]User, error)
	GetUsers() ([]User, error)
	GetAdmin() ([]User, error)
	DeleteUser(TelegramID int64) error
	DeleteAll() error
}

//News from users
type Newser interface {
	CreateNews(Text string, ContributerID int64) error
	GetAllNews() ([]News, error)
	GetNews(string) ([]News, error)
	GetNotAsReadNews() ([]News, error)
	GetAsReadNews() ([]News, error)
	GetNewsWithDate(time time.Time) ([]News, error)
	GetNewsFromConsumer(ConsumerID int64) ([]News, error)
	UpdateNews(NewNews News) error
	DeleteNews(NewsID uuid.UUID) error
}

//Question from Telgram bot
type Questions interface {
	CreateQuestion(Text string, ContributerID int64, MesasgeID int64) error
	GetAllQuestions() ([]Question, error)
	GetNotAnswerQuestion() ([]Question, error)
	GetAsAnswerQuestion() ([]Question, error)
	GetQuestionFromConsumer(ConsumerID int64) ([]Question, error)
	UpdateQuestion(NewQuestion Question) error
	DeleteQuestion(QuestionID uuid.UUID) error
	MarkAsAnswer(QuestionID uuid.UUID) error
}

type Firms interface {
	CreateFirm(FirmName string) error
	UpdateFirm(NewFirm Firm) error
	GetFirm(Name string) ([]Firm, error)
	DeleteFirm(ID string) error
}

type Models interface {
	CreateModel(FirmID string, ModelName string) error
	UpdateModel(NewModel Model) error
	GetModel(Name string) ([]Model, error)
	DeleteModel(ID string) error
}

type Manuals interface {
	CreateManual(FirmID string, ModelID string, FileUniqID string, Version string) error
	UpdateManual(NewManual Manual) error
	GetManual(Name string) ([]Manual, error)
	DeleteModel(ID string) error
}
