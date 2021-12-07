package bot

import (
	tele "gopkg.in/tucnak/telebot.v3"
	"gopkg.in/tucnak/telebot.v3/layout"
)

type Bot struct {
	*tele.Bot
	*layout.Layout
}

func IsChannelMessage(b *Bot) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {

			if c.Message().FromGroup() && c.Message().SenderChat.Type == "channel" {
				msg := c.Message()
				b.Delete(msg)
			}
			return next(c)
		}
	}
}
func New(path string) (*Bot, error) {
	lt, err := layout.New(path)
	if err != nil {
		return nil, err
	}

	b, err := tele.NewBot(lt.Settings())
	if err != nil {
		return nil, err
	}

	if cmds := lt.Commands(); cmds != nil {
		if err := b.SetCommands(cmds); err != nil {
			return nil, err
		}
	}

	return &Bot{
		Bot:    b,
		Layout: lt,
	}, nil
}

func (b *Bot) Start() {
    b.Use(IsChannelMessage(b))
	b.Handle("/start", func(context tele.Context) error {
		return context.Send("Просто добавь меня в группу и дай права.")
	})
	b.Handle(tele.OnText, func(context tele.Context) error {
		return nil
	})
	b.Handle(tele.OnVideo, func(context tele.Context) error {
		return nil
	})
	b.Handle(tele.OnAudio, func(context tele.Context) error {
		return nil
	})
	b.Handle(tele.OnSticker, func(context tele.Context) error {
		return nil
	})
	b.Handle(tele.OnAnimation, func(context tele.Context) error {
		return nil
	})

	b.Bot.Start()
}