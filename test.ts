
// 1. leetcode 1개 풀기
// 2. go 2개 투토리얼 확인
// 3. nestjs DDD best praticess
// - https://github.com/kanghyungmin/domain-driven-hexagon?tab=readme-ov-file
//


function sortPeople(names: string[], heights: number[]): string[] {
    let mappedArr : Array<{height: number; name: string}> = names.map((name, index) => ({
        height : heights[index],
        name : name,
    }))
    mappedArr.sort( (a,b,) => {
        if(a.height < b.height)
            return 1;
        return -1
    })

    let retVal : string[] = []
    for(const i of mappedArr.values()) {
        retVal.push(i.name)
    }

    return retVal
};



let names = ["Mary","John","Emma"], heights = [180,165,170]
console.log(sortPeople(names,heights))
names = ["Alice","Bob","Bob"], heights = [155,185,150]
console.log(sortPeople(names,heights))