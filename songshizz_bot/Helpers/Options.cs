using System;
using System.IO;
using Newtonsoft.Json;
using JsonSerializer = System.Text.Json.JsonSerializer;

namespace SongshizzBot.Helpers
{
    public class Options
    {
        public string BotToken { get; set; }
        public string SpotifyToken { get; set; }
        public ulong GuildId { get; set; }

        public static Options LoadConfig()
        {
            if (File.Exists(Environment.CurrentDirectory + "/options.json"))
                return JsonSerializer.Deserialize<Options>(File.ReadAllText(Environment.CurrentDirectory + "/options.json"));

            Console.WriteLine("Missing config file! Please restart the bot and fill the `botToken` with your token!");
            File.WriteAllText(Environment.CurrentDirectory + "/options.json", JsonConvert.SerializeObject(new Options()
            {
                BotToken = "TOKEN_HERE!",
                SpotifyToken = "SPOTIFY_CLIENT_SECRET_HERE!",
                GuildId = 0
            }, Formatting.Indented));
            Environment.Exit(-404);
            return null;
        }
    }
}