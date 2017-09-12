package commands

type Command interface {
	CallName() string
	GetRequestBody() []byte
	SetToken(token string)
}
