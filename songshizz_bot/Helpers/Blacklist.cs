using System;
using System.Collections.Generic;
using System.IO;
using Newtonsoft.Json;

namespace SongshizzBot.Helpers
{
    public class Blacklist
    {
        public static List<ulong> userBlacklist;

        public static List<ulong> guildBlacklist;

        public static void Save()
        {
            File.WriteAllText($"{Environment.CurrentDirectory}/users-blacklist.json",
                JsonConvert.SerializeObject(userBlacklist, Formatting.Indented));
            
            File.WriteAllText($"{Environment.CurrentDirectory}/guild-blacklist.json",
                JsonConvert.SerializeObject(guildBlacklist, Formatting.Indented));
        }
    }
}