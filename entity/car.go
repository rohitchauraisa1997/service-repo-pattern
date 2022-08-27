package entity

type Car struct {
	CarDetails `json:"Car"`
}

type CarDetails struct {
	ID    int64  `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  string `json:"year"`
}
