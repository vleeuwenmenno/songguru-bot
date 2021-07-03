using System;
using System.Linq;
using System.Threading.Tasks;
using DSharpPlus;
using DSharpPlus.Entities;
using DSharpPlus.SlashCommands;
using Newtonsoft.Json;
using RestSharp;

namespace DiscordCodBot
{
    public class SongwhipCommands : SlashCommandModule
    {
        public Options _options { get; set; }
        
        [SlashCommand("songwhip", "Convert a song link to a Songwhip link (Spotify, Deezer, YouTube and more")]
        public async Task Songwhip(InteractionContext ctx, [Option("song_link", "The song link from a streaming service to convert to a Songwhip link.")] string song_link)
        {
            Console.WriteLine($"/songwhip requested by {ctx.Member.Username}#{ctx.Member.Id} on {ctx.Guild.Name} ({ctx.Guild.Id}):");
            
            await ctx.CreateResponseAsync(InteractionResponseType.DeferredChannelMessageWithSource);
            string desc = "";
            DiscordUser user = ctx.Member;
            SongwhipInfo info = GetSongwhip(song_link);

            if (info == null)
            {
                await ctx.EditResponseAsync(new DiscordWebhookBuilder().AddEmbed(new DiscordEmbedBuilder
                {
                    Color = DiscordColor.DarkRed,
                    Title = "Not found~",
                    Description = "We couldn't find anything with the provided link."
                }.Build()));
                Console.WriteLine($"   Request fulfilled with 404 :(");
                return;
            }
                
            DiscordWebhookBuilder builder = new DiscordWebhookBuilder();

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
            
            var mainEmbed = new DiscordEmbedBuilder
                {
                    ImageUrl = info.Image,
                    Color = DiscordColor.Purple,
                    Description = desc
                }.
                WithFooter($"Shared by {ctx.Member.DisplayName}", ctx.Member.AvatarUrl).
                WithAuthor($"{string.Join(" ", info.Artists.Select(x => x.Name))} - {info.Name}", info.Url, info.Artists.First().Image);

            builder.AddEmbed(mainEmbed.Build());
            await ctx.EditResponseAsync(builder);
            Console.WriteLine($"   Request fulfilled with data! :D");
        }

        private SongwhipInfo GetSongwhip(string link)
        {
            _options = Options.LoadConfig();

            try
            {
                var client = new RestClient(_options.songwhipEndpoint);
                var request = new RestRequest(Method.POST);
                request.AddHeader("Content-Type", "application/json");
                request.AddParameter("application/json", $"{{\n\t\"url\":\"{link}\"\n}}", ParameterType.RequestBody);
                IRestResponse response = client.Execute(request);
                return JsonConvert.DeserializeObject<SongwhipInfo>(response.Content);
            }
            catch (Exception)
            {
                return null;
            }
        }
    }
}