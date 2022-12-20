using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Concatenation_of_Consecutive_Binary_Numbers
{
    public class Solution
    {
        public int ConcatenatedBinary(int n)
        {
            long res = 1;
            for(int i = 2; i <= n; i++)
            {
                int a = 1;
                while (a <= i)
                {
                    res <<= 1;

                    a <<= 1; 
                }

                while (res >= (Math.Pow(10, 9) + 7))
                    res -= (int)(Math.Pow(10, 9) + 7);

                res += i;
            }

            res = (int)Math.Abs(res % (Math.Pow(10, 9) + 7));

            return (int)res;
        }
    }

    internal class Program
    {
        static void Main(string[] args)
        {
            Solution solution = new Solution();

            Console.WriteLine(solution.ConcatenatedBinary(12));
        }
    }
}
