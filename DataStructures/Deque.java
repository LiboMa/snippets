import java.util.*;

public class Deque {

    public static void main(String[] args)  throws IllegalStateException {
    
        Deque<Integer> dq = new ArrayDeque<Integer>();

        dq.add(1);
        dq.add(2);
        dq.add(3);
        dq.add(4);

        System.out.println("Deuqe: " + dq);
        System.out.println("Deque's head: " + dq.getFirst());
    }
}
