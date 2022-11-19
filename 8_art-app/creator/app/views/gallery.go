package views

import (
	"creator/app/models"
	"encoding/json"
	"fmt"
)

func PrintGallery(g *models.Gallery) error {
	if res, err := json.Marshal(g); err != nil {
		return err
	} else {
		fmt.Println(string(res))
		return nil
	}
}
