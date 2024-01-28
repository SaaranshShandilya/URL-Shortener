### A simple URL shortener made in golang.
- Takes the original full url as input
- Saves the original url in mongo db
- Returns the shortened url
- When the shortened url is used with the base URL, the user is directed to the original URL

1. Hit a post request at localhost:8080/shorten

Sample request:
```
{
    "URL":"https://github.com"
}
```
Sample response:
```
{
    "table": {
        "FullUrl": "https://github.com",
        "ShortUrl": "kKGKFWJVXj"
    }
}
```
2. Use the shortened url that will redirect you to the base URL

```
localhost:8080/kKGKFWJVXj
```

