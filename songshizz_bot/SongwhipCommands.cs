using System.Linq;
using System.Threading.Tasks;
using DSharpPlus;
using DSharpPlus.Entities;
using DSharpPlus.SlashCommands;
using SongshizzBot.Helpers;

namespace SongshizzBot
{
    public class SongshizzCommands : ApplicationCommandModule
    {
        [SlashCommand("info", "About this bot ...")]
        public async Task Info(InteractionContext ctx)
        {
            await ctx.CreateResponseAsync(InteractionResponseType.ChannelMessageWithSource, new DiscordInteractionResponseBuilder()
                .AddEmbed(new DiscordEmbedBuilder
                    {
                        Color = DiscordColor.Purple,
                        Description = "Songshizz bot made by [Menno van Leeuwen](https://github.com/vleeuwenmenno)\nSongwhip made by [Wilson](https://songwhip.com/faq)\n\nI would like to thank Wilson especially for making his API publicly available for everyone to use!\n\n[Add bot to your server](https://discord.com/api/oauth2/authorize?client_id=860899901020700684&permissions=2147837952&scope=bot%20applications.commands) - [GitHub repository](https://github.com/vleeuwenmenno/songwhip-bot)",
                        Title = $"Songshizz bot {Utilities.Version}"
                    }.
                    WithFooter($"About requested by {ctx.Member.DisplayName}", ctx.Member.AvatarUrl)
                    .Build())
            );
        }
        
        [SlashCommand("changelog", "Changelogs about the bot")]
        public async Task ChangeLog(InteractionContext ctx)
        {
            await ctx.CreateResponseAsync(InteractionResponseType.ChannelMessageWithSource, new DiscordInteractionResponseBuilder()
                .AddEmbed(new DiscordEmbedBuilder
                    {
                        Color = DiscordColor.Purple,
                        Description = "** Update v1.4.0 **\n> - Replaced opt-in/opt-out with mention mode. Now you can make the bot resolve only when you really want it to!\n** Major update v1.3.1 **\n> - Renamed the bot to Songshizz\n> - Streamlined resolver flow to be more efficient and improved code readability.\n> - Added support for Spotify playlists\n> -Updated codebase to run on .NET 6.0\n> - Added this changelog command.\n",
                        Title = $"Songshizz bot {Utilities.Version}"
                    }.
                    WithFooter($"About requested by {ctx.Member.DisplayName}", ctx.Member.AvatarUrl)
                    .Build())
            );
        }

        [SlashCommand("toggle-mode",
            "Toggle between mentioning or auto-resolving music links mode.")]
        public async Task ToggleMode(InteractionContext ctx)
        {
            if (Blacklist.userMentionMode.Any(x => ctx.Member.Id == x))
            {
                Blacklist.userMentionMode.RemoveAll(x => x == ctx.Member.Id);
                Blacklist.Save();

                await ctx.CreateResponseAsync(InteractionResponseType.ChannelMessageWithSource,
                    new DiscordInteractionResponseBuilder()
                        .AddEmbed(new DiscordEmbedBuilder
                            {
                                Color = DiscordColor.DarkGreen,
                                Description =
                                    "You've been opted in of auto-resolving music links. <a:Happy:395314894364213248>\nThis means your links will be resolved automatically unless you mention Songshizz then the bot will ignore you.",
                                Title = $"Mode toggled to auto-resolving"
                            }
                            .WithFooter($"Toggle mode requested by {ctx.Member.DisplayName}", ctx.Member.AvatarUrl)
                            .Build())
                );
            }
            else
            {
                Blacklist.userMentionMode.Add(ctx.Member.Id);
                Blacklist.Save();
            
                await ctx.CreateResponseAsync(InteractionResponseType.ChannelMessageWithSource, new DiscordInteractionResponseBuilder()
                    .AddEmbed(new DiscordEmbedBuilder
                        {
                            Color = DiscordColor.DarkGreen,
                            Description =
                                "You've been opted in of mention resolving music links. <a:Happy:395314894364213248>\nThis means your links will not be resolved automatically unless you mention Songshizz then the bot will try to resolve the provided link.",
                            Title = $"Mode toggled to mention-resolving"
                        }
                        .WithFooter($"Toggle mode requested by {ctx.Member.DisplayName}", ctx.Member.AvatarUrl)
                        .Build())
                );
            }
        }
    }
}