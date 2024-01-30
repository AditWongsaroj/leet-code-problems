#include <vector>
#include <algorithm>
#include <cmath>

bool is_integer(float);

float medianTwoSortedArrays(std::vector<int> arr1, std::vector<int> arr2){
    auto size = arr1.size() + arr2.size();
    std::vector<int> arr_combined;
    arr_combined.reserve(size);
    arr_combined.insert(arr_combined.end(), arr1.begin(), arr1.end());
    arr_combined.insert(arr_combined.end(), arr2.begin(), arr2.end());
    std::sort( arr_combined.begin(), arr_combined.end());

    auto mid = size / 2.0;
    if(is_integer(mid)){
        return (arr_combined[mid-1] + arr_combined[mid])/2.0;
    }
    return arr_combined[mid];
}

bool is_integer(float k)
{
  return std::floor(k) == k;
}

