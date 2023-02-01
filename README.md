# gin-juice

A simple way to enhance the Gin and Docker development process. 

## Description

This project faciliates getting a baseline Gin and Docker web application fleshed out for you to give Gin a more "Django" feel.

## Getting Started

### Dependencies

* Docker - Any modern version should be fine. See 
https://docs.docker.com/get-docker/
* Git bash or WSL as a terminal to leverage if using Windows. 
* Bash if using Linux/OSX

### Installing
* Ensure you are in an empty directory where you are ready to start work.
* Copy the below string and replace <GIN_PROJECT> with your own project name.
```
curl -L -o tmp.zip https://github.com/anthony-hopkins/gin-juice/archive/refs/heads/main.zip && \
unzip tmp.zip && \
cp -r ./gin-juice-main/* . && \
bash ./sip.sh <GIN_PROJECT>
```

## What now?
You can run the follow command to make sure your environment is properly set up. It will provide you a simple hello gin page:
```
docker-compose up
```
Just navigate to http://localhost:80 and you should see your gin app running within it's container.

Simply make changes to your web application and rerun your docker-compose 

Now all that is left to do is build something cool. Have fun!

## Authors

Contributors names and contact info

Anthony Hopkins - anthony.b.hopkins@gmail.com

## Acknowledgments

Shout out to these guys for basically providing a sweet README.md template:
* [awesome-readme](https://github.com/matiassingers/awesome-readme)
