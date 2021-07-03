#!/bin/bash
dotnet publish -c Release --runtime linux-x64 -p:PublishSingleFile=true --self-contained true -o build/
