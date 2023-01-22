# TODO

- [ ] Write backend kubernetes manifest
- [ ] Attach database

# Snippets

```
docker run --name mariadb --rm -ti --network host -e MYSQL_RANDOM_ROOT_PASSWORD=yes -e MYSQL_DATABASE=quotes -e MYSQL_USER=quotes -e MYSQL_PASSWORD=password mariadb:10
```

```
docker run --rm -ti --name adminer --network host adminer:4
```
