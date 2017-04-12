# tvdb [![CircleCI](https://circleci.com/gh/danesparza/tvdb.svg?style=svg)](https://circleci.com/gh/danesparza/tvdb)
TVDB v2.0 API wrapper for Go

## Example

``` Go
// Create a request & a client
request := tvdb.SearchRequest{Name: "Looney Tunes"}
client := tvdb.Client{}

//  Search for the series
matches, err := client.SeriesSearch(request)

//  Check for errors or use matches...
if matches[0].ID != 72514 {
  log.Printf("Didn't get the series ID back that we expected")
}
```
