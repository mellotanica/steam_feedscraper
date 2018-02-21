package scrapers

import (
	"github.com/mmcdole/gofeed"
	"feedscraper/games_cache"
)

func ParseFitGirlRepack(item *gofeed.Item) (*games_cache.Game, error){
	name := ""
	genre := ""
	link := ""

	return gameFromLinkAndName(link, name, genre)
}
