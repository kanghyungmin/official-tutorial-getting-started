
// 1. leetcode 1개 풀기
// 2. go 2개 투토리얼 확인
// 3. nestjs DDD best praticess
// - https://khalilstemmler.com/articles/categories/domain-driven-design/ (6/17)
// - https://junho2343.github.io/posts/clean-architecture-hexagonal-architecture-with-nestjs
// - https://eocoding.tistory.com/36


function frequencySort(nums: number[]): number[] {
    const freqMap = new Map<number, number>();

    // Count the frequency of each number in the array
    for (const num of nums) {
        freqMap.set(num, (freqMap.get(num) || 0) + 1)
        
        }

    // Sort the array based on frequency (increasing) and then by value (decreasing)
    nums.sort((a, b) => {
        const freqA = freqMap.get(a)!;
        const freqB = freqMap.get(b)!;
        
        // If frequencies are different, sort by frequency
        if (freqA !== freqB) {
            return freqA - freqB;
        }
        
        // If frequencies are the same, sort by value in decreasing order
        return b - a;
    });

    return nums;
}

let nums = [-1,1,-6,4,5,-6,1,4,1]
// nums=[1,1,2,2,2,3]

console.log(frequencySort(nums)); // [1,3,3,2,2]




