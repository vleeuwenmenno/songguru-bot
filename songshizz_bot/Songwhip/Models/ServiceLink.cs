using System.Text.Json.Serialization;

namespace Songwhip.Models;

public class ServiceLink
{
    [JsonPropertyName("link")]
    public string Link { get; set; }

    [JsonPropertyName("countries")]
    public object Countries { get; set; }
}