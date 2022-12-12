# Church

A project demo grpc and RESTFULL api with caddy reverse proxy


## Installation

If you want to deploy with docker

```bash
    cd ./deployments
    docker-compose up --force-recreate -d 
    
    #NOTE: if you docker-compose down it'll not re init the database
    #incase you want to re init database
    docker-compose down --volumes
```

This project required Go 1.18.3 and Caddy

To install Caddy (Ubuntu)

``` bash
sudo apt install -y debian-keyring debian-archive-keyring apt-transport-https
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list
sudo apt update
sudo apt install caddy
```

Copy Caddy file into Caddy config directory

```bash
 vm /etc/caddy/Caddyfile /etc/caddy/Caddyfile.bak
 cp ./deployments/Caddyfile /etc/caddy/Caddyfile
```

Start Caddy service

```bash
sudo systemctl start caddy
```

Run project

```bash
cd service/minhdq
go build cmd/church/main.go
sudo chmod +x ./main

#command argument

# help - list commands
./main --help

# serve - serve http server
./main --env ./configs/.env serve

# authen - serve grpc server
./main --env ./configs/.env authen

# client - grpc client
./main --env ./configs/.env client

```

## Testing

if you want to test grpc more specifically user grpcurl

Install grpcurl

```bash
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

Testing grpc server

```bash
grpcurl -plaintext -d '{"LoginID":"12", "Password":"password"}'  localhost:8888 authentication.Resgister.Resgis
```