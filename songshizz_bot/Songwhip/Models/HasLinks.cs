using System.Text.Json.Serialization;

namespace Songwhip.Models;

public class HasLinks
{
    [JsonPropertyName("tidal")]
    public bool Tidal { get; set; }

    [JsonPropertyName("amazon")]
    public bool Amazon { get; set; }

    [JsonPropertyName("deezer")]
    public bool Deezer { get; set; }

    [JsonPropertyName("itunes")]
    public bool Itunes { get; set; }

    [JsonPropertyName("napster")]
    public bool Napster { get; set; }

    [JsonPropertyName("pandora")]
    public bool Pandora { get; set; }

    [JsonPropertyName("spotify")]
    public bool Spotify { get; set; }

    [JsonPropertyName("youtube")]
    public bool Youtube { get; set; }

    [JsonPropertyName("amazonMusic")]
    public bool AmazonMusic { get; set; }

    [JsonPropertyName("itunesStore")]
    public bool ItunesStore { get; set; }

    [JsonPropertyName("youtubeMusic")]
    public bool YoutubeMusic { get; set; }
}