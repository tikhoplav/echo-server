# Go Echo Server

> Made as a demo to [this project](https://www.upwork.com/jobs/microservice-GET-request-reply-with-requester_~0168da9362b3e6d72a)

## Usage example

```
$ docker-compose up
$ curl "localhost:8080/?ip&referrer&user-agent" -H "Referer: http://example.com"

{
	"agent": "curl/7.58.0",
	"ip": "127.0.0.1",
	"referrer": "http://example.com"
}
```
