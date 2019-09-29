# Barulho

## Goal

API to query by video by song name, artist or the city.

## Launch

```bash
./db
go run app/main.go
```

Query the API:

The GraphQL queries are defined in the `schema.graphql` file.

```bash
curl -XPOST -d '{"query": "{ hello }"}' localhost:8080/query`

|> curl -XPOST -d '{"query": "{ getVideosBySongName(name: \"Dance\") { name cityName artistName } }"}' localhost:8080/query
{"data":{"getVideosBySongName":[{"name":"I Wanna Dance With Somebody (Whitney Houston Cover)","cityName":"Los Angeles","artistName":"Marcus Very Ordinary"},{"name":"Dance Dance Dance","cityName":"Melbourne","artistName":"Blackchords"},{"name":"Unwillingness To Dance","cityName":"Oxford","artistName":"FaceOmeter"},{"name":"Worry Dance","cityName":"Minneapolis","artistName":"Diet Folk"},{"name":"Dance On My Own","cityName":"London","artistName":"M.O"},{"name":"Bedroom Dance","cityName":"Auckland","artistName":"Napier Avalon Jones"},{"name":"Dancehall","cityName":"Brighton","artistName":"Tribes"},{"name":"And So We Dance","cityName":"NYC","artistName":"Kellylee Evans"}]}}

|> curl -XPOST -d '{"query": "{ getVideosByArtistName(name: \"Whitney\") { name cityName artistName } }"}' localhost:8080/query
{"data":{"getVideosByArtistName":[{"name":"Mountain House","cityName":"Seattle","artistName":"Whitney Lyman"}]}}
```

## Notes

A video record example from the data source:
```json
{
  "video_uid": "wWRIyRRyTCI",
  "song": {
    "id": 3976,
    "artist_id": 2745,
    "title": "Brunch",
    "cached_slug": "brunch",
    "city_id": 58,
    "artist": {
      "id": 2745,
      "title": "Xoe Wise",
      "cached_slug": "xoe-wise"
    },
    "city": {
      "id": 58,
      "title": "Chicago",
      "cached_slug": "chicago"
    }
  }
}
```
