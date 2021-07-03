using System;
using System.IO;
using Newtonsoft.Json;
using JsonSerializer = System.Text.Json.JsonSerializer;

namespace songwhip_bot
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

            Console.WriteLine("Missing config file! Please restart the bot and fill the `botToken` with your token!");
            File.WriteAllText(Environment.CurrentDirectory + "/options.json", JsonConvert.SerializeObject(new Options()
            {
                songwhipEndpoint = "https://songwhip.com/",
                botToken = "TOKEN_HERE!"
            }, Formatting.Indented));
            Environment.Exit(-404);
            return null;
        }
    }
}