/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* prev;
    Node* next;
    Node* child;
};
*/
class Solution {
public:
    Node* flatten(Node* head) {
        if (head == NULL) {
            return head;
        }
        
        Node dummy;
        preorder(head, &dummy);
        dummy.next->prev = NULL;
        return dummy.next;
    }
    
    Node* preorder(Node* cur, Node* prev) {
        if (cur == NULL) {
            return prev;
        }
        
        cur->prev = prev;
        prev->next = cur;
        
        Node* tmp = cur->next; // 相同层的下一个节点
        Node* tail = cur;
        if (cur->child != NULL) {
            tail = preorder(cur->child, cur); // 返回展开后最后一个节点
            cur->child = NULL;
        }
        
        return preorder(tmp, tail);
    }
};