using System;
using System.Linq;
using System.Threading.Tasks;
using DSharpPlus;
using DSharpPlus.Entities;
using DSharpPlus.EventArgs;
using Html2Markdown;
using songshizz_bot;
using Songwhip;
using Songwhip.Models;
using SpotifyAPI.Web;

namespace SongshizzBot
{
    public class Songshizz
    {
        public static async Task ProcessMessage(MessageCreateEventArgs e, DiscordClient discord)
        {
            if (e.Author.Id == discord.CurrentUser.Id)
                return;
            
            string link = Utilities.ExtractLink(e.Message.Content);

            if (!Utilities.IsLinkMessage(link))
                return;

            if (Utilities.IsBlackListed(e))
                return;

            if (!Utilities.IsYouTubeLink(link))
                await e.Message.DeleteAsync();

            if (Utilities.IsSpotifyPlaylist(link))
            {
                string playlistId = link.Split("/playlist/")[1].Split('?')[0];
                SpotifyClient spotify = new SpotifyClient(Program.Options.SpotifyToken);

                try
                {
                    FullPlaylist playlist = await spotify.Playlists.Get(playlistId);
                    
                    if (!await PostEmbed(playlist, e, discord))
                        await PostFailEmbed(e, discord);
                }
                catch (APIException)
                {
                    await PostFailEmbed(e, discord);
                }
            }
            else
            {
                Tuple<SongwhipInfo, DiscordMessage> result = await FetchSongwhipInfo(e, discord);

                if (!await PostEmbed(result.Item1, e, discord))
                    await PostFailEmbed(e, discord);
                
                if (!Utilities.IsYouTubeLink(link))
                    await result.Item2.DeleteAsync();
            }
        }

        private static async Task PostFailEmbed(MessageCreateEventArgs e, DiscordClient discord)
        {
            string link = Utilities.ExtractLink(e.Message.Content);
            if (Utilities.IsYouTubeLink(link))
                return;
            
            DiscordEmbed embed = new DiscordEmbedBuilder
                {
                    Color = DiscordColor.DarkBlue,
                    Title = $"{e.Author.Username} shared a music link ...",
                    Description = $"We couldn't find anything with this link on Songwhip but maybe it's too new or {e.Author.Username} might have an unconventional music style?\n\n🔗 [Original music link]({link})"
                }
                .WithFooter($"Shared by {e.Author.Username}", e.Author.AvatarUrl)
                .Build();
            
            await discord.SendMessageAsync(e.Message.Channel, embed);
        }

        private static async Task<bool> PostEmbed(SongwhipInfo info, MessageCreateEventArgs e, DiscordClient discord)
        {
            string link = Utilities.ExtractLink(e.Message.Content);
            if (info == null)
                return false;

            if (Utilities.IsYouTubeLink(link)) // If it is a YouTube link and we got this far it probably means we managed to find a album/artist/song link for it.
                await e.Message.DeleteAsync();

            DiscordEmbedBuilder mainEmbed = new DiscordEmbedBuilder
                {
                    ImageUrl = info.Image,
                    Color = DiscordColor.Purple,
                    Description = BuildDescription(info, link)
                }
                .WithFooter($"Shared by {e.Author.Username}", e.Author.AvatarUrl)
                .WithAuthor($"{string.Join(" ", info.Artists.Select(x => x.Name))} - {info.Name}", info.Url,
                    info.Artists.First().Image);

            await discord.SendMessageAsync(e.Message.Channel, mainEmbed.Build());
            return true;
        }
        
        private static async Task<bool> PostEmbed(FullPlaylist info, MessageCreateEventArgs e, DiscordClient discord)
        {
            string link = Utilities.ExtractLink(e.Message.Content);
            
            if (info == null)
                return false;

            string ownerImage = info.Owner.Images == null ? info.Images.First().Url : info.Owner.Images.First().Url;
            DiscordEmbedBuilder mainEmbed = new DiscordEmbedBuilder
                {
                    ImageUrl = info.Images.First().Url,
                    Color = DiscordColor.Purple,
                    Description = BuildDescription(info, link)
                }
                .WithFooter($"Shared by {e.Author.Username}", e.Author.AvatarUrl)
                .WithAuthor($"Spotify playlist - {info.Name}", info.Href, ownerImage);

            await discord.SendMessageAsync(e.Message.Channel, mainEmbed.Build());
            return true;
        }

        private static async Task<Tuple<SongwhipInfo, DiscordMessage>> FetchSongwhipInfo(MessageCreateEventArgs e, DiscordClient discord)
        {
            Console.WriteLine($"Trying to fetch data for {e.Message.Author.Username}#{e.Message.Author.Id} on {e.Guild.Name} ({e.Guild.Id})");
            DiscordEmbed embed = new DiscordEmbedBuilder
            {
                Color = DiscordColor.Blurple,
                Title = "Detected a music link",
                Description = "Just a moment while we try to resolve this music link for you ... "
            }.Build();
            
            string link = Utilities.ExtractLink(e.Message.Content);
            DiscordMessage msg = null;
            
            if (!Utilities.IsYouTubeLink(link))
                msg = await discord.SendMessageAsync(e.Message.Channel, embed);
            
            return new Tuple<SongwhipInfo, DiscordMessage>(await SongwhipApi.GetInfo(link), msg);
        }
        
        private static string BuildDescription(SongwhipInfo info, string originalMessage)
        {
            string link = Utilities.ExtractLink(originalMessage);
            string desc = "";
            desc =
                $"**Release date:** {info.ReleaseDate:dd-MM-yyyy}\n**Track name:** {info.Name}\n**Artist:** {string.Join(" ", info.Artists.Select(x => x.Name))}\nListen on ";

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

            desc += $"\n\n🔗 [Streaming services]({info.Url})";
            desc += $"\n🔗 [Original link]({link})\n";
            return desc;
        }
        
        private static string BuildDescription(FullPlaylist info, string originalMessage)
        {
            string link = Utilities.ExtractLink(originalMessage);
            string desc = $"";

            desc += $"\n{new Converter().Convert(info.Description)}\n";
            desc += $"\n** Tracks: ** {info.Tracks.Total}";
            desc += $"\n** Creator: ** {info.Owner.DisplayName}";
            
            desc += $"\n\n🔗 [Open playlist]({link})\n";
            return desc;
        }
    }
}