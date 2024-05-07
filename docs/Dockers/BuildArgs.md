

# Docker build with --build-arg with multiple arguments from file



Posted *Nov 3, 2018*

By *[André Ilhicas dos Santos](https://twitter.com/ilhicas)*

*1 min* read

# Introduction

On my last post I described how to use a **single** `ARG` with a list of packages to be installed on an alpine image.



But what happens when your Dockefile contains a multitude of `BUILD ARGS` that you wish to pass along the build from a single file that contains them.

Let’s check an rather simple Dockerfile what I’m trying to explain here.

```
1 2 3 4 5 6 7 8 9 10 11 12 13 FROM python:latest ARG ARG_1=first_argument ARG ARG_2=second_argument ARG ARG_3=third_argument ARG ARG_4=fourth_argument ARG ARG_12=first_second_argument ARG ARG_22=second_second_argument ARG ARG_32=third_second_argument ARG ARG_42=fourth_second_argument ARG ARG_13=first_third_argument ARG ARG_23=second_third_argument ARG ARG_33=third_third_argument ARG ARG_43=fourth_third_argument 
```

If you wan’t to override all this arguments using the `--build-arg` would require you to have the following command (or similar)

```
docker build -t custom:stuff --build-arg ARG_1=1st --build-arg ARG_2=2nd --build-arg ARG_3=3rd ... #you get the point 
```

Now imagine you would want to have those build args defined in a file similar to a `.env` or `env_file` that you wish to use in order to have different multiple builds, for staging production etc.. , as that would be much more pratical.

## –build-arg with multiple arguments from file with bash

> This could obvisouly be achieved on many other terminal interpreters, for simplicity and since my main OS is usually Linux with bash as interpreter, I’ve done it in bash

Now let’s create a file named `build.args` with the following contents

```
ARG_1=1st ARG_2=2nd ARG_3=3rd ARG_4=4th ARG_12=1st2nd ARG_22=2nd2nd ARG_32=3rd2nd ARG_42=4th2nd ARG_13=1st3rd ARG_23=2nd3rd ARG_33=3rd3rd ARG_43=4th3rd 
```

Now, how to use this `build.args` to build an image and its respective build args from this file.

```
1 docker build -t custom:stuff $(for i in `cat build.args`; do out+="--build-arg $i " ; done; echo $out;out="") . 
```

```
out=""; for i in $(cat ./build.args); do out="$out--build-arg $i "; done; echo "$out"; out=""
```

It might not be the simplest single line command to read, but when using on CI/CD pipeline for instance, you could have multiple build environments from multiple files, while also avoiding a train of `--build-arg` in the shell.