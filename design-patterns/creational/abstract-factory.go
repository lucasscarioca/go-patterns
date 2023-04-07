package creational

import "fmt"

/*
# Abstract Factory Pattern

## Concept
Allows the client to create objects (products) without specifying their concrete
classes. Avaiable products that can be built with an Abstract Factory need to
share a common interface.

## Example explanation
In this pattern we create an abstract factory interface with a GetFactory Method
to instantiate concrete factories of related objects/entities that can have
their own internal logic.

The following example illustrates a messaging application where the client can
send emails and push notifications based on the given channel.
*/

// ##### Abstract factory interface #####
type IFactory interface {
	makeMail(to string, subject string, body string) IMail
	makePush(to string, text string) IPush
}

func GetMessageFactory(channel string) (IFactory, error) {
	if channel == "whatsapp" {
		return &WhatsApp{}, nil
	}

	if channel == "telegram" {
		return &Telegram{}, nil
	}

	return nil, fmt.Errorf("wrong channel type passed")
}

// ##### Concrete factory #####
type WhatsApp struct{}

func (w *WhatsApp) makeMail(to string, subject string, body string) IMail {
	return &WhatsAppMail{
		Mail: Mail{
			to,
			subject,
			body,
		},
	}
}

func (w *WhatsApp) makePush(to string, text string) IPush {
	return &WhatsAppPush{
		Push: Push{
			to,
			text,
		},
	}
}

// ##### Concrete factory #####
type Telegram struct{}

func (t *Telegram) makeMail(to string, subject string, body string) IMail {
	return &TelegramMail{
		Mail: Mail{
			to,
			subject,
			body,
		},
	}
}

func (t *Telegram) makePush(to string, text string) IPush {
	return &TelegramPush{
		Push: Push{
			to,
			text,
		},
	}
}

// ##### Abstract product #####
type IMail interface {
	setTo(to string)
	setSubject(subject string)
	setBody(text string)
	getTo() string
	getSubject() string
	getBody() string
}

type Mail struct {
	to      string
	subject string
	body    string
}

func (m *Mail) setTo(to string) {
	m.to = to
}

func (m *Mail) setSubject(subject string) {
	m.subject = subject
}

func (m *Mail) setBody(body string) {
	m.body = body
}

func (m *Mail) getTo() string {
	return m.to
}

func (m *Mail) getSubject() string {
	return m.subject
}

func (m *Mail) getBody() string {
	return m.body
}

// ##### Concrete product #####
type WhatsAppMail struct {
	Mail
}

// ##### Concrete product #####
type TelegramMail struct {
	Mail
}

// ##### Abstract product #####
type IPush interface {
	setTo(to string)
	getTo() string
	setText(text string)
	getText() string
}

type Push struct {
	to   string
	text string
}

func (p *Push) setTo(to string) {
	p.to = to
}

func (p *Push) setText(text string) {
	p.text = text
}

func (p *Push) getTo() string {
	return p.to
}

func (p *Push) getText() string {
	return p.text
}

// ##### Concrete product #####
type WhatsAppPush struct {
	Push
}

// ##### Concrete product #####
type TelegramPush struct {
	Push
}

// ##### Client Code #####
func RunAbstractFactoryExample() {
	whatsappFactory, _ := GetMessageFactory("whatsapp")
	telegramFactory, _ := GetMessageFactory("telegram")

	whatsappMail := whatsappFactory.makeMail("test@test.com", "Whatsapp Email subject", "Whatsapp Email body")
	whatsappPush := whatsappFactory.makePush("user1", "Whatsapp Push Notification Body")

	telegramMail := telegramFactory.makeMail("test@test.com", "Telegram Email subject", "Telegram email body")
	telegramPush := telegramFactory.makePush("user1", "Telegram Push Notification Body")

	printMailDetails(whatsappMail)
	printPushDetails(whatsappPush)

	printMailDetails(telegramMail)
	printPushDetails(telegramPush)
}

func printMailDetails(m IMail) {
	fmt.Printf("To: %s", m.getTo())
	fmt.Println()
	fmt.Printf("Subject: %s", m.getSubject())
	fmt.Println()
	fmt.Printf("Body: %s", m.getBody())
	fmt.Println()
}

func printPushDetails(p IPush) {
	fmt.Printf("To: %s", p.getTo())
	fmt.Println()
	fmt.Printf("Text: %s", p.getText())
	fmt.Println()
}
