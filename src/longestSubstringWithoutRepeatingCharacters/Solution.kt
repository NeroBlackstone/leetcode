package longestSubstringWithoutRepeatingCharacters

class Solution {
    fun lengthOfLongestSubstring(s: String): Int {
        var subString=""
        var longestStrSize=0
        var longestStrIndex=0
        var nowIndex=0
        s.forEach {
            if (subString.contains(it)) {
                val newSubStringStartIndex=subString.indexOf(it)+1
                nowIndex=newSubStringStartIndex+longestStrIndex
                subString=subString.substring(newSubStringStartIndex)+it
            }
            else
                subString+=it
            if (subString.length>longestStrSize){
                longestStrSize=subString.length
                longestStrIndex=nowIndex
            }
        }
        return longestStrSize
    }
}

fun main() {
    val solution=Solution()
    //val s="abcabcbb"
    val s="aab"
    print(solution.lengthOfLongestSubstring(s))
}