package repository

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	Users
	News
	Questions
	Firms
}

// User from telegram
type Users interface {
	CreateUser(TelegramID int64, username string) error
	UpdateUser(NewData User) error
	GetUser(TelegramID int64) ([]User, error)
	GetUsers() ([]User, error)
	GetAdmin() ([]User, error)
	DeleteUser(TelegramID int64) error
	DeleteAll() error
}

// News from users
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

// Question from Telgram bot
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
	CreateFirm(FirmName string) (primitive.ObjectID, error)
	UpdateFirm(NewFirm Firm) error
	GetFirm(Name string) ([]Firm, error)
	GetFirmById(ID string) (Firm, error)
	GetFirms() ([]Firm, error)
	GetApprovedFirms() ([]Firm, error)
	GetApprovedFirmsWithOffsetAndLimit(offset int64, limit int, approved bool) ([]Firm, error)
	GetAllFirmsWithOffsetAndLimit(offset int64, limit int) ([]Firm, error)
	DeleteFirm(ID string) error
}

type Manuals interface {
	CreateManual(NewManual Manual) error
	UpdateManual(NewManual Manual) error
	GetManuals() ([]Manual, error)
	UpdateEmbeddedFirm(NewFirm Firm) error
	DeleteManuals(ID string) error
	DeleteManualsByFirm(ID string) error
	GetManualByID(ID string) (Manual, error)
	GetManualsByFirmID(FirmID string) ([]Manual, error)
	GetApprovedManuals(approved bool) ([]Manual, error)
	GetApprovedManualsWithOffsetAndLimit(FirmID string, offset int64, limit int, approved bool) ([]Manual, error)
}
