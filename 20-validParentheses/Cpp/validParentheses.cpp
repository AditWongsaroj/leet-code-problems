#include <string>
#include <vector>
#include <iostream>

class Solution
{
public:
    bool isValid(std::string s)
    {
        std::vector<char> seen;

        for (int i = 0; i < s.size(); i++)
        {
            switch (s.at(i))
            {
            case '[':
                seen.push_back('[');
                break;
            case '{':
                seen.push_back('{');
                break;
            case '(':
                seen.push_back('(');
                break;
            case ']':
                if (seen.size() == 0 || seen[seen.size() - 1] != '[')
                {
                    return false;
                }
                seen.pop_back();
                break;
            case ')':
                if (seen.size() == 0 || seen[seen.size() - 1] != '(')
                {
                    return false;
                }
                seen.pop_back();
                break;
            case '}':
                if (seen.size() == 0 || seen[seen.size() - 1] != '{')
                {
                    return false;
                }
                seen.pop_back();
                break;
            }
        }
        return seen.size() == 0;
    }
};