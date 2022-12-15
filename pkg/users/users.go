package users

import (
	"encoding/json"
	"fmt"
	"ta/db"
)

type User struct {
	ID   int   `json:"id"`
	Data *Data `json:"data"`
}

type Data struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Interests  string `json:"interests"`
}

func NewUser(id int, first, last, interests string) *User {
	return &User{
		ID: id,
		Data: &Data{
			First_name: first,
			Last_name:  last,
			Interests:  interests,
		},
	}
}

func (s *User) Create() error {
	data, err := serialize(s.Data)
	if err != nil {
		return err
	}

	if _, err = db.DB.NamedExec(`INSERT INTO users (id, data) VALUES (:id, :data)`,
		map[string]interface{}{
			"id":   s.ID,
			"data": data,
		}); err != nil {
		return fmt.Errorf("error creating users: %w", err)
	}

	return err
}

func (s *User) Read() error {
	rows, err := db.DB.NamedQuery(`SELECT * FROM users WHERE id=:fn`, map[string]interface{}{"fn": s.ID})
	if err != nil {
		return fmt.Errorf("error geting users: %w", err)
	}

	var id byte
	var data []byte

	for rows.Next() {
		err := rows.Scan(&id, &data)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	if err = json.Unmarshal(data, &s.Data); err != nil {
		return fmt.Errorf("error unmarshaling: %w", err)
	}

	return nil
}

func (s *User) Update() error {
	data, err := serialize(s.Data)
	if err != nil {
		return err
	}

	if _, err = db.DB.NamedExec(`UPDATE users SET data=:data WHERE id=:id`,
		map[string]interface{}{
			"data": data,
			"id":   s.ID,
		}); err != nil {
		return fmt.Errorf("error updating users: %w", err)
	}
	return nil
}

func (s *User) Delete() error {
	if _, err := db.DB.NamedExec(`DELETE FROM users WHERE id =:id`,
		map[string]interface{}{
			"id": s.ID,
		}); err != nil {
		return fmt.Errorf("error deleting users: %w", err)
	}
	return nil
}

func serialize(data *Data) ([]byte, error) {
	d, err := json.Marshal(data)

	return d, err
}
