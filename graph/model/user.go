package model

type User struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	Password        string          `json:"password"`
	ProfileImageURL string          `json:"profileImageURL"`
	Description     *string         `json:"description,omitempty"`
	HaveAnime       []*AnimeRanking `json:"haveAnime,omitempty"`
}

func (User) IsNode()            {}
func (this User) GetID() string { return this.ID }
