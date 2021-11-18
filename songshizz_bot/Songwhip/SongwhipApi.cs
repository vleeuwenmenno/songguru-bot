using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Threading.Tasks;
using Songwhip.Models;
using System.Text.Json;

namespace Songwhip
{
    public class SongwhipApi
    {
        public static async Task<SongwhipInfo> GetInfo(string link)
        {
            var client = new HttpClient();
            var request = new HttpRequestMessage
            {
                Method = HttpMethod.Post,
                RequestUri = new Uri("https://songwhip.com/"),
                Content = new StringContent(JsonSerializer.Serialize(new Dictionary<string, string>() { { "url", link } }))
                {
                    Headers =
                    {
                        ContentType = new MediaTypeHeaderValue("application/json")
                    }
                }
            };

            using var response = await client.SendAsync(request);
            {
                if (!response.IsSuccessStatusCode) return null;
                
                var body = await response.Content.ReadAsStringAsync();
                return JsonSerializer.Deserialize<SongwhipInfo>(body);
            }
        }
    }
}