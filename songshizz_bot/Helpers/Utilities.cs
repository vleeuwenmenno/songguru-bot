using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;
using DSharpPlus;
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
                string v = "v1.5.2";
                if (File.Exists($"{Environment.CurrentDirectory}/BRANCH") && File.Exists($"{Environment.CurrentDirectory}/COMMIT"))
                {
                    string hash = File.ReadAllText($"{Environment.CurrentDirectory}/COMMIT").Replace("\n", "");
                    string branch = File.ReadAllText($"{Environment.CurrentDirectory}/BRANCH").Replace("\n", "");
                
                    return $"{v}-{hash}-{branch}";
                }

                return $"{v}-ucommit-ubranch";
            }
        }
        
        /// <summary>
        /// Checks if the requesting user has mentioning mode enabled and/or has properly mentioned the current bot user id.
        /// </summary>
        /// <param name="e">Message event arguments</param>
        /// <param name="discord">Discord client interface</param>
        /// <returns>Returns true unless mentioning mode is disabled or mentioning mode is enabled but the message contained a valid mention to the current bot user id.</returns>
        public static bool IsMentioningMode(MessageCreateEventArgs e, DiscordClient discord)
        {
            if (ListHelper.MentioningMode.Any(x => x == e.Message.Author.Id))
            {
                if (e.Message.Content.Contains(discord.CurrentUser.Id.ToString()))
                    return false;
                
                Console.WriteLine($"music link detected for {e.Message.Author.Username}#{e.Message.Author.Id} on {e.Guild.Name} ({e.Guild.Id}) but doing nothing because user has mentioning mode enabled but didn't mention the bot~");
                return true;
            }

            return false;
        }

        public static bool IsYouTubeLink(string msg)
        {
            if (string.IsNullOrEmpty(msg))
                return false;

            return (msg.StartsWith("https://youtube.com/") ||
                    msg.StartsWith("https://www.youtube.com/") ||
                    msg.StartsWith("https://youtu.be/"));
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
                    msg.StartsWith("https://youtu.be/") ||
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
            return msg.StartsWith("https://www.spotify.com/") ||
                   msg.StartsWith("https://spotify.com/") ||
                   msg.StartsWith("https://open.spotify.com/") && 
                   msg.Contains("/playlist/");
        }

        public static bool IsDeezerPlaylist(string msg)
        {
            return msg.StartsWith("https://deezer.com/") ||
                   msg.StartsWith("https://www.deezer.com/") && 
                   msg.Contains("/playlist/");
        }

        public static string[] ExtractMentions(string messageContent, ulong currentUserId)
        {
            // Let's search in message content for mentions <@userid>
            var matches = Regex.Matches(messageContent, @"<@(\d+)>");
            var mentions = new List<string>();
            foreach (Match match in matches)
            {
                // If the mention is not the current bot user id, add it to the list
                if (match.Groups[1].Value != currentUserId.ToString())
                    mentions.Add(match.Groups[1].Value);
            }
            
            return mentions.ToArray();
        }
    }
}