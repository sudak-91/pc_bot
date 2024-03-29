package service

import (
	"log"
	"os"

	command "github.com/sudak-91/telegrambotgo/Command"
	methods "github.com/sudak-91/telegrambotgo/telegram_api/methods"
	types "github.com/sudak-91/telegrambotgo/telegram_api/types"
)

type TelegramUpdater struct {
	command.ICommandService
}

func NewTelegramUpdater() *TelegramUpdater {
	return &TelegramUpdater{
		command.NewTelegramCommandService(),
	}

}

func (t *TelegramUpdater) CallbackQueryService(Query types.CallbackQuery) ([]byte, error) {
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

func (t *TelegramUpdater) ChannelPostService(Post types.Message) ([]byte, error) {
	return nil, nil
}

func (t *TelegramUpdater) ChatJoinRequsetService(JoinRequest types.ChatJoinRequest) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) ChatMemberService(MemberService types.ChatMemberUpdated) ([]byte, error) {
	return nil, nil

}
func (t *TelegramUpdater) ChosenInlineResultService(TelegramChosenInlien types.ChosenInlineResult) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) EditedChannelPostService(Message types.Message) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) EditedMessageService(Message types.Message) ([]byte, error) {
	log.Println("Edited log service")
	log.Println(Message)
	return t.Execute("/default", Message)
}
func (t *TelegramUpdater) InlineQueryService(InlineQuery types.InlineQuery) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) MessageService(Message types.Message) ([]byte, error) {
	return t.messageService(Message)

}
func (t *TelegramUpdater) MyChatMemberService(MyChatMember types.ChatMemberUpdated) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) PollService(Poll types.Poll) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) PollAnswerService(PollAnswer types.PollAnwer) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) PreCheckoutPollService(CheckoutPoll types.PreCheckoutQuery) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) ShippingService(Shipping types.ShippingQuery) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) ChatUserUpdateService(UserUpdate types.Update) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) Default(Update types.Update) ([]byte, error) {
	return t.Default(Update)
}
