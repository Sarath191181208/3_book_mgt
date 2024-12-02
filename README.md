### Running the app 

```bash 
# start the mysql server
sudo docker run -d --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=db mysql 

# set the env vars in the current shell
 set -a && source .env && set +a

# run the go file 
go run ./cmd/main.go
```
