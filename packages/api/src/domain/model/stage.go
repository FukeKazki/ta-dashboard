package model

type Stage struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	TimeRecord TimeRecord `json:"timeRecord"`
}

type TimeRecord struct {
	FirstTime  string `json:"firstTime"`
	SecondTime string `json:"secondTime"`
	ThirdTime  string `json:"thirdTime"`
	TotalTime  string `json:"totalTime"`
}
