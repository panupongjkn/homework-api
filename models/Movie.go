package models

type Movie struct {
	ID            int    `json:"id"`
	Cover         string `json:"cover"`
	Name          string `json:"name"`
	DateRangeShow string `json:"date_range_show"`
}

func NewMovie(id int, cover, name, date_range_show string) Movie {
	movie := Movie{}
	movie.ID = id
	movie.Cover = cover
	movie.Name = name
	movie.DateRangeShow = date_range_show
	return movie
}
