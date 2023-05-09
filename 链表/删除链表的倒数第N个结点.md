力扣题目链接：https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

给你一个链表，删除链表的倒数第 `n` 个结点，并且返回链表的头结点
---------------------- C++ --------------------------------------

ListNode* removeNthFromEnd(ListNode* head, int n) {
    // 双指针经典应用：如果要删除倒数第n个节点，先让快指针往后移动n+1个位置，然后同时移动快慢指针，直到快指针指向空。此时慢指针指向被删除节点的前一个节点
    // 为啥是n+1？因为你想移除倒数第n个，你最少是有1个节点，我如果想删除某个节点，需要指向被删除节点q的前一个节点k(这里都在虚拟头结点的环境下考虑)
    // 然后让k指向q的下一个节点。因此，我要使用n+1步先来找到"目标距离"，获取快节点和慢节点的“差一值”间距
    // 用虚拟头结点最好，如果不用的话需要单独判断，比如：总共3个节点，让删除倒数第三个，再用上面的方法，越界了。
         // 虚拟头结点
        ListNode *hunmmy = (ListNode*)malloc(sizeof(ListNode));
        hunmmy->next = head;
        ListNode* fast = hunmmy;
        ListNode* slow = hunmmy;
        while(n-- && fast != NULL)
        {
            fast = fast->next;
        }
        // 这是+1的那一步，就是指向要删除节点的下一个节点
        fast = fast->next;
        while(fast != NULL)
        {
            slow = slow->next;
            fast = fast->next;
        }
        slow->next = slow->next->next;
        return hunmmy->next;
        }
