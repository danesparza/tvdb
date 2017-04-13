# tvdb [![CircleCI](https://circleci.com/gh/danesparza/tvdb.svg?style=svg)](https://circleci.com/gh/danesparza/tvdb)
TVDB v2.0 API wrapper for Go

## Example

``` Go
import "github.com/danesparza/tvdb"

...

// Create a client and search request
client := tvdb.Client{}
request := tvdb.SeriesEpisodesRequest{Name: "Looney Tunes"}

//  Search for the series
responses, err := client.SeriesSearch(request)

//  Check for errors or use responses...
if responses[0].ID != 72514 {
  log.Printf("Didn't get the series ID back that we expected")
}
  
```
