using System;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;
using DSharpPlus;
using DSharpPlus.EventArgs;
using DSharpPlus.SlashCommands;
using Microsoft.Extensions.DependencyInjection;
using Newtonsoft.Json;
using SongshizzBot;
using SongshizzBot.Helpers;
using Utilities = SongshizzBot.Utilities;

namespace songshizz_bot
{
    // Invite the bot with:
    // DEVELOPER BOT: https://discord.com/oauth2/authorize?client_id=861018175935873054&permissions=2147837952&scope=bot%20applications.commands
    // LIVE BOT: https://discord.com/api/oauth2/authorize?client_id=860899901020700684&permissions=2147870784&scope=bot%20applications.commands
    public class Program
    {
        private static DiscordClient _discord;
        private SlashCommandsExtension _slash;
        
        #region Events
        private static void OnProcessExit (object sender, EventArgs e)
        {
            _discord.DisconnectAsync();
            Console.WriteLine($"Songshizz bot {Utilities.Version} closing shop ~");
        }
        
        private static Task DiscordOnReady(DiscordClient sender, ReadyEventArgs e)
        {
            return Task.CompletedTask;
        }
        
        private async Task MessageCreated(DiscordClient sender, MessageCreateEventArgs e)
        {
            await Songshizz.ProcessMessage(e, _discord);
        }
        #endregion

        #region Setups
        private async Task SetupDiscordBot()
        {
            _discord = new DiscordClient(new DiscordConfiguration()
            {
                Token = Environment.GetEnvironmentVariable("BOT_TOKEN"),
                TokenType = TokenType.Bot,
                Intents = DiscordIntents.AllUnprivileged
            });
            
            _discord.MessageCreated += MessageCreated;
            _discord.Ready += DiscordOnReady;
            
            _slash = _discord.UseSlashCommands(new SlashCommandsConfiguration
            {
                Services = new ServiceCollection()
                    .BuildServiceProvider()
            });
            
            _slash.RegisterCommands<SongshizzCommands>();
            
            await _discord.ConnectAsync();
        }

        private async Task SetupBlacklist()
        {
            if (File.Exists($"{Environment.CurrentDirectory}/users-blacklist.json"))
                Blacklist.userMentionMode = JsonConvert.DeserializeObject<List<ulong>>(await File.ReadAllTextAsync($"{Environment.CurrentDirectory}/users-blacklist.json"));
         
            if (File.Exists($"{Environment.CurrentDirectory}/guild-blacklist.json"))
                Blacklist.guildBlacklist = JsonConvert.DeserializeObject<List<ulong>>(await File.ReadAllTextAsync($"{Environment.CurrentDirectory}/guild-blacklist.json"));
            
            if (Blacklist.userMentionMode == null)
                Blacklist.userMentionMode = new List<ulong>();

            if (Blacklist.guildBlacklist == null)
                Blacklist.guildBlacklist = new List<ulong>();
        }
        #endregion
        
        private static void Main(string[] args) =>
            new Program().MainAsync().GetAwaiter().GetResult();
        
        private async Task MainAsync()
        {
            Console.WriteLine($"Songshizz bot {Utilities.Version} starting ...");
            AppDomain.CurrentDomain.ProcessExit += OnProcessExit;

            await SetupBlacklist();
            await SetupDiscordBot();
            
            Console.WriteLine("Songshizz bot is ready!");
            await Task.Delay(-1);
        }

    }
}