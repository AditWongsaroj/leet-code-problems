#define DOCTEST_CONFIG_IMPLEMENT_WITH_MAIN
#include "doctest.h"

#include "validParentheses.cpp"

#include <variant>
#include <vector>
#include <string>
#include <iostream>

TEST_CASE("testing the validParentheses function")
{
    struct Target
    {
        std::vector<std::string> inputs;
        bool want;
    };

    struct Target targets[7]{
        {{"()", "{}", "[]"}, true},
        {{"()[]{}"}, true},
        {{"([{}])", "[]({})", "[()]{}"}, true},
        {{"(]", "{)"}, false},
        {{"([)]"}, false},
        {{"]", ")", "}"}, false},
        {{"(", "[", "{"}, false},
    };

    for (int i = 0; i < (sizeof(targets) / sizeof(targets[0])); i++)
    {
        auto target = targets[i].inputs;

        bool want = targets[i].want;

        for (int j = 0; j < target.size(); j++)
        {
            auto got = Solution().isValid(target[j]);
            std::cout << "Test " << i + 1 << " - " << j + 1 << ": '" << target[j] << "'\n";
            std::cout << "\t\tgot: " << got << ", want: " << want << "\n\n";
            CHECK(got == want);
        }
    };
}