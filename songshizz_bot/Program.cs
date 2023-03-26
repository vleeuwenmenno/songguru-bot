using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Net.Http;
using System.Threading.Tasks;
using DSharpPlus;
using DSharpPlus.Entities;
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
        private async Task LoadEnvironmentVariables(string filePath)
        {
            if (!File.Exists(filePath))
                return;

            foreach (var line in await File.ReadAllLinesAsync(filePath))
            {
                var parts = line.Split(
                    '=',
                    StringSplitOptions.RemoveEmptyEntries);

                if (parts.Length != 2)
                    continue;

                Environment.SetEnvironmentVariable(parts[0], parts[1]);
            }
        }

        private async Task SetupDiscordBot()
        {
            var token = Environment.GetEnvironmentVariable("BOT_TOKEN");
            _discord = new DiscordClient(new DiscordConfiguration()
            {
                Token = token,
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
            
            // Let's log what guilds the bot is in
            var client = new HttpClient();
            client.DefaultRequestHeaders.Add("Authorization", "Bot " + token);
            
            // TODO: This is a hacky way to get the bot's guilds, but it works for now
            // Seems like the DSharpPlus library doesn't properly fetch the guilds so we have to do it manually
            var guildsResponse = await client.GetAsync(@"https://discord.com/api/v6/users/@me/guilds");
            var response = await guildsResponse.Content.ReadAsStringAsync();
            var guilds = JsonConvert.DeserializeObject<List<DiscordGuild>>(response);
            
            Console.WriteLine($"Bot is in {guilds.Count} guilds:");
            foreach (var guild in guilds)
            {
                Console.WriteLine($"- {guild.Name} ({guild.Id})");
            }
            
            await _discord.UpdateStatusAsync(new DiscordActivity($" for music links in {guilds.Count} guilds!", ActivityType.Watching));
        }
        #endregion
        
        private static void Main(string[] args) =>
            new Program().MainAsync().GetAwaiter().GetResult();
        
        private async Task MainAsync()
        {
            Console.WriteLine($"Songshizz bot {Utilities.Version} starting ...");
            AppDomain.CurrentDomain.ProcessExit += OnProcessExit;
            
            await LoadEnvironmentVariables(Environment.CurrentDirectory + "/.env");
            await ListHelper.Load();
            await SetupDiscordBot();
            
            Console.WriteLine("Songshizz bot is ready!");
            await Task.Delay(-1);
        }

    }
}