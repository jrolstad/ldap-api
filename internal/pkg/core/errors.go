package core

import (
	"errors"
	"fmt"
)

func CombineErrors(toMap []error) error {
	combinedMessage := ""
	for _, item := range toMap {
		combinedMessage = fmt.Sprintf("%v%v;", combinedMessage, item.Error())
	}

	return errors.New(combinedMessage)
}
