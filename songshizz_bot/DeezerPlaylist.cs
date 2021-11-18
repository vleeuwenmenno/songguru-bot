using System;
using System.Net.Http;
using System.Threading.Tasks;

namespace SongshizzBot;

public class DeezerPlaylist
{
    public string imageUrl { get; set; }
    public string title { get; set; }
    public string creator { get; set; }
    public string description { get; set; }
        
    public static async Task<DeezerPlaylist> ScrapeInfo(string playlistId)
    {
        var clientHandler = new HttpClientHandler
        {
            UseCookies = false,
        };
        var client = new HttpClient(clientHandler);
        var request = new HttpRequestMessage
        {
            Method = HttpMethod.Get,
            RequestUri = new Uri($"https://www.deezer.com/en/playlist/{playlistId}"),
        };
        using (var response = await client.SendAsync(request))
        {
            response.EnsureSuccessStatusCode();
                
            var body = await response.Content.ReadAsStringAsync();
            
            DeezerPlaylist info = new DeezerPlaylist();
            string desc = body.Split("<div id=\"naboo_playlist_description\" class=\"quote\" itemprop=\"description\">")[1] .Split("</div>")[0].Trim();
            
            info.imageUrl = body.Split("<img id=\"naboo_playlist_image\" itemprop=\"image\" src=\"")[1].Split("\"")[0].Trim();
            info.title = body.Split("<h1 class=\"heading-1\" id=\"naboo_playlist_title\" itemprop=\"name\">")[1].Split("</h1>")[0].Trim();
            info.creator = body.Split("<span id=\"naboo_playlist_information_creator\">")[1].Split("</span>")[0].Trim().Replace("\t", "").Substring(2);
            info.description = string.IsNullOrEmpty(desc) ? "" : desc.Substring(1, desc.Length - 2);
            
            return info;
        }
    }
}