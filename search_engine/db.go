package searchengine

import (
	"encoding/json"
	"os"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getData() ([]User, error) {
	file, err := os.Open("search_engine/users.json")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var users []User

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&users); err != nil {
		return nil, err
	}

	return users, nil
}
