package dto

type GetJWTInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTOutput struct {
	Token string `json:"token"`
}

type MatchDTO struct {
	HomeTeam struct {
		Name string `json:"name"`
	} `json:"homeTeam"`
	AwayTeam struct {
		Name string `json:"name"`
	} `json:"awayTeam"`
	Score struct {
		FullTime struct {
			HomeTeam *int `json:"homeTeam"`
			AwayTeam *int `json:"awayTeam"`
		} `json:"fullTime"`
	} `json:"score"`
}
