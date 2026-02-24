package usecase

import (
	"fmt"
	"os"
)

func (u *UseCase) CreateFile(bid string) error {
	r, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer r.Close()
	_, err = r.WriteString(fmt.Sprintf("Dólar: %s\n", bid))

	if err != nil {
		return err
	}

	return nil
}
