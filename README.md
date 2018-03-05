# steam_feedscraper
steam_feedscraper is a simple [Go](https://golang.org) http proxy adding browsing functionalities to steam store based on news feeds scraping.
Running `go run application.go` and visiting `http://localhost:8080/review` will lead you to the steam store page of a game extracted from the feeds that has not already been reviewed through the app and is not already in your wishlist.

This tool works great alongside with [steam_deallist](https://github.com/mellotanica/steam_deallist).

## Scrapers

Two demo scrapers are already implemented reading from games updates sites.
The steam store links are extracted from the feeds or queried directly to steam based on game name if link is missing, deleting duplicated entries.
An optional steam wishlist is scraped too to avoid inserting already wanted games into the list.
Already reviewed items are kept in a persistent storage, filtering them out from the active pending list.

## steam store enhance

An http proxy forwards requests on the port specified as argument (`8080` is the default) and modifies steam store pages adding two buttons and two shortcuts for games that belong to the active pending list (the proxy is completely transparent for any other webpage).
The `Next` button (`n` key) marks the game as reviewed and loads the next entry from the list, the `Wishlist` button (`w` key) adds the game to the wishlist if the user has logged in on steam store.
The main entry point for games review is the `/review` endpoint.