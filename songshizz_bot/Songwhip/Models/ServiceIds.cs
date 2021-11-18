using System.Text.Json.Serialization;

namespace Songwhip.Models;

public class ServiceIds
{
    [JsonPropertyName("tidal")]
    public string Tidal { get; set; }

    [JsonPropertyName("amazon")]
    public string Amazon { get; set; }

    [JsonPropertyName("deezer")]
    public string Deezer { get; set; }

    [JsonPropertyName("itunes")]
    public string Itunes { get; set; }

    [JsonPropertyName("yandex")]
    public string Yandex { get; set; }

    [JsonPropertyName("discogs")]
    public string Discogs { get; set; }

    [JsonPropertyName("napster")]
    public string Napster { get; set; }

    [JsonPropertyName("pandora")]
    public string Pandora { get; set; }

    [JsonPropertyName("spotify")]
    public string Spotify { get; set; }

    [JsonPropertyName("soundcloud")]
    public string Soundcloud { get; set; }

    [JsonPropertyName("musicBrainz")]
    public string MusicBrainz { get; set; }

    [JsonPropertyName("youtubeMusic")]
    public string YoutubeMusic { get; set; }
}