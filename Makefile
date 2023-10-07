build:
	docker buildx build --load --platform linux/amd64 -t stardebris/songshizz-bot .
  
run: 
	docker-compose up

latesttag=`git describe --tags --abbrev=0 --always`
.PHONY: publish
publish:
	docker buildx build --platform linux/amd64 -t stardebris/songshizz-bot -t stardebris/songshizz-bot:${latesttag} --push .