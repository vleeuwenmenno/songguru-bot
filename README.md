# SongGuru Discord bot

A discord bot that converts your song links to a universal format. No longer limited to sharing music with friends that use the same streaming service!
#### [Add this Bot to my Discord server!](https://discord.ly/songguru) - [Direct](https://discord.com/api/oauth2/authorize?client_id=860899901020700684&permissions=689610975232&scope=bot)

![alt text](https://github.com/vleeuwenmenno/songguru-bot/raw/develop/images/preview.png)


## Deployment

### How to
Deployment is done using docker and can be done using the provided `docker-compose.yml` file.

1. Create a new folder and create the `docker-compose.yml` file.
2. Update the `docker-compose.yml` with the correct mount point for the configs folder.
3. Create the `configs/config.yaml` with the correct settings.
4. Run the container using `docker-compose up -d`
5. Enjoy~ `docker stats <container>`

Alternatively you can build a container yourself as follows:

1. Clone the repository `git clone git@github.com:vleeuwenmenno/songguru-bot.git`
2. Checkout to your desired branch `git checkout production`
3. Build a docker image `make`
4. Update your config files and your `docker-compose.yml` to run your local image.
5. Run the docker image `docker compose up -d`
6. Enjoy~ `docker stats <container>`

## FAQ

1. How do I change my server/member preferences?
   - In the guild/server you want to do this, use the command `/settings`. This will generate a one-time link where you can edit your settings and your guild/server settings if you have the proper role for that.

2. I already had the bot on my server, but I can't edit the admin settings?
   - Due to the v2.0.0 update, the bot adds new features that require the "manage roles" permission to work properly. You may need to invite the bot to your server again. The invite link provided in this README has been updated to include the proper permissions. Alternatively, you can create a new role called `SongGuruAdmin`, which should also work. This needs to be verified (#28).

3. I shared the settings link, and now all my settings have changed?!
   - Please avoid sharing this link. You can generate a new one, and all links are only valid for 15 minutes. Although no harm can be done with the link, it's always better to keep it private. In the future, we might change the way these links work to prevent multiple sessions by locking them to your IP or using another method.

4. I tried changing the admin settings, but even after adding the admin role to my Discord user on the guild, I still can't change admin settings?
   - Make sure you have added the proper role and generated a new settings link after doing so. The link won't detect role changes once it has been generated, so you need to generate a new one.

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