package battle

import (
	"errors"
	"fmt"
)

type Battle struct {
	//Each players field configuration
	Player_one_field_config []int `json:"player_one_field_config,omitempty"`
	Player_two_field_config []int `json:"player_two_field_config,omitempty"`
	//Each players name
	Player_one_name string `json:"player_one_name,omitempty"`
	Player_two_name string `json:"player_two_name,omitempty"`
	//The amount of fields to use
	Field_total int `json:"total_fields,omitempty"`
	//The score fore each field
	Field_score []int `json:"file_score,omitempty"`
	//The score of both players
	Player_one_score int `json:"player_one_score,omitempty"`
	Player_two_score int `json:"player_two_score,omitempty"`
	//The max amount of units per a field
	Unit_limit int `json:"unit_limit,omitempty"`
	//The total unit count per player
	Unit_total int `json:"unit_total,omitempty"`
	//The winner of a battle
	Winner string `json:"winner,omitempty"`
}

//calcuate_winner determines the winner of a battle
//It sets the player score for each player and sets the winner flag, which could be
//either players name, or "draw".
//An error could be returned
func (b *Battle) calcuate_winner() error {
	player_one_score := 0
	player_two_score := 0
	for i := 0; i < b.Field_total; i++ {
		if b.Player_one_field_config[i] < b.Player_two_field_config[i] {
			player_one_score += 0
			player_two_score += b.Field_score[i]
		} else if b.Player_one_field_config[i] > b.Player_two_field_config[i] {
			player_two_score += 0

			player_one_score += b.Field_score[i]

		} else if b.Player_one_field_config[i] == b.Player_two_field_config[i] {
			player_two_score += 0
			player_one_score += 0

		} else if b.Player_one_field_config[i] == b.Player_two_field_config[i] {
			player_two_score += 0
			player_one_score += 0

		} else {
			return errors.New("Could not calculate scores, bad configuration")
		}
	}
	if player_one_score < player_two_score {
		b.Winner = b.Player_two_name
	} else if player_one_score > player_two_score {
		b.Winner = b.Player_one_name
	} else if player_one_score == player_two_score {
		b.Winner = "draw"
	}
	b.Player_one_score = player_one_score
	b.Player_two_score = player_two_score
	return nil
}

//validate checks to see if the battle object is correct before starting a game
//An error is returned
func (b *Battle) validate() error {
	fmt.Print("Hey\n")
	if len(b.Player_two_field_config) < 2 || len(b.Player_one_field_config) < 2 {
		return errors.New("Player config size must be greater than 2")
	}
	if b.Player_one_name == b.Player_two_name || b.Player_one_name == "" || b.Player_two_name == "" {
		return errors.New("Player names cannot be the same, or empty")
	}
	if b.Field_total != len(b.Player_two_field_config) || b.Field_total != len(b.Player_one_field_config) || b.Field_total != len(b.Field_score) {
		return errors.New("Player field configs don't match field_total or field score hasn't got the correct number of elements")
	}
	player_one_sum := 0
	player_two_sum := 0
	for index, element := range b.Player_one_field_config {
		player_one_sum += element
		player_two_sum += b.Player_two_field_config[index]
		if b.Unit_limit > 0 {
			if element > b.Unit_limit {
				return errors.New("Player " + b.Player_one_name + " unit placement on field is greater than unit_limit")
			}
			if b.Player_two_field_config[index] > b.Unit_limit {
				return errors.New("Player " + b.Player_two_name + " unit placement on field is greater than unit_limit")
			}
		}
	}
	if player_two_sum > b.Unit_total || player_one_sum > b.Unit_total {
		return errors.New("Player field config sum does not equal unit_total")
	}

	return nil
}
