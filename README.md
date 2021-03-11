# iso8601ms #

Package iso8601 is a simple Go package for encoding `time.Time` in JSON in ISO 8601
format with millisecond precision, with the converted UTC timezone.

Standard iso8601 with millisecond precision: [https://docs.jsonata.org/date-time](https://docs.jsonata.org/date-time)

```go
// Local time zone
// Marshal to "2021-03-11T01:14:28.625-07:00"
t := iso8601ms.Time(time.Now())
jsonBytes, _ := json.Marshal(t)
fmt.Println(string(jsonBytes))
```

```go
// UTC
// Marshal to "2021-03-11T01:14:28.625Z"
t := iso8601ms.Time(time.Now().UTC())
jsonBytes, _ := json.Marshal(t)
fmt.Println(string(jsonBytes))
```
