package rawg_sdk_go

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Creator struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Image           string `json:"results"`
	ImageBackground string `json:"image_background"`
	GamesCount      int    `json:"games_count"`
}

type CreatorDetailed struct {
	Creator
	Description  string   `json:"description"`
	ReviewsCount int      `json:"reviews_count"`
	Rating       string   `json:"rating"`
	RatingTop    int      `json:"rating_top"`
	Updated      DateTime `json:"updated"`
}

type Developer struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	GamesCount      int    `json:"games_count"`
	ImageBackground string `json:"image_background"`
}

type DeveloperDetailed struct {
	Developer
	Description string `json:"description"`
}

type Genre struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	GamesCount      int    `json:"games_count"`
	ImageBackground string `json:"image_background"`
}

type GenreDetailed struct {
	Genre
	Description string `json:"description"`
}

type Platform struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	GamesCount      int    `json:"games_count"`
	ImageBackground string `json:"image_background"`
	Image           string `json:"image"`
	YearStart       int    `json:"year_start"`
	YearEnd         int    `json:"year_end"`
}

type PlatformDetailed struct {
	Platform
	Description string `json:"description"`
}

type Publisher struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	GamesCount      int    `json:"games_count"`
	ImageBackground string `json:"image_background"`
}

type PublisherDetailed struct {
	Publisher
	Description string `json:"description"`
}

type Store struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Domain          string `json:"domain"`
	Slug            string `json:"slug"`
	GamesCount      int    `json:"games_count"`
	ImageBackground string `json:"image_background"`
}

type StoreDetailed struct {
	Store
	Description string `json:"description"`
}

type Tag struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	GamesCount      int    `json:"games_count"`
	ImageBackground string `json:"image_background"`
	Language        string `json:"language"`
}

type TagDetailed struct {
	Tag
	Description string `json:"description"`
}
