#define DOCTEST_CONFIG_IMPLEMENT_WITH_MAIN
#include "doctest.h"

#include "validParentheses.cpp"

#include <variant>
#include <vector>
#include <string>
#include <print>

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
            std::print("Test {0} - {1}: '{2}'\n", i + 1, j + 1, target[j]);
            std::print("\tgot: {0}\n\twant: {1}\n\n", got, want);
            CHECK(got == want);
        }
    };
}