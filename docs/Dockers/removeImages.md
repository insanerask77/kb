### Danglin images remove



```
docker rmi -f $(docker images -f "dangling=true" -q)
docker system prune -af
```

