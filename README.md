# Prototype WebApp Gin

- start using ``go run main.go``


## Example requests

- create user: 
```
curl --request POST \
  --url http://localhost:8080/users \
  --header 'Content-Type: application/json' \
  --data '{
    "id": 1,
    "firstname": "domi",
    "lastname": "illi",
    "fullname": "domi illi",
    "email": "hallo@example.com"
  }'
```

- get user: ``curl --request GET --url http://localhost:8080/users/1``