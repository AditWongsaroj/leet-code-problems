#include <map>
#include <string>
#include <iostream>
#include <minmax.h>

class Solution
{
public:
  int lengthOfLongestSubstring(std::string s)
  {
    int longest = 0;
    auto sz = s.size();
    auto seen_map = std::map<char, int>{};
    
    for (auto i = 0; i < sz; i++)
    {
      auto ch = s[i];
      if(seen_map.contains(ch)){
        if (auto sms = seen_map.size(); longest < sms)
        {
          longest = sms;
        }
        auto idx = seen_map[ch];

        for (auto it = seen_map.begin(); it != seen_map.end();)
        {
          if (it->second <= idx)
            it = seen_map.erase(it);
          else
            ++it;
        }
      }
      seen_map[ch] = i;
    }
    if (auto sms = seen_map.size(); longest < sms)
    {
      longest = sms;
    }

    return longest;
  }
  int lolss10xfaster(std::string s)
  {
    int len = s.size();
    int char_map[127] = { 0 };
    int longest = 0;
    for (int left = 0, right = 0; right < len; right++)
    {
      left = max(char_map[s[right]], left);
      longest = max(longest, right - left + 1);
      char_map[s[right]] = right + 1;
    }
    return longest;
  }
};

int main()
{
  std::cout << Solution().lolss10xfaster("dvdf");
}