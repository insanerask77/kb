## Clone all repos in a organization

#### http

```
curl -s -H "Authorization: token $YOUR_TOKEN" \
"https://api.github.com/orgs/$YOUR_ORG/repos?per_page=100"\
| jq -r ".[].clone_url" \
| xargs -L1 git clone
```

ssh

```
curl -s -H "Authorization: token $YOUR_TOKEN" \
"https://api.github.com/orgs/$YOUR_ORG/repos?per_page=100"\
| jq -r ".[].ssh_url" \
| xargs -L1 git clone
```



