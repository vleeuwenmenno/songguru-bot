using System;
using System.IO;
using System.Text.Json;
using System.Threading.Tasks;
using DSharpPlus;
using DSharpPlus.EventArgs;
using DSharpPlus.SlashCommands;
using Microsoft.Extensions.DependencyInjection;
using Newtonsoft.Json;

namespace DiscordCodBot
{
    // Invite the bot with:
    // https://discord.com/api/oauth2/authorize?client_id=860899901020700684&permissions=2147764224&scope=applications.commands%20bot
    public class Program
    {
        public static Options _options;
        public DiscordClient _discord;
        public SlashCommandsExtension _slash;

        static void Main(string[] args) =>
            new Program().MainAsync().GetAwaiter().GetResult();
        
        public async Task MainAsync()
        {
            Console.WriteLine("Songwhip bot v1.0 starting ...");
            
            _options = Options.LoadConfig();

            _discord = new DiscordClient(new DiscordConfiguration()
            {
                Token = _options.botToken,
                TokenType = TokenType.Bot,
                Intents = DiscordIntents.AllUnprivileged
            });
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

        private Task DiscordOnReady(DiscordClient sender, ReadyEventArgs e)
        {
            return Task.CompletedTask;
        }
    }
}