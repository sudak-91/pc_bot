package service

import (
	"log"
	"os"

	command "github.com/sudak-91/telegrambotgo/Command"
	methods "github.com/sudak-91/telegrambotgo/TelegramAPI/Methods"
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

type TelegramUpdater struct {
	command.ICommandService
}

func NewTelegramUpdater() *TelegramUpdater {
	return &TelegramUpdater{
		command.NewTelegramCommandService(),
	}

}

func (t *TelegramUpdater) CallbackQueryService(Query types.TelegramCallbackQuery) ([]byte, error) {
	var answerCallback methods.AnswerCallBackQuery
	answerCallback.CallbackQueryId = Query.ID
	data, err := t.Execute(Query.Data, Query)
	if err != nil {
		answerCallback.Text = "Внутренняя ошибка"
		answerCallback.ShowAlert = true
		err := methods.AnswerCallBackQueryMethod(os.Getenv("BOT_KEY"), answerCallback)
		if err != nil {
			log.Println("_________CakkbackSendError__________")
		}
		return data, err
	}
	answerCallback.Text = "OK"
	err = methods.AnswerCallBackQueryMethod(os.Getenv("BOT_KEY"), answerCallback)
	if err != nil {
		log.Println("_________CakkbackSendError__________")
	}
	return data, nil
}

func (t *TelegramUpdater) ChannelPostService(Post types.TelegramMessage) ([]byte, error) {
	return nil, nil
}

func (t *TelegramUpdater) ChatJoinRequsetService(JoinRequest types.TelegramChatJoinRequest) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) ChatMemberService(MemberService types.TelegramChatMemberUpdated) ([]byte, error) {
	return nil, nil

}
func (t *TelegramUpdater) ChosenInlineResultService(TelegramChosenInlien types.TelegramChosenInlineResult) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) EditedChannelPostService(Message types.TelegramMessage) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) EditedMessageService(Message types.TelegramMessage) ([]byte, error) {
	log.Println("Edited log service")
	log.Println(Message)
	return t.Execute("/default", Message)
}
func (t *TelegramUpdater) InlineQueryService(InlineQuery types.TelegramInlineQuery) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) MessageService(Message types.TelegramMessage) ([]byte, error) {
	return t.messageService(Message)

}
func (t *TelegramUpdater) MyChatMemberService(MyChatMember types.TelegramChatMemberUpdated) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) PollService(Poll types.TelegramPoll) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) PollAnswerService(PollAnswer types.TelegramPollAnwer) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) PreCheckoutPollService(CheckoutPoll types.TelegramPreCheckoutQuery) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) ShippingService(Shipping types.TelegramShippingQuery) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) ChatUserUpdateService(UserUpdate types.TelegramUpdate) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) Default(Update types.TelegramUpdate) ([]byte, error) {
	return t.Default(Update)
}
