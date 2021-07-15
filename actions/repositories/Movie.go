package repositories

import (
	"homework-api/models"
)

func GetMovie() ([]models.Movie, error) {
	movie1 := models.NewMovie(1, "https://www.img02.xyz/assets/movie_poster/2/201f545eebfd4154b7095b14e62f1961.jpg", "Ironman 1", "13ส.ค.2561 - 21ส.ค.2561")
	movie2 := models.NewMovie(2, "https://i.pinimg.com/originals/dd/fd/24/ddfd249a481ece6180df276b61b893a8.jpg", "Ironman 2", "13ส.ค.2561 - 21ส.ค.2561")
	movie3 := models.NewMovie(3, "https://i.pinimg.com/originals/f8/f4/0c/f8f40cd2f8f6819c8ee6b8370dab15d4.png", "Ironman 3", "13ส.ค.2561 - 21ส.ค.2561")
	movie4 := models.NewMovie(4, "https://m.media-amazon.com/images/I/71krfsS5kNL._AC_SY550_.jpg", "Captain America 1", "13ส.ค.2561 - 21ส.ค.2561")
	movie5 := models.NewMovie(5, "https://images-na.ssl-images-amazon.com/images/I/91OxromzoSL._AC_SL1500_.jpg", "Captain America 2", "13ส.ค.2561 - 21ส.ค.2561")
	movie6 := models.NewMovie(6, "https://cf.shopee.co.th/file/df8c560c9086b80e74b9e2f896c75c25", "Captain America 3", "13ส.ค.2561 - 21ส.ค.2561")
	movie7 := models.NewMovie(7, "https://m.media-amazon.com/images/I/71krfsS5kNL._AC_SY550_.jpg", "The Avenger", "13ส.ค.2561 - 21ส.ค.2561")
	all_movie := []models.Movie{movie1, movie2, movie3, movie4, movie5, movie6, movie7}
	return all_movie, nil
}

func GetMoviePopular() ([]models.Movie, error) {
	movie1 := models.NewMovie(1, "https://www.img02.xyz/assets/movie_poster/2/201f545eebfd4154b7095b14e62f1961.jpg", "Ironman 1", "13ส.ค.2561 - 21ส.ค.2561")
	movie3 := models.NewMovie(3, "https://i.pinimg.com/originals/f8/f4/0c/f8f40cd2f8f6819c8ee6b8370dab15d4.png", "Ironman 3", "13ส.ค.2561 - 21ส.ค.2561")
	movie6 := models.NewMovie(6, "https://cf.shopee.co.th/file/df8c560c9086b80e74b9e2f896c75c25", "Captain America 3", "13ส.ค.2561 - 21ส.ค.2561")
	all_movie := []models.Movie{movie1, movie3, movie6}
	return all_movie, nil
}

func GetMovieIscoming() ([]models.Movie, error) {
	movie7 := models.NewMovie(7, "https://m.media-amazon.com/images/I/71krfsS5kNL._AC_SY550_.jpg", "The Avenger", "13ส.ค.2561 - 21ส.ค.2561")
	all_movie := []models.Movie{movie7}
	return all_movie, nil
}

func GetMovieSoldout() ([]models.Movie, error) {
	movie1 := models.NewMovie(1, "https://www.img02.xyz/assets/movie_poster/2/201f545eebfd4154b7095b14e62f1961.jpg", "Ironman 1", "13ส.ค.2561 - 21ส.ค.2561")
	movie3 := models.NewMovie(3, "https://i.pinimg.com/originals/f8/f4/0c/f8f40cd2f8f6819c8ee6b8370dab15d4.png", "Ironman 3", "13ส.ค.2561 - 21ส.ค.2561")
	movie6 := models.NewMovie(6, "https://cf.shopee.co.th/file/df8c560c9086b80e74b9e2f896c75c25", "Captain America 3", "13ส.ค.2561 - 21ส.ค.2561")
	all_movie := []models.Movie{movie1, movie3, movie6}
	return all_movie, nil
}
