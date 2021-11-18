using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Threading.Tasks;
using DSharpPlus;
using DSharpPlus.Entities;
using DSharpPlus.SlashCommands;
using Newtonsoft.Json;
using RestSharp;
using SongshizzBot.Helpers;
using Songwhip;
using Songwhip.Models;
using JsonSerializer = System.Text.Json.JsonSerializer;

namespace SongshizzBot
{
    public class SongshizzCommands : SlashCommandModule
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
                        Description = "** Major update v1.3.1 **\n> - Renamed the bot to Songshizz\n> - Streamlined resolver flow to be more efficient and improved code readability.\n> - Added support for Spotify playlists\n> -Updated codebase to run on .NET 6.0\n> - Added this changelog command.\n",
                        Title = $"Songshizz bot {Utilities.Version}"
                    }.
                    WithFooter($"About requested by {ctx.Member.DisplayName}", ctx.Member.AvatarUrl)
                    .Build())
            );
        }
        
        [SlashCommand("opt-in", "Enable auto-resolving music links for yourself. Songshizz bot will try resolve your music links.")]
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
        
        [SlashCommand("opt-out", "Disable auto-resolving music links for yourself. Songshizz bot won't bother you ever again!")]
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
    }
}