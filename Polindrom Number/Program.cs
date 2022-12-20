using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Polindrom_Number
{
    public class Solution
    {
        public bool IsPalindrome(int x)
        {
            if (x < 0)
                return false;

            if((x > 9) && (x < 100))
            {
                if((x % 10 == x / 10))
                    return true;
                else
                    return false;
            }

            if(x < 10 && x > 0)
                return true;

            int numLen = (int)Math.Ceiling(Math.Log10((double)x));

            for(int i = 0; i <= numLen/2; i++)
            {

/*                Console.WriteLine((int)(x % Math.Pow(10, i + 1)) / (int)Math.Pow(10, i));
                Console.WriteLine((int)(x % Math.Pow(10, numLen - i)) / (int)Math.Pow(10, numLen - i - 1));
                Console.WriteLine();*/

                if ((int)(x % Math.Pow(10, i + 1)) / (int)Math.Pow(10, i) !=
                    (int)(x % Math.Pow(10, numLen - i)) / (int)Math.Pow(10, numLen - i - 1))
                    return false;
            }

            return true;
        }
    }

    internal class Program
    {
        static void Main(string[] args)
        {
            Solution solution = new Solution();

            Console.WriteLine(solution.IsPalindrome(101));
            ;
        }
    }
}
