package addTwoNumbers
import ListNode

class Solution {
    fun addTwoNumbers(l1: ListNode?, l2: ListNode?): ListNode? {
        var flag=false
        val result=ListNode(0)
        var now=result
        var ll1=l1
        var ll2=l2
        do {
            val ll1Value= ll1?.value ?: 0
            val ll2Value= ll2?.value ?: 0
            var sum=ll1Value+ll2Value
            if (flag) sum++
            if (sum<10){
                now.value=sum
                flag=false
            }else{
                now.value=sum-10
                flag=true
            }
            ll1=ll1?.next
            ll2=ll2?.next
            val hasNext=ll1!=null||ll2!=null||flag
            if (hasNext) {
                now.next = ListNode(0)
                now = now.next as ListNode
            }
        } while (hasNext)
        return result
    }
}

fun main() {
    val ll1ln1=ListNode(2)
    val ll1ln2=ListNode(4)
    val ll1ln3=ListNode(3)
    ll1ln1.next=ll1ln2
    ll1ln2.next=ll1ln3

    val ll2ln1=ListNode(5)
    val ll2ln2=ListNode(6)
    val ll2ln3=ListNode(4)
    ll2ln1.next=ll2ln2
    ll2ln2.next=ll2ln3
    /*val ll2ln1=ListNode(5)
    val ll1ln1=ListNode(5)*/

    val solution=Solution()
    val result=solution.addTwoNumbers(ll1ln1,ll2ln1)
    var node=result
    do {
        print(node?.value)
        node=node?.next
    }while (node!=null)
}