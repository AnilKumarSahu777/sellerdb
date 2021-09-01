## sellerdb
### API : http://localhost:10091/

1. Build the Docker image of sellerdb -> docker build . -t sellerdb
2. Build the Docker image of sellerapp -> docker build . -t sellerapp (go to seller app directory)
3. Run docker-compose up
4. Open the postman and POST the Request to  http://localhost:10090/
5. Pass the BODY {"url":"https://www.amazon.com/s?k=nintendo+switch&ref=nb_sb_noss_1"}
```json
{"url":"https://www.amazon.com/s?k=nintendo+switch&ref=nb_sb_noss_1"}
```
### cURL Example : 
> Request
```cURL
curl --location --request POST 'http://localhost:10090/' \
--header 'Content-Type: text/plain' \
--data-raw '{"url":"https://www.amazon.com/s?k=nintendo+switch&ref=nb_sb_noss_1"}'
```

> Response 200