# Twitter

Golang Microservice to tweet

This microservice will be used as an microservice example in this event: https://www.eventbrite.com.br/e/minicurso-microservicos-do-conceito-ao-codigo-tickets-37631005350

## Usage
```shell
curl -i -XPOST -H 'X-ConsumerKey: consumerKey' -H 'X-ConsumerSecret: consumerSecret' -H 'X-AccessToken: accessToken' -H 'X-AccessTokenSecret: accessTokenSecret' -d '{ "message":"Testing again using json2"  }' 'http://localhost:8080/posts'
```
