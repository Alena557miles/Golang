package views

import (
	"creator/app/models"
	"encoding/json"
	"fmt"
)

func PrintArt(a *models.Art) error {
	if res, err := json.Marshal(a); err != nil {
		return err
	} else {
		fmt.Println(string(res))
		return nil
	}
}
