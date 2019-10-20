start: generate-cert server client 

server: 
	nohup go run ./server.go & echo $! > pid.txt

client:
	go run ./client.go

generate-cert:
	openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt -subj '/CN=localhost'

quit:
	kill -9 $(cat pid.txt)
	rm -f pid.txt