package rawg

import (
	"fmt"
	"strings"
	"time"
)

// DateRange - two dates: from and to
type DateRange struct {
	From time.Time
	To   time.Time
}

// GamesFilter: Filter to search games in RAWG
type GamesFilter struct {
	page              int
	pageSize          int
	search            string
	parentPlatforms   []int
	platforms         []int
	stores            []int
	developers        []interface{}
	publishers        []interface{}
	genres            []interface{}
	tags              []interface{}
	creators          []interface{}
	dates             []*DateRange
	platformsCount    int
	excludeCollection int
	excludeAdditions  bool
	excludeParents    bool
	excludeGameSeries bool
	ordering          string
}

// NewGamesFilter creates filter
func NewGamesFilter() *GamesFilter {
	return &GamesFilter{}
}

// SetPage sets "page" parameter
func (filter *GamesFilter) SetPage(page int) *GamesFilter {
	filter.page = page
	return filter
}

// SetPageSize sets "page_size" parameter
func (filter *GamesFilter) SetPageSize(pageSize int) *GamesFilter {
	filter.pageSize = pageSize
	return filter
}

// SetSearch sets "search" parameter
func (filter *GamesFilter) SetSearch(search string) *GamesFilter {
	filter.search = search
	return filter
}

// SetParentPlatforms sets "parent_platforms" parameter
func (filter *GamesFilter) SetParentPlatforms(parentPlatforms ...int) *GamesFilter {
	filter.parentPlatforms = parentPlatforms
	return filter
}

// SetPlatforms sets "platforms" parameter
func (filter *GamesFilter) SetPlatforms(platforms ...int) *GamesFilter {
	filter.platforms = platforms
	return filter
}

// SetStores sets "stores" parameter
func (filter *GamesFilter) SetStores(stores ...int) *GamesFilter {
	filter.stores = stores
	return filter
}

// SetDevelopers sets "developers" parameter
func (filter *GamesFilter) SetDevelopers(developers ...interface{}) *GamesFilter {
	filter.developers = developers
	return filter
}

// SetPublishers sets "publishers" parameter
func (filter *GamesFilter) SetPublishers(publishers ...interface{}) *GamesFilter {
	filter.publishers = publishers
	return filter
}

// SetGenres sets "genres" parameter
func (filter *GamesFilter) SetGenres(genres ...interface{}) *GamesFilter {
	filter.genres = genres
	return filter
}

// SetTags sets "tags" parameter
func (filter *GamesFilter) SetTags(tags ...interface{}) *GamesFilter {
	filter.tags = tags
	return filter
}

// SetCreators sets "creators" parameter
func (filter *GamesFilter) SetCreators(creators ...interface{}) *GamesFilter {
	filter.creators = creators
	return filter
}

// SetDates sets "dates" parameter
func (filter *GamesFilter) SetDates(ranges ...*DateRange) *GamesFilter {
	filter.dates = ranges
	return filter
}

// SetPlatformsCount sets "platforms_count" parameter
func (filter *GamesFilter) SetPlatformsCount(count int) *GamesFilter {
	filter.platformsCount = count
	return filter
}

// ExcludeCollection sets "exclude_collection" parameter
func (filter *GamesFilter) ExcludeCollection(collection int) *GamesFilter {
	filter.excludeCollection = collection
	return filter
}

// WithoutAdditions sets "exclude_additions" parameter
func (filter *GamesFilter) WithoutAdditions() *GamesFilter {
	filter.excludeAdditions = true
	return filter
}

// WithoutParents sets "exclude_parents" parameter
func (filter *GamesFilter) WithoutParents() *GamesFilter {
	filter.excludeParents = true
	return filter
}

// WithoutGameSeries sets "exclude_game_series" parameter
func (filter *GamesFilter) WithoutGameSeries() *GamesFilter {
	filter.excludeGameSeries = true
	return filter
}

// SetOrdering sets results ordering
func (filter *GamesFilter) SetOrdering(ordering string) *GamesFilter {
	filter.ordering = ordering
	return filter
}

// GetParams returns filter parameters as Map
func (filter *GamesFilter) GetParams() map[string]interface{} {
	params := make(map[string]interface{})

	params["page"] = 1
	if filter.page != 0 {
		params["page"] = filter.page
	}

	params["page_size"] = 20
	if filter.pageSize != 0 {
		params["page_size"] = filter.pageSize
	}

	if filter.search != "" {
		params["search"] = filter.search
	}

	if len(filter.parentPlatforms) != 0 {
		params["parent_platforms"] = strings.Trim(strings.Replace(fmt.Sprint(filter.parentPlatforms), " ", ",", -1), "[]")
	}

	if len(filter.platforms) != 0 {
		params["platforms"] = strings.Trim(strings.Replace(fmt.Sprint(filter.platforms), " ", ",", -1), "[]")
	}

	if len(filter.stores) != 0 {
		params["stores"] = strings.Trim(strings.Replace(fmt.Sprint(filter.stores), " ", ",", -1), "[]")
	}

	if len(filter.developers) != 0 {
		params["developers"] = strings.Trim(strings.Replace(fmt.Sprint(filter.developers), " ", ",", -1), "[]")
	}

	if len(filter.publishers) != 0 {
		params["publishers"] = strings.Trim(strings.Replace(fmt.Sprint(filter.publishers), " ", ",", -1), "[]")
	}

	if len(filter.genres) != 0 {
		params["genres"] = strings.Trim(strings.Replace(fmt.Sprint(filter.genres), " ", ",", -1), "[]")
	}

	if len(filter.tags) != 0 {
		params["tags"] = strings.Trim(strings.Replace(fmt.Sprint(filter.tags), " ", ",", -1), "[]")
	}

	if len(filter.creators) != 0 {
		params["creators"] = strings.Trim(strings.Replace(fmt.Sprint(filter.creators), " ", ",", -1), "[]")
	}

	if len(filter.dates) != 0 {
		parts := make([]string, 0)
		for _, dateRange := range filter.dates {
			parts = append(parts, dateRange.From.Format("2006-01-02")+","+dateRange.To.Format("2006-01-02"))
		}

		params["dates"] = strings.Join(parts, ".")
	}

	if filter.platformsCount != 0 {
		params["platforms_count"] = filter.platformsCount
	}

	if filter.excludeCollection != 0 {
		params["exclude_collection"] = filter.excludeCollection
	}

	if filter.excludeAdditions {
		params["exclude_additions"] = filter.excludeAdditions
	}

	if filter.excludeParents {
		params["exclude_parents"] = filter.excludeParents
	}

	if filter.excludeGameSeries {
		params["exclude_game_series"] = filter.excludeGameSeries
	}

	if filter.ordering != "" {
		params["ordering"] = filter.ordering
	}

	return params
}
