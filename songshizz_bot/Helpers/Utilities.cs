using System;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;
using DSharpPlus.EventArgs;
using SongshizzBot.Helpers;

namespace SongshizzBot
{
    public static class Utilities
    {
        public static string Version
        {
            get
            {
                string v = "v1.2.0";
                if (File.Exists($"{Environment.CurrentDirectory}/BRANCH") && File.Exists($"{Environment.CurrentDirectory}/COMMIT"))
                {
                    string hash = File.ReadAllText($"{Environment.CurrentDirectory}/COMMIT").Replace("\n", "");
                    string branch = File.ReadAllText($"{Environment.CurrentDirectory}/BRANCH").Replace("\n", "");
                
                    return $"{v}-{hash}-{branch}";
                }

                return $"{v}-ucommit-ubranch";
            }
        }
        
        public static bool IsBlackListed(MessageCreateEventArgs e)
        {
            if (Blacklist.userBlacklist.Any(x => x == e.Message.Author.Id))
            {
                Console.WriteLine(
                    $"music link detected for {e.Message.Author.Username}#{e.Message.Author.Id} on {e.Guild.Name} ({e.Guild.Id}) but doing nothing because user opted-out~");
                return true;
            }

            return false;
        }

        public static bool IsYouTubeLink(string msg)
        {
            if (string.IsNullOrEmpty(msg))
                return false;

            return (msg.StartsWith("https://youtube.com/") ||
                    msg.StartsWith("https://www.youtube.com/"));
        }

        public static bool IsLinkMessage(string msg)
        {
            if (string.IsNullOrEmpty(msg))
                return false;
            
            return (msg.StartsWith("https://www.spotify.com/") ||
                    msg.StartsWith("https://spotify.com/") ||
                    msg.StartsWith("https://open.spotify.com/") ||
                    msg.StartsWith("https://deezer.com/") ||
                    msg.StartsWith("https://www.deezer.com/") ||
                    msg.StartsWith("https://youtube.com/") ||
                    msg.StartsWith("https://www.youtube.com/") ||
                    msg.StartsWith("https://music.youtube.com/") ||
                    msg.StartsWith("https://music.apple.com/") ||
                    msg.StartsWith("https://tidal.com/") ||
                    msg.StartsWith("https://www.tidal.com/"));
        }

        public static string ExtractLink(string messageContent)
        {
            Match url = Regex.Match(messageContent, @"(http|ftp|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?");
            return url.ToString();
        }
        
        public static string ExtractLinkMessage(string messageContent)
        {
            Match url = Regex.Match(messageContent, @"(http|ftp|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?");
            return messageContent.Replace(url.ToString(), "");
        }

        public static bool IsSpotifyPlaylist(string msg)
        {
            return IsLinkMessage(msg) && msg.Contains("/playlist/");
        }
    }
}