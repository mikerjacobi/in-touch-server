package models

import "fmt"

//Friend base friend object
type Friend struct {
	ID   string `json:"friendID"`
	Name string `json:"name"`
}

//LoadFriends fetch friends from the db given a query
func LoadFriends(query map[string]string) ([]*Friend, error) {
	fmt.Printf("here")
	db := map[string]string{
		"a": "anne",
		"b": "bob",
		"c": "charlie",
	}

	name := db[query["friend_id"]]
	if name == "" {
		return nil, ErrNotFound
	}

	f := []*Friend{
		{query["friend_id"], name},
	}
	return f, nil
}
