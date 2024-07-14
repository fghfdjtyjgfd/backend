package model


type Book struct {
	Model
	Title string `json:"title"`
	Author string `json:"author"`
	Genre string `json:"genre"`
	Year_of_publication string `json:"year_of_publication"`

}