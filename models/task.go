package models

type user struct {
	Username string `json:"username"`
}

type Task struct {
    ID     int      `json:"id"`
    Name   string   `json:"name"`
    User   user     `json:"user"`
    Price  int      `json:"priceLimit"`
    DaysLeft string `json:"timeLeft"`
}
