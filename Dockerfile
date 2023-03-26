# TODO: For now use preview version of .NET 8.0 SDK and runtime
# TODO: .NET 7.0 has issues with fetching nuget packages

FROM mcr.microsoft.com/dotnet/sdk:8.0-preview-jammy AS build-env
WORKDIR /songshizz_bot

# Copy everything
COPY . ./
# Restore as distinct layers
RUN dotnet restore
# Build and publish a release
RUN dotnet publish -c Release -o out

# Build runtime image
FROM mcr.microsoft.com/dotnet/aspnet:8.0-preview-jammy
WORKDIR /songshizz_bot
COPY --from=build-env /songshizz_bot/out .
ENTRYPOINT ["dotnet", "songshizz_bot.dll"]