#include <iostream>
#include <vector>
#include <unordered_map>
using namespace std;

/*

149. Max Points on a Line(Hard)
    Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane, 
    return the maximum number of points that lie on the same straight line.

https://leetcode.com/problems/max-points-on-a-line/description/

*/

typedef pair<float, int> Int_Pair;
typedef unordered_map<float, int> hashT;

class Solution {
public:
    int maxPoints(vector<vector<int>>& points) {

        if (points.size() <= 2) return points.size();

        int res = 0;
        float dX, dY, tg;

        for (int i = 0; i < points.size(); i++) {
            hashT hash;

            for (int j = 0; j < points.size(); j++) 
            {
                if (i == j) continue;


                dX = points[i][0] - points[j][0];
                dY = points[i][1] - points[j][1];

                if (dX == 0)
                    tg = 1000;
                else
                    tg = dY / dX;

                if (hash.find(tg) != hash.end())
                    hash[tg]++;
                else
                    hash.insert(Int_Pair(tg, 2));
            }

            for (hashT::const_iterator i = hash.begin(); i != hash.end(); ++i) {
                if (res < i->second)
                    res = i->second;

                //cout << i->first << " " << i->second << endl;
            }
        }

        return res;
    }
};

int main()
{
    vector<vector<int>> points = { {-6, -1}, {3, 1}, {12, 3} };

    Solution sol;
    cout << "Res = " << sol.maxPoints(points) << endl;
}
