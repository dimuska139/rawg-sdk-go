package rawg_sdk_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewGamesFilter(t *testing.T) {
	from1, _ := time.Parse("2006-01-02", "2010-01-01")
	to1, _ := time.Parse("2006-01-02", "2018-12-31")
	dateRangeFirst := DateRange{
		From: from1,
		To:   to1,
	}

	from2, _ := time.Parse("2006-01-02", "1960-01-01")
	to2, _ := time.Parse("2006-01-02", "1969-12-31")
	dateRangeSecond := DateRange{
		From: from2,
		To:   to2,
	}

	filter := NewGamesFilter().
		SetPage(1).
		SetPageSize(2).
		SetSearch("gta5").
		SetParentPlatforms(1, 2).
		SetPlatforms(3, 4).
		SetStores(5, 6).
		SetDevelopers(7, "feral-interactive").
		SetPublishers(8, "electronic-arts").
		SetGenres(9, "action", "indie").
		SetTags("singleplayer", 31).
		SetCreators(28, "mike-morasky").
		SetDates(&dateRangeFirst, &dateRangeSecond).
		SetPlatformsCount(10).
		ExcludeCollection(123).
		WithoutAdditions().
		WithoutParents().
		WithoutGameSeries().
		SetOrdering("-name")

	assert.Equal(t, map[string]interface{}{
		"creators":            "28,mike-morasky",
		"dates":               "2010-01-01,2018-12-31.1960-01-01,1969-12-31",
		"developers":          "7,feral-interactive",
		"exclude_additions":   true,
		"exclude_collection":  123,
		"exclude_game_series": true,
		"exclude_parents":     true,
		"genres":              "9,action,indie",
		"ordering":            "-name",
		"page":                1,
		"page_size":           2,
		"parent_platforms":    "1,2",
		"platforms":           "3,4",
		"platforms_count":     10,
		"publishers":          "8,electronic-arts",
		"search":              "gta5",
		"stores":              "5,6",
		"tags":                "singleplayer,31",
	}, filter.GetParams())

}
