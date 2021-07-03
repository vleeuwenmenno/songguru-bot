using System;
using System.IO;

namespace songwhip_bot
{
    public class Utilities
    {
        public static string Version
        {
            get
            {
                string v = "v1.1.0";
                if (File.Exists($"{Environment.CurrentDirectory}/BRANCH") && File.Exists($"{Environment.CurrentDirectory}/COMMIT"))
                {
                    string hash = File.ReadAllText($"{Environment.CurrentDirectory}/COMMIT").Replace("\n", "");
                    string branch = File.ReadAllText($"{Environment.CurrentDirectory}/BRANCH").Replace("\n", "");
                
                    return $"{v}-{hash}-{branch}";
                }

                return $"{v}-ucommit-ubranch";
            }
        }
    }
}