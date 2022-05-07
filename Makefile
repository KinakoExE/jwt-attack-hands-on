TOKEN := $(TOKEN)

local-none-attack:
	cd none-attack/ && go run main.go

local-brute-force-secret:
	cd brute-force-secret/ && go run main.go

docker-none-attack:
	docker build --target none-attack -t kinakoexe/none-attack:latest .
	docker run -it -p 5555:5555 kinakoexe/none-attack:latest

docker-brute-force-secret:
	docker build --target brute-force-secret -t kinakoexe/brute-force-secret:latest .
	docker run -it -p 5555:5555 kinakoexe/brute-force-secret:latest

# make send-token TOKEN="<jwt token>"
send-token:
	curl http://localhost:5555/admin -H "Authorization: Bearer $(TOKEN)"


get-token:
	curl http://localhost:5555/token
