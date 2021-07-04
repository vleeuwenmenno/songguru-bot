using System;
using System.Linq;
using System.Threading.Tasks;
using DSharpPlus;
using DSharpPlus.Entities;
using DSharpPlus.SlashCommands;
using Newtonsoft.Json;
using RestSharp;

namespace songwhip_bot
{
    public class SongwhipCommands : SlashCommandModule
    {
        [SlashCommand("info", "About this bot, with thanks to Wilson!")]
        public async Task Info(InteractionContext ctx)
        {
            await ctx.CreateResponseAsync(InteractionResponseType.ChannelMessageWithSource, new DiscordInteractionResponseBuilder()
                .AddEmbed(new DiscordEmbedBuilder
                    {
                        Color = DiscordColor.Purple,
                        Description = "Bot made by [Menno van Leeuwen](https://github.com/vleeuwenmenno)\nSongwhip made by [Wilson](https://songwhip.com/faq)\n\nI would like to thank Wilson especially for making his API publicly available for everyone to use!\n\n[Add bot to your server](https://discord.com/api/oauth2/authorize?client_id=860899901020700684&permissions=2147837952&scope=bot%20applications.commands) - [GitHub repository](https://github.com/vleeuwenmenno/songwhip-bot)",
                        Title = $"Songwhip bot {Utilities.Version}"
                    }.
                    WithFooter($"About requested by {ctx.Member.DisplayName}", ctx.Member.AvatarUrl)
                    .Build())
            );
        }
        
        [SlashCommand("opt-in", "Enable auto-resolving music links for yourself. Songwhip bot will try resolve your music links.")]
        public async Task OptIn(InteractionContext ctx)
        {
            Blacklist.userBlacklist.RemoveAll(x => x == ctx.Member.Id);
            Blacklist.Save();
            
            await ctx.CreateResponseAsync(InteractionResponseType.ChannelMessageWithSource, new DiscordInteractionResponseBuilder()
                .AddEmbed(new DiscordEmbedBuilder
                    {
                        Color = DiscordColor.DarkGreen,
                        Description = "You've been opted in of auto-resolving music links. <a:Happy:395314894364213248>",
                        Title = $"Opted in~"
                    }.
                    WithFooter($"Opt-in requested by {ctx.Member.DisplayName}", ctx.Member.AvatarUrl)
                    .Build())
            );
        }
        
        [SlashCommand("opt-out", "Disable auto-resolving music links for yourself. Songwhip bot won't bother you ever again!")]
        public async Task OptOut(InteractionContext ctx)
        {
            Blacklist.userBlacklist.Add(ctx.Member.Id);
            Blacklist.Save();
            
            await ctx.CreateResponseAsync(InteractionResponseType.ChannelMessageWithSource, new DiscordInteractionResponseBuilder()
                .AddEmbed(new DiscordEmbedBuilder
                    {
                        Color = DiscordColor.DarkRed,
                        Description = "You've been opted out of auto-resolving music links. <a:Sad:395314892485296148>",
                        Title = $"Opted out~"
                    }.
                    WithFooter($"Opt-out requested by {ctx.Member.DisplayName}", ctx.Member.AvatarUrl)
                    .Build())
            );
        }
        
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

        public static SongwhipInfo GetSongwhip(string link)
        {
            Options _options = Options.LoadConfig();

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