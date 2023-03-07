package character

import (
	"strconv"
)

type Character struct {
	Name       string  `json:"name"`
	Gender     string  `json:"gender"`
	HeigthStr  string  `json:"height"`
	HeigthFeet float32 `json:"heightFeet"`
	HeigthCM   float32 `json:"heightCM"`
}

func (c *Character) SetHeightInFeet() error {
	hCM, err := strconv.Atoi(c.HeigthStr)
	if err != nil {
		return err
	}
	c.HeigthCM = float32(hCM)
	c.HeigthFeet = c.HeigthCM / 30.48
	return nil
}

type CharacterResponse struct {
	Results []Character `json:"results"`
	Count   int         `json:"count"`
}

type CharacterAscendingOrder bool

var ASCENDING CharacterAscendingOrder = true
var DESCENDING CharacterAscendingOrder = false
