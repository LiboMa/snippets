// Java program to count number of nodes in a linked list 
  
/* Linked list Node*/
import java.time.*;
import java.util.concurrent.TimeUnit;
import java.util.*;

class Node 
{ 
    int data; 
    Node next; 
    Node(int d)  { data = d;  next = null; } 
} 
  
// Linked List class 
class LinkedList 
{ 
    Node head;  // head of list 
  
    /* Inserts a new Node at front of the list. */
    public void push(int new_data) 
    { 
        /* 1 & 2: Allocate the Node & 
                  Put in the data*/
        Node new_node = new Node(new_data); 
  
        /* 3. Make next of new Node as head */
        new_node.next = head; 
  
        /* 4. Move the head to point to new Node */
        head = new_node; 
    } 
  
    /* Returns count of nodes in linked list */
    public int getCount() 
    { 
        Node temp = head; 
        int count = 0; 
        while (temp != null) 
        { 
            count++; 
            temp = temp.next; 
        } 
        return count; 
    } 

    public List getData(){
        Node temp = head;
        List list = new ArrayList();
        while (temp != null){

            list.add(temp.data);
            temp = temp.next;
        }
        return list;
    }
  
    /* Driver program to test above functions. Ideally 
       this function should be in a separate user class. 
       It is kept here to keep code compact */
    public static void main(String[] args) 
    { 
        /* Start with the empty list */
        long start = System.nanoTime();
        LinkedList llist = new LinkedList(); 
        llist.push(1); 
        llist.push(3); 
        llist.push(1); 
        llist.push(2); 
        llist.push(1); 
  
        System.out.println("Count of nodes is " + llist.getCount()); 
        System.out.println(llist.getData()); 
        long finish = System.nanoTime();
        System.out.printf("%f ms", (finish - start)/1000.0/1000.0);
    } 
} 
