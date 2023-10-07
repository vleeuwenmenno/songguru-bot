# Songwhip Discord bot

A songwhip discord bot that converts your song links to a nice looking embed.
#### [Add this Bot to my Discord server!](https://discord.ly/songwhip) - [Direct](https://discord.com/api/oauth2/authorize?client_id=860899901020700684&permissions=689610975232&scope=bot)

![alt text](https://github.com/vleeuwenmenno/songwhip-bot/raw/develop/images/preview.png)


## Deployment

### How to
Deployment is done using docker and can be done using the provided `docker-compose.yml` file.

1. Create a new folder and create the `docker-compose.yml` file.
2. Update the `docker-compose.yml` with the correct mount point for the configs folder.
3. Create the `configs/config.yaml` with the correct settings.
4. Run the container using `docker-compose up -d`
5. Enjoy~ `docker stats <container>`

Alternatively you can build a container yourself as follows:

1. Clone the repository `git clone git@github.com:vleeuwenmenno/songwhip-bot.git`
2. Checkout to your desired branch `git checkout production`
3. Build a docker image `make`
4. Update your config files and your `docker-compose.yml` to run your local image.
5. Run the docker image `docker compose up -d`
6. Enjoy~ `docker stats <container>`

## Contribute

I am happy to accept contributions, fork the project make your changes and create a PR. Whenver I have time I will review it and merge it ;)

## Licensing

This project is licensed under the MIT License.

```
Copyright 2023 Menno van Leeuwen

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```