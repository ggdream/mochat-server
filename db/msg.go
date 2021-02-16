package db

import (
	"github.com/ggdream/mochat-server/model/msg"
)


func AddMsg(ms *msg.Message) error {
	_, err := Get().Exec("insert into msg(me, you, time) values (?,?,?)", ms.Me, ms.You, ms.Time)
	if err != nil {
		return err
	}
	return nil
}

func SearchMsg(ms *msg.DescMsg) ([]*msg.Message, error) {
	rows, err := Get().Query("select (me, you, time) from msg where (me=? and you=?) or (me=? and you=?) order by time,id desc limit ?,?", ms.Me, ms.You, ms.You, ms.Me, ms.Start, ms.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make([]*msg.Message, 0)
	for rows.Next() {
		msger := msg.NewMessage()
		if err := rows.Scan(0, msger.Me, msger.You, msger.Time); err != nil {
			return nil, err
		}
		data = append(data, msger)
	}

	return data, nil
}

