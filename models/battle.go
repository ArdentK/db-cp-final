package models

type Battle struct {
	ID            int `json:"id"`
	IDWinner      int `json:"id_winner"`
	IDFighter2    int `json:"id_fighter"`
	IDCompetition int `json:"id_competition"`
	ScoreWinner   int `json:"score_winner"`
	ScoreFighter2 int `json:"score_fighter2"`
}
