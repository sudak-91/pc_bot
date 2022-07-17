package service

import (
	command "github.com/sudak-91/telegrambotgo/Command"
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

func (t *TelegramUpdater) CallbackQueryService(Query *types.TelegramCallbackQuery) ([]byte, error) {
	return t.Execute(Query.Data, Query)
}

func (t *TelegramUpdater) ChannelPostService(Post *types.TelegramMessage) ([]byte, error) {
	return nil, nil
}

func (t *TelegramUpdater) ChatJoinRequsetService(JoinRequest *types.TelegramChatJoinRequest) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) ChatMemberService(MemberService *types.TelegramChatMemberUpdated) ([]byte, error) {
	return nil, nil

}
func (t *TelegramUpdater) ChosenInlineResultService(*types.TelegramChosenInlineResult) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) EditedChannelPostService(*types.TelegramMessage) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) EditedMessageService(*types.TelegramMessage) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) InlineQueryService(*types.TelegramInlineQuery) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) MessageService(Message *types.TelegramMessage) ([]byte, error) {
	//TODO: Rewrite Message Service
	return t.Execute(Message.Text, Message)
}
func (t *TelegramUpdater) MyChatMemberService(MyChatMember *types.TelegramChatMemberUpdated) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) PollService(Poll *types.TelegramPoll) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) PollAnswerService(PollAnswer *types.TelegramPollAnwer) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) PreCheckoutPollService(CheckoutPoll *types.TelegramPreCheckoutQuery) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) ShippingService(Shipping *types.TelegramShippingQuery) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) ChatUserUpdateService(UserUpdate *types.TelegramUpdate) ([]byte, error) {
	return nil, nil
}
func (t *TelegramUpdater) Default(Update types.TelegramUpdate) ([]byte, error) {
	return t.Default(Update)
}
