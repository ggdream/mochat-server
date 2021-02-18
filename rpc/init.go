package rpc


var rpc *Rpc


func Get() *Rpc {
	return rpc
}

func Init() (err error) {
	rpc, err = New()
	return err
}
