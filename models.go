package rawg

// Role: Position (job) of the video game creator
type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// Creator: Main information about creator
type Creator struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Image           string `json:"results"`
	ImageBackground string `json:"image_background"`
	GamesCount      int    `json:"games_count"`
}

// CreatorDetailed: Detailed information about creator
type CreatorDetailed struct {
	Creator
	Description  string   `json:"description"`
	ReviewsCount int      `json:"reviews_count"`
	Rating       string   `json:"rating"`
	RatingTop    int      `json:"rating_top"`
	Updated      DateTime `json:"updated"`
}

// Developer: Main information about video game developer
type Developer struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	GamesCount      int    `json:"games_count"`
	ImageBackground string `json:"image_background"`
}

// DeveloperDetailed: Detailed information about video game developer
type DeveloperDetailed struct {
	Developer
	Description string `json:"description"`
}

// Genre: Game genre
type Genre struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	GamesCount      int    `json:"games_count"`
	ImageBackground string `json:"image_background"`
}

// GenreDetailed: Detailed information of video game genre
type GenreDetailed struct {
	Genre
	Description string `json:"description"`
}

// Platform: Game platform
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

// PlatformDetailed: Detailed information of video game platform
type PlatformDetailed struct {
	Platform
	Description string `json:"description"`
}

// ParentPlatform: Parent video game platform. For instance, for PS2 and PS4 the “parent platform” is PlayStation.
type ParentPlatform struct {
	Platform struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"platform"`
}

// Publisher: Main information about video game publisher
type Publisher struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	GamesCount      int    `json:"games_count"`
	ImageBackground string `json:"image_background"`
}

// PublisherDetailed: Detailed information about video game developer
type PublisherDetailed struct {
	Publisher
	Description string `json:"description"`
}

// Store: Video game storefront
type Store struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Domain          string `json:"domain"`
	Slug            string `json:"slug"`
	GamesCount      int    `json:"games_count"`
	ImageBackground string `json:"image_background"`
}

// StoreDetailed: Video game storefront (detailed information)
type StoreDetailed struct {
	Store
	Description string `json:"description"`
}

// Tag: Tag of the video game
type Tag struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	GamesCount      int    `json:"games_count"`
	ImageBackground string `json:"image_background"`
	Language        string `json:"language"`
}

// TagDetailed: Tag of the video game (detailed information)
type TagDetailed struct {
	Tag
	Description string `json:"description"`
}

// EsrbRating: Entertainment Software Rating Board
type EsrbRating struct {
	ID   int    `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// GameStore: "Video game - storefront" relation
type GameStore struct {
	ID      int    `json:"id"`
	GameID  int    `json:"game_id"`
	StoreID int    `json:"store_id"`
	Url     string `json:"url"`
}

// Screenshot of the game
type Screenshot struct {
	ID        int    `json:"id"`
	Image     string `json:"image"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	IsDeleted bool   `json:"is_deleted"`
}

// Rating of the game
type Rating struct {
	ID      int     `json:"id"`
	Title   string  `json:"title"`
	Count   int     `json:"count"`
	Percent float32 `json:"percent"`
}

// AddedByStatus: Number of games added by status
type AddedByStatus struct {
	Yet     int `json:"yet"`
	Owned   int `json:"owned"`
	Beaten  int `json:"beaten"`
	Toplay  int `json:"toplay"`
	Dropped int `json:"dropped"`
	Playing int `json:"playing"`
}

// Requirement: Game requirements
type Requirement struct {
	Minimum     string `json:"minimum"`
	Recommended string `json:"recommended"`
}

// Position (job)
type Position struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// Clip (Trailer) of the game
type Clip struct {
	Clip    string            `json:"clip"`
	Clips   map[string]string `json:"clips"`
	Video   string            `json:"video"`
	Preview string            `json:"preview"`
}

// MetacriticPlatform: Metascore value of the game
type MetacriticPlatform struct {
	Metascore int    `json:"metascore"`
	Url       string `json:"url"`
	Platform  *struct {
		Platform int    `json:"platform"`
		Name     string `json:"name"`
		Slug     string `json:"slug"`
	} `json:"platform"`
}

// Game: Main information of the game
type Game struct {
	ID               int            `json:"id"`
	Slug             string         `json:"slug"`
	Name             string         `json:"name"`
	Released         DateTime       `json:"released"`
	Tba              bool           `json:"tba"`
	ImageBackground  string         `json:"background_image"`
	Rating           float32        `json:"rating"`
	RatingTop        int            `json:"rating_top"`
	Ratings          []*Rating      `json:"ratings"`
	RatingsCount     int            `json:"ratings_count"`
	ReviewsTextCount int            `json:"reviews_text_count"`
	Added            int            `json:"added"`
	AddedByStatus    *AddedByStatus `json:"added_by_status"`
	Metacritic       int            `json:"metacritic"`
	Playtime         int            `json:"playtime"`
	SuggestionsCount int            `json:"suggestions_count"`
	ReviewsCount     int            `json:"reviews_count"`
	SaturatedColor   string         `json:"saturated_color"`
	DominantColor    string         `json:"dominant_color"`
	Platforms        []*struct {
		Platform       *Platform    `json:"platform"`
		ReleasedAt     DateTime     `json:"released_at"`
		RequirementsEn *Requirement `json:"requirements_en"`
		RequirementsRu *Requirement `json:"requirements_ru"`
	} `json:"platforms"`
	ParentPlatforms []*struct {
		Platform struct {
			ID   int    `json:"id"`
			Slug string `json:"slug"`
			Name string `json:"name"`
		}
	} `json:"parent_platforms"`
	Genres []*Genre `json:"genres"`
	Stores []*struct {
		ID    int    `json:"id"`
		Store *Store `json:"store"`
		UrlEn string `json:"url_en"`
		UrlRu string `json:"url_ru"`
	} `json:"stores"`
	Clip             *Clip  `json:"clip"`
	Tags             []*Tag `json:"tags"`
	ShortScreenshots []*struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"short_screenshots"`
}

// GameDetailed: Detailed information of the game
type GameDetailed struct {
	ID                        int                   `json:"id"`
	Slug                      string                `json:"slug"`
	Name                      string                `json:"name"`
	NameOriginal              string                `json:"name_original"`
	Description               string                `json:"description"`
	Metacritic                int                   `json:"metacritic"`
	MetacriticPlatforms       []*MetacriticPlatform `json:"metacritic_platforms"`
	Released                  DateTime              `json:"released"`
	Tba                       bool                  `json:"tba"`
	Updated                   DateTime              `json:"updated"`
	ImageBackground           string                `json:"background_image"`
	ImageBackgroundAdditional string                `json:"background_image_additional"`
	Website                   string                `json:"website"`
	Rating                    float32               `json:"rating"`
	RatingTop                 int                   `json:"rating_top"`
	Ratings                   []*Rating             `json:"ratings"`
	Reactions                 map[string]int        `json:"reactions"`
	Added                     int                   `json:"added"`
	AddedByStatus             *AddedByStatus        `json:"added_by_status"`
	Playtime                  int                   `json:"playtime"`
	ScreenshotsCount          int                   `json:"screenshots_count"`
	MoviesCount               int                   `json:"movies_count"`
	CreatorsCount             int                   `json:"creators_count"`
	AchievementsCount         int                   `json:"achievements_count"`
	ParentAchievementsCount   int                   `json:"parent_achievements_count"`
	RedditUrl                 string                `json:"reddit_url"`
	RedditName                string                `json:"reddit_name"`
	RedditDescription         string                `json:"reddit_description"`
	RedditLogo                string                `json:"reddit_logo"`
	RedditCount               int                   `json:"reddit_count"`
	TwitchCount               int                   `json:"twitch_count"`
	YoutubeCount              int                   `json:"youtube_count"`
	ReviewsTextCount          int                   `json:"reviews_text_count"`
	RatingsCount              int                   `json:"ratings_count"`
	SuggestionsCount          int                   `json:"suggestions_count"`
	AlternativeNames          []string              `json:"alternative_names"`
	MetacriticUrl             string                `json:"metacritic_url"`
	ParentsCount              int                   `json:"parents_count"`
	AdditionsCount            int                   `json:"additions_count"`
	GameSeriesCountCount      int                   `json:"game_series_count"`
	ReviewsCount              int                   `json:"reviews_count"`
	SaturatedColor            string                `json:"saturated_color"`
	DominantColor             string                `json:"dominant_color"`
	ParentPlatforms           []*struct {
		Platform struct {
			ID   int    `json:"id"`
			Slug string `json:"slug"`
			Name string `json:"name"`
		}
	} `json:"parent_platforms"`
	Platforms []*struct {
		Platform     *Platform    `json:"platform"`
		ReleasedAt   DateTime     `json:"released_at"`
		Requirements *Requirement `json:"requirements"`
	} `json:"platforms"`
	Stores []*struct {
		ID    int    `json:"id"`
		Url   string `json:"url"`
		Store *Store `json:"store"`
	} `json:"stores"`
	Developers     []*Developer `json:"developers"`
	Genres         []*Genre     `json:"genres"`
	Tags           []*Tag       `json:"tags"`
	Publishers     []*Publisher `json:"publishers"`
	EsrbRating     *EsrbRating  `json:"esrb_rating"`
	Clip           *Clip        `json:"clip"`
	DescriptionRaw string       `json:"description_raw"`
}

// Achievement (a meta-goal defined outside a game's parameters)
type Achievement struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Percent     float32 `json:"percent,string"`
}

// Movie of the game
type Movie struct {
	ID      int               `json:"id"`
	Name    string            `json:"name"`
	Preview string            `json:"preview"`
	Data    map[string]string `json:"data"`
}

// Reddit comment
type Reddit struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Text        string   `json:"text"`
	Image       string   `json:"image"`
	Url         string   `json:"url"`
	Username    string   `json:"username"`
	UsernameUrl string   `json:"username_url"`
	Created     DateTime `json:"created"`
}

// Twitch stream
type Twitch struct {
	ID          int      `json:"id"`
	ExternalID  int      `json:"external_id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Created     DateTime `json:"created"`
	Published   DateTime `json:"published"`
	Thumbnail   string   `json:"thumbnail"`
	ViewCount   int      `json:"view_count"`
	Language    string   `json:"language"`
}

// YoutubeThumbnail: Thumbnail of the youtube video
type YoutubeThumbnail struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Youtube video associated with the game
type Youtube struct {
	ID            int      `json:"id"`
	ExternalID    string   `json:"external_id"`
	ChannelID     string   `json:"channel_id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Created       DateTime `json:"created"`
	ViewCount     int      `json:"view_count"`
	CommentsCount int      `json:"comments_count"`
	LikeCount     int      `json:"like_count"`
	DislikeCount  int      `json:"dislike_count"`
	FavoriteCount int      `json:"favorite_count"`
	Thumbnails    *struct {
		High          *YoutubeThumbnail
		Medium        *YoutubeThumbnail
		Default       *YoutubeThumbnail
		SdDefault     *YoutubeThumbnail
		MaxResDefault *YoutubeThumbnail
	}
}

// GameDeveloper: Detailed information of game developer (with games list)
type GameDeveloper struct {
	ID              int         `json:"id"`
	Name            string      `json:"name"`
	Slug            string      `json:"slug"`
	Image           string      `json:"image"`
	ImageBackground string      `json:"image_background"`
	GamesCount      int         `json:"games_count"`
	Games           []*Game     `json:"games"`
	Positions       []*Position `json:"positions"`
}
