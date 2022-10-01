package server

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"sync"

	pubrep "github.com/sudak-91/pc_bot/pkg/repository"
	update "github.com/sudak-91/telegrambotgo/Service"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

var Util *Utl

type FMSStage int

const (
	Addnews FMSStage = iota
	AddQuestion
	SendAnswerTo
	AddManualInfo
	AddManual
	AddManualDocument
	EditFirm
	ConfirmFirm
	EditManual
)

type Server struct {
	port                 string
	key                  string
	updater              update.Updater
	updateNewsSignal     chan bool
	updateQuestionSignal chan bool
	once                 sync.Once
}

type SendAnswer struct {
	QuestionID    uuid.UUID
	MessageID     int32
	ContributerID int64
}

func NewServer(Port string, key string, Upd update.Updater) *Server {
	return &Server{
		port:    Port,
		key:     key,
		updater: Upd,
	}
}

type Utl struct {
	StageMutex      *sync.RWMutex
	Stage           map[int64]FMSStage
	AnswerCtx       map[int64]SendAnswer
	Manual          map[int64]pubrep.Manual
	EditFirm        map[int64]pubrep.Firm
	EditManualMutex *sync.RWMutex
	EditManual      map[int64]pubrep.Manual
	AdminID         int64
}

func (s *Server) Run(AdminID int64) {
	s.once.Do(func() {
		Util = &Utl{
			StageMutex:      &sync.RWMutex{},
			Stage:           make(map[int64]FMSStage),
			AdminID:         AdminID,
			AnswerCtx:       make(map[int64]SendAnswer),
			Manual:          make(map[int64]pubrep.Manual),
			EditFirm:        make(map[int64]pubrep.Firm),
			EditManualMutex: &sync.RWMutex{},
			EditManual:      make(map[int64]pubrep.Manual)}

	})
	mux := http.NewServeMux()
	mux.HandleFunc("/"+s.key, s.Handl)
	log.Println("Server start")
	log.Fatal(http.ListenAndServe(":"+s.port, mux))
}

func (s *Server) Handl(w http.ResponseWriter, r *http.Request) {
	var b []byte
	log.Println("New request")
	buffer := bytes.NewBuffer(b)
	buffer.ReadFrom(r.Body)
	fmt.Printf("%+v", buffer.String())

	k, err := s.updater.Update(buffer.Bytes())
	if err != nil {
		log.Printf("!!!!_____ HAS ERROR: %s", err.Error())
	}

	w.Header().Set(http.CanonicalHeaderKey("Content-Type"), "application/json")
	_, err = w.Write(k)
	if err != nil {
		panic(err.Error())
	}
}
