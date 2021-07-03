using System;
using System.IO;
using Newtonsoft.Json;
using JsonSerializer = System.Text.Json.JsonSerializer;

namespace DiscordCodBot
{
    public class Options
    {
        public string botToken { get; set; }
        public string songwhipEndpoint { get; set; }
        public ulong guildId { get; set; }

        public static Options LoadConfig()
        {
            if (File.Exists(Environment.CurrentDirectory + "/options.json"))
                return JsonSerializer.Deserialize<Options>(File.ReadAllText(Environment.CurrentDirectory + "/options.json"));
            
            return new Options();
        }
    }
}