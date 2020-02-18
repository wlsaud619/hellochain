package types

import (
	"fmt"
	"strings"

	codec "github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// module의 이름
	ModuleName = "greeter"

	// module의 store를 등록하는 데 사용
	StoreKey = ModuleName
)

var (
	// amino로 인코딩해야 하는 모듈의 type을 포함
	ModuleCdc = codec.New()
)

// `json:~`과 `yaml:~`은json을marshalling 할 때 필드 이름을 지정하는 데 사용
type Greeting struct {
	Sender		sdk.AccAddress	`json:"sender" yaml:"sender"`
	Recipient	sdk.AccAddress	`json:"receiver" yaml:"receiver"`
	Body		string		`json:"body" yaml:"body"`

}

// 주어진 주소에 대한 모든 인사말(greeting)을 저장
type GreetingsList []Greeting

// Greeting을 반환 
func NewGreeting(sender sdk.AccAddress, body string, receiver sdk.AccAddress) Greeting {
	return Greeting {
		Recipient:	receiver,
		Sender:		sender,
		Body:		body,
	}
}

// fmt.Stringer 구현
func (g Greeting) String() string {
	return strings.TrimSpace(
		fmt.Sprintf(`Sender: %s, Recipient: %s, Body: %s`, g.Sender.String(), g.Recipient.String(), g.Body),
	)
}

// 주어진 주소에 대한 인사말을 포함하는Query에 대한 응답을 정의
type QueryResGreetings map[string][]Greeting

func (q  QueryResGreetings) String() string {
	b := ModuleCdc.MustMarshalJSON(q)
	return string(b)
}

// 새 인스턴스 구성
func NewQueryResGreetings() QueryResGreetings {
	return make(map[string][]Greeting)
}
