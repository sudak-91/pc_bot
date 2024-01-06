package repository

import (
	"context"

	"github.com/google/uuid"
	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Questionmongo struct {
	col *mongo.Collection
}

func NewQuestionmongo(db *mongo.Database) *Questionmongo {
	var q Questionmongo
	q.col = db.Collection("Question")
	return &q
}

func (q *Questionmongo) CreateQuestion(Text string, ContributerID int64, MessageID int64) error {
	var ques pubrep.Question
	var err error
	ques.QuestionID = uuid.New()
	if err != nil {
		return err
	}
	ques.Text = Text
	ques.ContributerID = ContributerID
	ques.MessageID = MessageID
	data, err := bson.Marshal(ques)
	if err != nil {
		return err
	}
	_, err = q.col.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	return nil
}
func (q *Questionmongo) GetAllQuestions() ([]pubrep.Question, error) {
	rslt, err := q.col.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var qu []pubrep.Question
	err = rslt.All(context.TODO(), &qu)
	if err != nil {
		return nil, err
	}
	return qu, nil

}

func (q *Questionmongo) GetNotAnswerQuestion() ([]pubrep.Question, error) {
	rslt, err := q.col.Find(context.TODO(), bson.D{{"asanswer", false}})
	if err != nil {
		return nil, err
	}
	var qu []pubrep.Question
	err = rslt.All(context.TODO(), &qu)
	if err != nil {
		return nil, err
	}
	return qu, nil
}
func (q *Questionmongo) GetAsAnswerQuestion() ([]pubrep.Question, error) {
	rslt, err := q.col.Find(context.TODO(), bson.D{{"asanswer", true}})
	if err != nil {
		return nil, err
	}
	var qu []pubrep.Question
	err = rslt.All(context.TODO(), &qu)
	if err != nil {
		return nil, err
	}
	return qu, nil
}
func (q *Questionmongo) GetQuestionFromConsumer(ConsumerID int64) ([]pubrep.Question, error) {
	rslt, err := q.col.Find(context.TODO(), bson.D{{"contributerid", ConsumerID}})
	if err != nil {
		return nil, err
	}
	var qu []pubrep.Question
	err = rslt.All(context.TODO(), &qu)
	if err != nil {
		return nil, err
	}
	return qu, nil
}
func (q *Questionmongo) UpdateQuestion(NewQuestion pubrep.Question) error {
	return nil
}
func (q *Questionmongo) DeleteQuestion(QuestionID uuid.UUID) error {
	return nil
}
func (q *Questionmongo) MarkAsAnswer(QuestionID uuid.UUID) error {
	filter := bson.D{{"_id", QuestionID}}
	upd := bson.D{{"$set", bson.D{{"asanswer", true}}}}
	_, err := q.col.UpdateOne(context.TODO(), filter, upd)
	if err != nil {
		return err
	}
	return nil
}
