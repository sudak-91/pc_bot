package command

import "log"

type Default struct {
}

func (d *Default) Handl(data interface{}) ([]byte, error) {
	log.Println("Default")
	return nil, nil
}
