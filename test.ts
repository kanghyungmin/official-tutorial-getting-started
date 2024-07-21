
// 1. leetcode 3개 풀기
// 2. nestjs DDD best praticess
// - https://github.com/kanghyungmin/domain-driven-hexagon?tab=readme-ov-file
// 3. go 2개 투토리얼 확인



//   Definition for a binary tree node.
function luckyNumbers (matrix: number[][]): number[] {
    let minValues : number[] =[], maxValues : number[] = []
    let [row, col] = [matrix.length , matrix[0].length ]
    for(let i=0; i<row;i++) {
        minValues.push(Math.min(...matrix[i]))
    }

    for(let i=0; i<col;i++) {
        let temp = 0;
        for(let j=0; j<row;j++) {
            temp = Math.max(matrix[j][i], temp)
        }
        maxValues.push(temp)
    }
    let retVal : number[] =[]

    minValues.map( v => {
        if(maxValues.includes(v))
            retVal.push(v);
    })

    return retVal as number[]
};
  
let matrix = [[1,10,4,2],[9,3,8,7],[15,16,17,12]]
console.log(luckyNumbers(matrix))