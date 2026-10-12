package storage

import (
	"encoding/json"
	"os"
)

func SaveToFile(target any) error {
	bytes, err := json.MarshalIndent(target, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("expenseList.json", bytes, 0644)
	return err
}

func ReadFromFile(target any) error {
	bytes, err := os.ReadFile("expenseList.json")
	if err != nil {
		if os.IsNotExist(err) {
			err = SaveToFile(target)
			return err
		}
		return err
	}

	err = json.Unmarshal(bytes, target)
	return err
}
