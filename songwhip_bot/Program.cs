using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using System.Threading.Tasks;
using DSharpPlus;
using DSharpPlus.Entities;
using DSharpPlus.EventArgs;
using DSharpPlus.SlashCommands;
using Microsoft.Extensions.DependencyInjection;
using Newtonsoft.Json;

namespace songwhip_bot
{
    // Invite the bot with:
    // DEVELOPER BOT: https://discord.com/oauth2/authorize?client_id=861018175935873054&permissions=2147837952&scope=bot%20applications.commands
    // LIVE BOT: https://discord.com/api/oauth2/authorize?client_id=860899901020700684&permissions=2147870784&scope=bot%20applications.commands
    public class Program
    {
        public static Options _options;
        public static DiscordClient _discord;
        public SlashCommandsExtension _slash;

        static void OnProcessExit (object sender, EventArgs e)
        {
            _discord.DisconnectAsync();
            Console.WriteLine($"Songwhip bot {Utilities.Version} closing shop ~");
        }
        
        static void Main(string[] args) =>
            new Program().MainAsync().GetAwaiter().GetResult();
        
        public async Task MainAsync()
        {
            AppDomain.CurrentDomain.ProcessExit += new EventHandler (OnProcessExit); 
            Console.WriteLine($"Songwhip bot {Utilities.Version} starting ...");
            
            _options = Options.LoadConfig();

            if (File.Exists($"{Environment.CurrentDirectory}/users-blacklist.json"))
                Blacklist.userBlacklist = JsonConvert.DeserializeObject<List<ulong>>(File.ReadAllText($"{Environment.CurrentDirectory}/users-blacklist.json"));
         
            if (File.Exists($"{Environment.CurrentDirectory}/guild-blacklist.json"))
                Blacklist.guildBlacklist = JsonConvert.DeserializeObject<List<ulong>>(File.ReadAllText($"{Environment.CurrentDirectory}/guild-blacklist.json"));
            
            if (Blacklist.userBlacklist == null)
                Blacklist.userBlacklist = new List<ulong>();

            if (Blacklist.guildBlacklist == null)
                Blacklist.guildBlacklist = new List<ulong>();
            
            _discord = new DiscordClient(new DiscordConfiguration()
            {
                Token = _options.botToken,
                TokenType = TokenType.Bot,
                Intents = DiscordIntents.AllUnprivileged
            });
            
            _discord.MessageCreated += MessageCreated;
            _discord.Ready += DiscordOnReady;
            
            _slash = _discord.UseSlashCommands(new SlashCommandsConfiguration
            {
                Services = new ServiceCollection()
                    .AddSingleton<Options>(_options)
                    .BuildServiceProvider()
            });
            
            _slash.RegisterCommands<SongwhipCommands>();
            
            await _discord.ConnectAsync();
            Console.WriteLine("Songwhip bot is ready!");
            await Task.Delay(-1);
        }

        private async Task MessageCreated(DiscordClient sender, MessageCreateEventArgs e)
        {
            if (e.Author.Id == _discord.CurrentUser.Id)
                return;
            
            string msg = e.Message.Content;
            if (msg.StartsWith("https://www.spotify.com/") ||
                msg.StartsWith("https://spotify.com/") ||
                msg.StartsWith("https://open.spotify.com/") ||
                msg.StartsWith("https://deezer.com/") ||
                msg.StartsWith("https://www.deezer.com/") ||
                msg.StartsWith("https://youtube.com/") ||
                msg.StartsWith("https://www.youtube.com/") ||
                msg.StartsWith("https://music.youtube.com/") ||
                msg.StartsWith("https://music.apple.com/") ||
                msg.StartsWith("https://tidal.com/") ||
                msg.StartsWith("https://www.tidal.com/"))
            {
                if (Blacklist.userBlacklist.Any(x => x == e.Message.Author.Id))
                {                
                    Console.WriteLine($"music link detected for {e.Message.Author.Username}#{e.Message.Author.Id} on {e.Guild.Name} ({e.Guild.Id}) but doing nothing because user opted-out~");
                    return;
                }

                DiscordEmbed embed = null;
                DiscordMessage loadingMsg = null;
                if (!msg.StartsWith("https://youtube.com/") &&
                    !msg.StartsWith("https://www.youtube.com/"))
                {
                    Console.WriteLine($"music link detected for {e.Message.Author.Username}#{e.Message.Author.Id} on {e.Guild.Name} ({e.Guild.Id})");
                    await e.Message.DeleteAsync();
                    embed = new DiscordEmbedBuilder
                    {
                        Color = DiscordColor.Blurple,
                        Title = "Detected a music link",
                        Description = "Just a moment while we try to resolve this music link for you ... "
                    }.Build();
                    loadingMsg = await _discord.SendMessageAsync(e.Message.Channel, embed);
                }

                string desc = "";
                SongwhipInfo info = SongwhipCommands.GetSongwhip(msg);

                if (!msg.StartsWith("https://youtube.com/") &&
                    !msg.StartsWith("https://www.youtube.com/") && 
                    info == null)
                {
                    embed = new DiscordEmbedBuilder
                    {
                        Color = DiscordColor.DarkRed,
                        Title = "Not found~",
                        Description = $"We couldn't find anything with this music link, maybe it's too new or you might have a unconventional music style? :upsidedown:\n[Original link]({msg})"
                    }.Build();
                    
                    Console.WriteLine($"   Request fulfilled with 404 :(");
                    await loadingMsg.DeleteAsync();
                    await _discord.SendMessageAsync(e.Message.Channel, embed);
                    return;
                }
                else if (info == null)
                    return;

                desc = $"**Release date:** {info.ReleaseDate.ToString("dd-MM-yyyy")}\n**Track name:** {info.Name}\n**Artist:** {string.Join(" ", info.Artists.Select(x => x.Name))}\nListen on ";
                
                if (info.Links.Spotify)
                    desc += "<:spotify:860992370954469407> ";
                
                if (info.Links.Deezer)
                    desc += "<:deezer:860992333914570772> ";
                
                if (info.Links.Itunes)
                    desc += "<:applemusic:860995200797507604> ";
                
                if (info.Links.YoutubeMusic)
                    desc += "<:youtubemusic:860994648888836118> ";
                
                if (info.Links.Youtube)
                    desc += "<:youtube:860992285483335730> ";
                
                if (info.Links.Pandora)
                    desc += "<:pandora:860992558519418910> ";
                
                if (info.Links.Tidal)
                    desc += "<:tidal:860992188434612245> ";
                
                desc += $"\n[Streaming services]({info.Url})\n";
                
                DiscordEmbedBuilder mainEmbed = new DiscordEmbedBuilder
                    {
                        ImageUrl = info.Image,
                        Color = DiscordColor.Purple,
                        Description = desc
                    }.
                    WithFooter($"Shared by {e.Author.Username}", e.Author.AvatarUrl).
                    WithAuthor($"{string.Join(" ", info.Artists.Select(x => x.Name))} - {info.Name}", info.Url, info.Artists.First().Image);

                if (!msg.StartsWith("https://youtube.com/") &&
                    !msg.StartsWith("https://www.youtube.com/"))
                    await loadingMsg.DeleteAsync();
                
                else
                    await e.Message.DeleteAsync();
                

                await _discord.SendMessageAsync(e.Message.Channel, mainEmbed.Build());
                Console.WriteLine($"   Request fulfilled with data! :D");
            }
        }

        private Task DiscordOnReady(DiscordClient sender, ReadyEventArgs e)
        {
            return Task.CompletedTask;
        }
    }
}