# atom
atom go http service template


## docker compose 
```
version: '3'
services:
  db:
    image: mariadb:latest
    restart: always
    ports:
      - 3306:3306
    environment:
      MARIADB_DATABASE: demos
      MARIADB_ROOT_HOST: "%"
      MARIADB_ROOT_PASSWORD: root
```