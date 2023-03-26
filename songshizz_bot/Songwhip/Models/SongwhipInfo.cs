using System;
using System.Collections.Generic;
using System.Text.Json.Serialization;

namespace Songwhip.Models;

public class SongwhipInfo
{
    [JsonPropertyName("type")]
    public string Type { get; set; }

    [JsonPropertyName("id")]
    public int Id { get; set; }

    [JsonPropertyName("path")]
    public string Path { get; set; }

    [JsonPropertyName("name")]
    public string Name { get; set; }

    [JsonPropertyName("url")]
    public string Url { get; set; }

    [JsonPropertyName("sourceUrl")]
    public string SourceUrl { get; set; }

    [JsonPropertyName("sourceCountry")]
    public string SourceCountry { get; set; }

    [JsonPropertyName("image")]
    public string Image { get; set; }

    [JsonPropertyName("links")]
    public HasLinks Links { get; set; }

    [JsonPropertyName("linksCountries")]
    public List<string> LinksCountries { get; set; }

    [JsonPropertyName("artists")]
    public List<Artist> Artists { get; set; }

    [JsonPropertyName("overrides")]
    public object Overrides { get; set; }
}