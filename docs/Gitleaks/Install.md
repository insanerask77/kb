These instructions are intended to be followed in a development/localhost environment to guarantee that the committed changes are not leaking secrets.



**gitleaks has support for Github Actions:** [![(https://github.com/fluidicon.png)

### Requirements:

- Python3.8^
- [gitleaks](https://github.com/gitleaks/gitleaks)
- [pre-commit](https://pre-commit.com/) python library

 

### Setup

1. Install pre-commit 

   ```
   python -m pip install pre-commit
   ```

2. Navigate to the root folder of the repository and create a `.pre-commit-config.yaml` file with the following content: 

```
repos:
  - repo: local
    hooks:
      - id: gitleaks
        name: check for secrets
        entry: gitleaks detect --verbose
        language: system
```



1. Setup the git hook 

```
pre-commit autoupdate && pre-commit install
```

 

With this setup, every time you invoke the `git commit` command, the hook will be called and the repository will be scanned for secrets or passwords.

### Additional information

Given the fact that we rely on different `.env` files for development and testing purposes, it is possible to add the `exclude` attribute within the `pre-commit-config.yaml` file to skip these in case it is necessary to add them.

### Installation via Scripts

We have a new folder in the repostory named .gitleaks inside we have a copuple of scripts to setup our pecommit.

First of all execute gitleaks installation and after that the precommit installation.

### Troubleshooting

Gitleaks detecting me old leaks:

In this case we have to update the baseline, sometimes it can happen if we are working on an old branch or some secret was pushed into another branch that was merged to develop or main. We just need to update the gitleaks baseline with this command.



```
gitleaks detect --report-path ./baseline.json
```

After you push all the changes applied to the baseline file and its attached job.

I want to make sure that I don’t have any leak in my code.

Before committing your job you can scan your branch manually with this command.





```
gitleaks detect -b baseline.json -v
```

If you detect any new leaks, contact your repository administrator or manager to revoke this leaked secret.