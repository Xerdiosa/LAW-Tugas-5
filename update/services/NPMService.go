package services

import "os"

type INPMService interface {
	AddNPM(npm string, name string) (err error)
}

type NPMService struct {
}

func (n NPMService) AddNPM(npm string, name string) (err error) {
	file, err := os.Create("../record/" + npm)

	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString(name)
	return err
}
