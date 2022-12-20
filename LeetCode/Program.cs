using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Globalization;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

public class Solution
{

    public string ReverseWords(string s)
    {
        List<string> words = new List<string>();

        string str = "";
        string word = "";
       // int firstWordIndex = 0, lastWordIndex = 0;

        for (int i = s.Length - 1; i >= 0; i--)
        {
            if (s[i] != ' ')
            {
                word += s[i];
            }
            else
            {
                words.Add(word);
                word = "";
            }
        }
        words.Add(word);

        for(int i = words.Count - 1; i >= 0; i--)
        {
            str += words[i];

            if(i != 0)
            str += ' ';  
        }

        return str;
    } 
}

namespace LeetCode
{
    internal class Program
    {
        static void Main(string[] args)
        {
            string s = "God Ding";

            Solution sol = new Solution();

            Console.WriteLine(sol.ReverseWords(s));

            Console.WriteLine(sol.ReverseWords(s) == "doG gniD");

        }
    }
}
