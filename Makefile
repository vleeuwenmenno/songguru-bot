build:
	docker buildx create --use
	docker buildx build --platform linux/amd64,linux/arm64/v8 -t stardebris/songshizz-bot .
  
run: 
	docker-compose up

latesttag=`git describe --tags --abbrev=0 --always`
.PHONY: publish
publish:	
	docker buildx create --use
	docker buildx build --platform linux/amd64,linux/arm64/v8 -t stardebris/songshizz-bot -t stardebris/songshizz-bot:${latesttag} --push .
