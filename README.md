# order-assignment

1. make app.env file with info from config.json file.
```
port=8080
db_User=db_user
db_Password=kothinPassword
db_Host=172.18.0.2
db_Port=3306
db_Schema=order
secret_key=kothinKey
```

2. give permission to build_and_push.sh to execute. 
```
chmod +x build_and_pushs.sh
```

3. run build_and_push.sh to run database and app in containers.
```
./build_and_push.sh
```

4. inspect to check ip of the both containers.
```
docker inspect <container_name>
```

5. now trigger appropriate apis to get result.

6. api documentation is attached.


