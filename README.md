# Barulho

## Goal

API to query by video by song name, artist or the city.

## Launch

```bash
./db
go run app/main.go
```

Query the API:
```bash
curl -XPOST -d '{"query": "{ hello }"}' localhost:8080/query`
curl -XPOST -d '{"query": "{ getVideosBySongName(name: \"Dance\") { name cityName artistName } }"}' localhost:8080/query
```

## Notes

A video record example:
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
