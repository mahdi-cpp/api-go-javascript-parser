package model

type MovieArtists struct {
	Id   int `json:"id"`
	Cast []struct {
		Adult              bool    `json:"adult"`
		Gender             int     `json:"gender"`
		Id                 int     `json:"id"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		Popularity         float64 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
		CastId             int     `json:"cast_id"`
		Character          string  `json:"character"`
		CreditId           string  `json:"credit_id"`
		Order              int     `json:"order"`
	} `json:"cast"`
	Crew []struct {
		Adult              bool    `json:"adult"`
		Gender             int     `json:"gender"`
		Id                 int     `json:"id"`
		KnownForDepartment string  `json:"known_for_department"`
		Name               string  `json:"name"`
		OriginalName       string  `json:"original_name"`
		Popularity         float64 `json:"popularity"`
		ProfilePath        string  `json:"profile_path"`
		CreditId           string  `json:"credit_id"`
		Department         string  `json:"department"`
		Job                string  `json:"job"`
	} `json:"crew"`
}
