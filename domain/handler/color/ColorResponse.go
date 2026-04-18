package color

type ColorResponse struct {
	Id  int    `json:"id"`
	Rgb string `json:"rgb"`
}

type ColorListResponse struct {
	Colors []ColorResponse `json:"colors"`
}
