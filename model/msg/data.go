package msg


type Status struct {
	Code		int32		`json:"code"`
	Message		string		`json:"message"`
}

type Message struct {
	ID				string	`json:"id"`
	From			string	`json:"from"`
	To				string	`json:"to"`
}

func NewMessage() *Message {
	return new(Message)
}

type DescMsg struct {
	From			string	`json:"from" form:"from"`
	To				string	`json:"to" form:"to" binding:"required"`
	Start			uint32	`json:"start" form:"start" binding:"required"`
	Offset			uint32	`json:"offset" form:"offset" binding:"required"`
}

func NewDescMsg() *DescMsg {
	return new(DescMsg)
}

type Result struct {
	Code			int32		`json:"code"`
	Message			string		`json:"message"`
	Records			[]*Message	`json:"records"`
}
