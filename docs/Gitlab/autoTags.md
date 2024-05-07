# Tag your Git repo using Gitlab’s CI/CD pipeline

![img](https://miro.medium.com/max/700/0*CMIL3ryN85uG9LRq)

Photo by [Pankaj Patel](https://unsplash.com/@pankajpatel?utm_source=medium&utm_medium=referral) on [Unsplash](https://unsplash.com/?utm_source=medium&utm_medium=referral)

Git Tagging is an important feature of git in order to mark release points (v1.0, v2.0, and so on). Developers know the advantage of tagging and proper tagging is required when we release a new version.

Nowadays, DevOps engineers are handling the build and deployment process with the help of CI/CD pipelines. After each release, we may need to tag that repo. Here, we are looking at methods of tagging using Gitlab’s CI/CD pipeline code (.gitlab-ci.yml)

Here, we are describing 2 methods for tagging:

1. Using Gitlab’s username and password
2. Using Gitlab Personal Access Token

# **Method 1: Using Gitlab’s username and password**

This is the simplest method among the two, as we are directly using our Gitlab’s username and password in the pipeline code for tagging. Given below is the code sample that we can use for tagging from CI/CD pipeline:

```
image: docker:19.03.12services:
  - docker:19.03.12-dindstages: 
  - taggingTagging from pipeline: 
  stage: tagging
  script: 
    # You can write your pipeline code for build/deploy here
    - docker info
  after_script:
    - apk update && apk add git
    - git --version
    - git remote remove origin
    - git remote add origin https://username:password@gitlab.com/account-name/project-name
    - git config user.email <your-gitlab-email_id>
    - git config user.name <your-gitlab-username>
    - git tag -a v1.0 -m "Release version 1.0"
  only: 
    - master
```

**Explanation:
**The given code is a docker based pipeline code. We used the ***after_script\*** section for tagging. Here, we are installing git to the Alpine docker image. Then, we are removing the existing origin and add a new remote origin with our GitLab's username and password. We need to replace these values from the above code with our own values:

1. username
2. password
3. account-name
4. project-name
5. <your-gitlab-email_id>
6. <your-gitlab-username>

As we know, going with the simplest method has its own pros and cons.

***Pros:\***

1. No additional settings or configurations are required

***Cons:\***

1. Gitlab’s username and password are exposed in the pipeline code
2. If we change our Gitlab password, we need to change that in the pipeline code too

# **Method 2: Using Gitlab Personal Access Token**

We can use Gitlab’s Personal Access Tokens to authenticate over HTTP or SSH. You can find the details about Personal Access Tokens [here](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html). In order to use this method, we will be creating a Personal Access Token first. Given below are the steps for creating a Personal Access Token:

1. Login to your Gitlab account.
2. In the upper right corner, click your avatar and select ***Settings\***
3. On the user settings menu, select ***Access Tokens\***
4. Choose a name and an optional expiry date for the token
5. Check the scopes ***read_repository\*** and ***write_repository\*** from the list
6. Click the **Create personal access token** button
7. Save the personal access token somewhere safe. If you navigate away or refresh your page, and you did not save the token, you must create a new one

Given below is an image of the Personal Access token section in Gitlab:

![img](https://miro.medium.com/max/700/1*-0Son8eUYJcYjQ5uo8LYOA.png)

Access Token section in Gitlab

Once the Personal Access Token is created, we can use the same code from above for tagging with some slight modification. The pipeline code with Personal Access Token is given below:

```
image: docker:19.03.12services:
  - docker:19.03.12-dindstages: 
  - taggingTagging from pipeline: 
  stage: tagging
  script: 
    # You can write your pipeline code for build/deploy here
    - docker info
  after_script:
    - apk update && apk add git
    - git --version
    - git remote remove origin
    - git remote add origin https://oauth2:personal-access-token@gitlab.com/account-name/project-name
    - git config user.email <your-gitlab-email_id>
    - git config user.name <your-gitlab-username>
    - git tag -a v1.0 -m "Release version 1.0"
  only: 
    - master
```

**Explanation:
**Here, the main difference from the above method (Method 1) is, we have replaced ***username:password\*** with ***oauth2:personal-access-token\*** in the line for adding the remote origin (Remember to replace personal-access-token with the token that we created). So, the values that need to be replaced are:

1. personal-access-token
2. account-name
3. project-name
4. <your-gitlab-email_id>
5. <your-gitlab-username>

The advantages of using this method are:

1. We are not exposing Gitlab’s credentials
2. Personal Access Token can be revoked at any time and can create a new one

We can use Gitlab’s pipeline variables for storing these values (secrets) and then use those variables in our pipeline code. In this way, we can avoid using those credentials in code as plaintext.

Thanks ! ! !