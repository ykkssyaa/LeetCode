package org.LeetCode;

public class Main {
    public static void main(String[] args) {

        System.out.println(Solution.romanToInt("MDCXCV"));
        System.out.println(Solution.romanToInt("MCMXCIV"));
    }
};

class Solution {
    static public int romanToInt(String s) {

        boolean flag = true;
        int res = 0;
        for(int i = 0; i < s.length() - 1; i++){
            if(charToInt(s.charAt(i)) < charToInt(s.charAt(i + 1))){
                res += charToInt(s.charAt(i + 1)) - charToInt(s.charAt(i));
                i++;
                if(i >= s.length() - 1)
                    flag = false;
            }
            else{
                res += charToInt(s.charAt(i));
            }

        }

        if(flag){
            res += charToInt(s.charAt(s.length() - 1));
        }

        return res;
    }

    static private int charToInt(char roman){
        return switch (roman) {
            case 'I' -> 1;
            case 'V' -> 5;
            case 'X' -> 10;
            case 'L' -> 50;
            case 'C' -> 100;
            case 'D' -> 500;
            case 'M' -> 1000;
            default -> -1;
        };
    }
}