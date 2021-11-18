using System.Collections.Generic;
using System.Text.Json;
using System.Text.Json.Serialization;

namespace Songwhip.Models
{
    public class Links
    {
        [JsonPropertyName("tidal")]
        public List<ServiceLink> Tidal { get; set; }

        [JsonPropertyName("amazon")]
        public List<ServiceLink> Amazon { get; set; }

        [JsonPropertyName("deezer")]
        public List<ServiceLink> Deezer { get; set; }

        [JsonPropertyName("itunes")]
        public List<ServiceLink> Itunes { get; set; }

        [JsonPropertyName("napster")]
        public List<ServiceLink> Napster { get; set; }

        [JsonPropertyName("pandora")]
        public List<ServiceLink> Pandora { get; set; }

        [JsonPropertyName("spotify")]
        public List<ServiceLink> Spotify { get; set; }

        [JsonPropertyName("youtube")]
        public List<ServiceLink> Youtube { get; set; }

        [JsonPropertyName("amazonMusic")]
        public List<ServiceLink> AmazonMusic { get; set; }

        [JsonPropertyName("itunesStore")]
        public List<ServiceLink> ItunesStore { get; set; }

        [JsonPropertyName("youtubeMusic")]
        public List<ServiceLink> YoutubeMusic { get; set; }

        [JsonPropertyName("yandex")]
        public List<ServiceLink> Yandex { get; set; }

        [JsonPropertyName("discogs")]
        public List<ServiceLink> Discogs { get; set; }

        [JsonPropertyName("twitter")]
        public List<ServiceLink> Twitter { get; set; }

        [JsonPropertyName("facebook")]
        public List<ServiceLink> Facebook { get; set; }

        [JsonPropertyName("instagram")]
        public List<ServiceLink> Instagram { get; set; }

        [JsonPropertyName("wikipedia")]
        public List<ServiceLink> Wikipedia { get; set; }

        [JsonPropertyName("soundcloud")]
        public List<ServiceLink> Soundcloud { get; set; }

        [JsonPropertyName("musicBrainz")]
        public List<ServiceLink> MusicBrainz { get; set; }
    }
}