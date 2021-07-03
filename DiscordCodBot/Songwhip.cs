using System;
using System.Collections;
using System.Collections.Generic;
using Newtonsoft.Json;

namespace DiscordCodBot
{
    public class Links
    {
        [JsonProperty("tidal", NullValueHandling = NullValueHandling.Ignore)]
        public List<ServiceLink> Tidal { get; set; }

        [JsonProperty("amazon", NullValueHandling = NullValueHandling.Ignore)]
        public List<ServiceLink> Amazon { get; set; }

        [JsonProperty("deezer", NullValueHandling = NullValueHandling.Ignore)]
        public List<ServiceLink> Deezer { get; set; }

        [JsonProperty("itunes", NullValueHandling = NullValueHandling.Ignore)]
        public List<ServiceLink> Itunes { get; set; }

        [JsonProperty("napster", NullValueHandling = NullValueHandling.Ignore)]
        public List<ServiceLink> Napster { get; set; }

        [JsonProperty("pandora", NullValueHandling = NullValueHandling.Ignore)]
        public List<ServiceLink> Pandora { get; set; }

        [JsonProperty("spotify", NullValueHandling = NullValueHandling.Ignore)]
        public List<ServiceLink> Spotify { get; set; }

        [JsonProperty("youtube", NullValueHandling = NullValueHandling.Ignore)]
        public List<ServiceLink> Youtube { get; set; }

        [JsonProperty("amazonMusic", NullValueHandling = NullValueHandling.Ignore)]
        public List<ServiceLink> AmazonMusic { get; set; }

        [JsonProperty("itunesStore", NullValueHandling = NullValueHandling.Ignore)]
        public List<ServiceLink> ItunesStore { get; set; }

        [JsonProperty("youtubeMusic", NullValueHandling = NullValueHandling.Ignore)]
        public List<ServiceLink> YoutubeMusic { get; set; }

        [JsonProperty("yandex", NullValueHandling = NullValueHandling.Ignore)]
        public List<ServiceLink> Yandex { get; set; }

        [JsonProperty("discogs", NullValueHandling = NullValueHandling.Ignore)]
        public List<ServiceLink> Discogs { get; set; }

        [JsonProperty("twitter", NullValueHandling = NullValueHandling.Ignore)]
        public List<ServiceLink> Twitter { get; set; }

        [JsonProperty("facebook", NullValueHandling = NullValueHandling.Ignore)]
        public List<ServiceLink> Facebook { get; set; }

        [JsonProperty("instagram", NullValueHandling = NullValueHandling.Ignore)]
        public List<ServiceLink> Instagram { get; set; }

        [JsonProperty("wikipedia", NullValueHandling = NullValueHandling.Ignore)]
        public List<ServiceLink> Wikipedia { get; set; }

        [JsonProperty("soundcloud", NullValueHandling = NullValueHandling.Ignore)]
        public List<ServiceLink> Soundcloud { get; set; }

        [JsonProperty("musicBrainz", NullValueHandling = NullValueHandling.Ignore)]
        public List<ServiceLink> MusicBrainz { get; set; }
    }

    public class ServiceLink
    {
        [JsonProperty("link", NullValueHandling = NullValueHandling.Ignore)]
        public string Link { get; set; }

        [JsonProperty("countries", NullValueHandling = NullValueHandling.Ignore)]
        public object Countries { get; set; }
    }

    public class ServiceIds
    {
        [JsonProperty("tidal", NullValueHandling = NullValueHandling.Ignore)]
        public string Tidal { get; set; }

        [JsonProperty("amazon", NullValueHandling = NullValueHandling.Ignore)]
        public string Amazon { get; set; }

        [JsonProperty("deezer", NullValueHandling = NullValueHandling.Ignore)]
        public string Deezer { get; set; }

        [JsonProperty("itunes", NullValueHandling = NullValueHandling.Ignore)]
        public string Itunes { get; set; }

        [JsonProperty("yandex", NullValueHandling = NullValueHandling.Ignore)]
        public string Yandex { get; set; }

        [JsonProperty("discogs", NullValueHandling = NullValueHandling.Ignore)]
        public string Discogs { get; set; }

        [JsonProperty("napster", NullValueHandling = NullValueHandling.Ignore)]
        public string Napster { get; set; }

        [JsonProperty("pandora", NullValueHandling = NullValueHandling.Ignore)]
        public string Pandora { get; set; }

        [JsonProperty("spotify", NullValueHandling = NullValueHandling.Ignore)]
        public string Spotify { get; set; }

        [JsonProperty("soundcloud", NullValueHandling = NullValueHandling.Ignore)]
        public string Soundcloud { get; set; }

        [JsonProperty("musicBrainz", NullValueHandling = NullValueHandling.Ignore)]
        public string MusicBrainz { get; set; }

        [JsonProperty("youtubeMusic", NullValueHandling = NullValueHandling.Ignore)]
        public string YoutubeMusic { get; set; }
    }

    public class Artist
    {
        [JsonProperty("type", NullValueHandling = NullValueHandling.Ignore)]
        public string Type { get; set; }

        [JsonProperty("id", NullValueHandling = NullValueHandling.Ignore)]
        public int Id { get; set; }

        [JsonProperty("path", NullValueHandling = NullValueHandling.Ignore)]
        public string Path { get; set; }

        [JsonProperty("name", NullValueHandling = NullValueHandling.Ignore)]
        public string Name { get; set; }

        [JsonProperty("sourceUrl", NullValueHandling = NullValueHandling.Ignore)]
        public string SourceUrl { get; set; }

        [JsonProperty("sourceCountry", NullValueHandling = NullValueHandling.Ignore)]
        public string SourceCountry { get; set; }

        [JsonProperty("url", NullValueHandling = NullValueHandling.Ignore)]
        public string Url { get; set; }

        [JsonProperty("image", NullValueHandling = NullValueHandling.Ignore)]
        public string Image { get; set; }

        [JsonProperty("createdAt", NullValueHandling = NullValueHandling.Ignore)]
        public DateTime CreatedAt { get; set; }

        [JsonProperty("updatedAt", NullValueHandling = NullValueHandling.Ignore)]
        public DateTime UpdatedAt { get; set; }

        [JsonProperty("refreshedAt", NullValueHandling = NullValueHandling.Ignore)]
        public DateTime RefreshedAt { get; set; }

        [JsonProperty("linksCountries", NullValueHandling = NullValueHandling.Ignore)]
        public List<string> LinksCountries { get; set; }

        [JsonProperty("links", NullValueHandling = NullValueHandling.Ignore)]
        public Links Links { get; set; }

        [JsonProperty("description", NullValueHandling = NullValueHandling.Ignore)]
        public string Description { get; set; }

        [JsonProperty("serviceIds", NullValueHandling = NullValueHandling.Ignore)]
        public ServiceIds ServiceIds { get; set; }
    }

    public class SongwhipInfo
    {
        [JsonProperty("type", NullValueHandling = NullValueHandling.Ignore)]
        public string Type { get; set; }

        [JsonProperty("id", NullValueHandling = NullValueHandling.Ignore)]
        public int Id { get; set; }

        [JsonProperty("path", NullValueHandling = NullValueHandling.Ignore)]
        public string Path { get; set; }

        [JsonProperty("name", NullValueHandling = NullValueHandling.Ignore)]
        public string Name { get; set; }

        [JsonProperty("url", NullValueHandling = NullValueHandling.Ignore)]
        public string Url { get; set; }

        [JsonProperty("sourceUrl", NullValueHandling = NullValueHandling.Ignore)]
        public string SourceUrl { get; set; }

        [JsonProperty("sourceCountry", NullValueHandling = NullValueHandling.Ignore)]
        public string SourceCountry { get; set; }

        [JsonProperty("releaseDate", NullValueHandling = NullValueHandling.Ignore)]
        public DateTime ReleaseDate { get; set; }

        [JsonProperty("createdAt", NullValueHandling = NullValueHandling.Ignore)]
        public DateTime CreatedAt { get; set; }

        [JsonProperty("updatedAt", NullValueHandling = NullValueHandling.Ignore)]
        public DateTime UpdatedAt { get; set; }

        [JsonProperty("refreshedAt", NullValueHandling = NullValueHandling.Ignore)]
        public DateTime RefreshedAt { get; set; }

        [JsonProperty("image", NullValueHandling = NullValueHandling.Ignore)]
        public string Image { get; set; }

        [JsonProperty("config", NullValueHandling = NullValueHandling.Ignore)]
        public object Config { get; set; }

        [JsonProperty("links", NullValueHandling = NullValueHandling.Ignore)]
        public HasLinks Links { get; set; }

        [JsonProperty("linksCountries", NullValueHandling = NullValueHandling.Ignore)]
        public List<string> LinksCountries { get; set; }

        [JsonProperty("artists", NullValueHandling = NullValueHandling.Ignore)]
        public List<Artist> Artists { get; set; }

        [JsonProperty("overrides", NullValueHandling = NullValueHandling.Ignore)]
        public object Overrides { get; set; }
    }

    public class HasLinks
    {
        [JsonProperty("tidal", NullValueHandling = NullValueHandling.Ignore)]
        public bool Tidal { get; set; }

        [JsonProperty("amazon", NullValueHandling = NullValueHandling.Ignore)]
        public bool Amazon { get; set; }

        [JsonProperty("deezer", NullValueHandling = NullValueHandling.Ignore)]
        public bool Deezer { get; set; }

        [JsonProperty("itunes", NullValueHandling = NullValueHandling.Ignore)]
        public bool Itunes { get; set; }

        [JsonProperty("napster", NullValueHandling = NullValueHandling.Ignore)]
        public bool Napster { get; set; }

        [JsonProperty("pandora", NullValueHandling = NullValueHandling.Ignore)]
        public bool Pandora { get; set; }

        [JsonProperty("spotify", NullValueHandling = NullValueHandling.Ignore)]
        public bool Spotify { get; set; }

        [JsonProperty("youtube", NullValueHandling = NullValueHandling.Ignore)]
        public bool Youtube { get; set; }

        [JsonProperty("amazonMusic", NullValueHandling = NullValueHandling.Ignore)]
        public bool AmazonMusic { get; set; }

        [JsonProperty("itunesStore", NullValueHandling = NullValueHandling.Ignore)]
        public bool ItunesStore { get; set; }

        [JsonProperty("youtubeMusic", NullValueHandling = NullValueHandling.Ignore)]
        public bool YoutubeMusic { get; set; }
    }
}