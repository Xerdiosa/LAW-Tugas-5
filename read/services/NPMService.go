package services

import (
	"errors"
	"io/ioutil"
	"os"
	"time"
)

type INPMService interface {
	GetNPM(npm string) (name string, err error)
}

type NPMService struct {
}

func (n NPMService) GetNPM(npm string) (name string, err error) {
	if _, err := os.Stat("../record/" + npm); err == nil {
		content, _ := ioutil.ReadFile("../record/" + npm)
		name := string(content)
		time.Sleep(15 * time.Millisecond) // mensimulasikan delay pada database
		return name, err
	} else if os.IsNotExist(err) {
		err = errors.New("not found")
		return name, err
	} else {
		err = errors.New("other error")
		return name, err
	}
}
