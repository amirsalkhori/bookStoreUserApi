package users

import (
	"encoding/json"
)

type PublicUser struct {
	Id        int64  `json:"id"`
	// Name      string `json:"name"`
	// Family    string `json:"family"`
	CreatedAt string `json:"createdAt"`
	Status    bool   `json:"status"`
}

type PrivateUser struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Family    string `json:"family"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	Status    bool   `json:"status"`
}

func(user *User) Marshall(isPublic bool) interface{}{
	if isPublic {
		return PublicUser{
			Id: user.Id,
			CreatedAt: user.CreatedAt,
			Status: user.Status,
		}
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		return err
	}
	var privateUser PrivateUser
	err = json.Unmarshal(userJson, &privateUser)
	if err != nil {
		return err
	}
	
	return privateUser
}

func(users Users) Marshall(isPublic bool) interface{}{
	result := make([]interface{}, len(users))
	for index, user := range users{
		result[index] = user.Marshall(isPublic)
	}

	return result
}