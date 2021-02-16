package msg


type Message struct {
	Me				string	`json:"me" form:"me"`
	You				string	`json:"you" form:"you"`
	Time			int64	`json:"time" form:"time"`
}

func NewMessage() *Message {
	return new(Message)
}

type DescMsg struct {
	Me				string	`json:"me" form:"me" binding:"required"`
	You				string `json:"you" form:"you" binding:"required"`
	Start			int		`json:"start" form:"start" binding:"required"`
	End				int		`json:"end" form:"end"`
	Offset			int		`json:"offset" form:"offset" binding:"required"`
}

func NewDescMsg() *DescMsg {
	return new(DescMsg)
}
