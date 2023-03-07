package character

import (
	"encoding/json"
	"errors"
	"sort"

	"github.com/emekarr/coding-test-busha/logger"
	"github.com/emekarr/coding-test-busha/network"
	"go.uber.org/zap"
)

var CharacterService = characterService{network: network.NetworkController{BaseUrl: "https://swapi.dev/api/people"}}

type characterService struct {
	network network.NetworkController
}

// no advanced filter provided by swapi
func (ms *characterService) FetchAllMovieCharacters(page string) (*[]Character, error) {
	response := ms.network.Get("/?page="+page, nil, nil)
	if response.Error != nil {
		logger.Error(errors.New("error fetching characters from character service"), zap.Error(response.Error))
		return nil, response.Error
	}
	var parsedResponse CharacterResponse
	json.Unmarshal(*response.Body, &parsedResponse)
	return &parsedResponse.Results, nil
}

func (ms *characterService) SortMovieCharacters(payload *[]Character, field string, order CharacterAscendingOrder) error {
	fields := []string{"name", "gender", "height"}
	contains := false
	for _, f := range fields {
		if f == field {
			contains = true
			break
		}
	}
	if !contains {
		return errors.New("invalid field added for sorting")
	}
	sort.Slice(*payload, func(i, j int) bool {
		if field == "name" {
			if order == ASCENDING {
				return (*payload)[i].Name < (*payload)[j].Name
			} else {
				return (*payload)[i].Name > (*payload)[j].Name
			}
		} else if field == "gender" {
			if order == ASCENDING {
				return (*payload)[i].Gender < (*payload)[j].Gender
			} else {
				return (*payload)[i].Gender > (*payload)[j].Gender
			}
		} else {
			if order == ASCENDING {
				return (*payload)[i].HeigthCM < (*payload)[j].HeigthCM
			} else {
				return (*payload)[i].HeigthCM > (*payload)[j].HeigthCM
			}
		}
	})
	return nil
}

func (ms *characterService) FilterMovieCharacters(payload *[]Character, filter string) *[]Character {
	if filter != "male" && filter != "female" && filter != "n/a" {
		return payload
	}
	parsedCharacters := []Character{}
	for _, c := range *payload {
		if c.Gender == filter {
			parsedCharacters = append(parsedCharacters, c)
		}
	}
	return &parsedCharacters
}

func (ms *characterService) SetCharacterHeightValueFeet(payload *[]Character) *[]Character {
	parsedCharacters := []Character{}
	for _, c := range *payload {
		c.SetHeightInFeet()
		parsedCharacters = append(parsedCharacters, c)
	}
	return &parsedCharacters
}

func (ms *characterService) ClculateTotalCharacterHeightCM(payload *[]Character) float32 {
	var height float32
	for _, c := range *payload {
		c.SetHeightInFeet()
		height += c.HeigthCM
	}
	return height
}
