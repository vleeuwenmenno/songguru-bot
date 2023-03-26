using System;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;
using Newtonsoft.Json;

namespace SongshizzBot.Helpers
{
    public abstract class ListHelper
    {
        public static List<ulong> MentioningMode;
        public static List<ulong> KeepMessageMode;

        public static async Task Save()
        {
            // Make sure the data directory exists
            if (!Directory.Exists($"{Environment.CurrentDirectory}/data"))
                Directory.CreateDirectory($"{Environment.CurrentDirectory}/data");
            
            await File.WriteAllTextAsync($"{Environment.CurrentDirectory}/data/users-mentioning-mode.json",
                JsonConvert.SerializeObject(MentioningMode, Formatting.Indented));
            
            await File.WriteAllTextAsync($"{Environment.CurrentDirectory}/data/users-keep-message-mode.json",
                JsonConvert.SerializeObject(KeepMessageMode, Formatting.Indented));
        }

        public static async Task Load()
        {
            // Load the users-mentioning-mode.json file
            if (File.Exists($"{Environment.CurrentDirectory}/data/users-mentioning-mode.json"))
            {
                var json = await File.ReadAllTextAsync($"{Environment.CurrentDirectory}/data/users-mentioning-mode.json");
                MentioningMode = JsonConvert.DeserializeObject<List<ulong>>(json);
            }
            else
            {
                MentioningMode = new List<ulong>();
            }
            
            // Load the users-keep-message-mode.json file
            if (File.Exists($"{Environment.CurrentDirectory}/data/users-keep-message-mode.json"))
            {
                var json = await File.ReadAllTextAsync($"{Environment.CurrentDirectory}/data/users-keep-message-mode.json");
                KeepMessageMode = JsonConvert.DeserializeObject<List<ulong>>(json);
            }
            else
            {
                KeepMessageMode = new List<ulong>();
            }
        }
    }
}