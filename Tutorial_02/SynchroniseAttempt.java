import java.util.LinkedList;
import java.util.Queue;
import java.util.Random;
import java.util.concurrent.Semaphore;

public class SynchroniseAttempt extends Thread{
    private String name;
    private static Queue<Integer> integerQueue = new LinkedList<>();
    private static int numOperations = 10;
    private static final int MAX_NUMBER = 1000;
    private static Semaphore semaphore = new Semaphore(numOperations);

    SynchroniseAttempt(String name){
        this.name = name;
    }

    private void addElement() {
        Random random = new Random();
        int newElement = random.nextInt(MAX_NUMBER);
        try {
            semaphore.acquire();
            integerQueue.offer(newElement);
            System.out.println(this.name + " adds " + newElement);
            System.out.println(integerQueue.toString());
            semaphore.release();
        } catch (InterruptedException e){ }
    }

    private void removeElement() {
        try {
            semaphore.acquire();
            if (!integerQueue.isEmpty()) {
                int polledElement = integerQueue.poll();
                System.out.println(this.name + " removes " + polledElement);
                System.out.println(integerQueue.toString());
            } else {
                System.out.println("Queue is empty.");
                System.out.println(integerQueue.toString());

            }
            semaphore.release();
        } catch (InterruptedException e) { }

//            try {
//                while (integerQueue.isEmpty()) {
//                    wait();
//                }
//                int polledElement = integerQueue.poll();
//                System.out.println(this.name + " removes " + polledElement);
//                System.out.println(integerQueue.toString());
//
//            } catch (InterruptedException e) {
//        }
    }

    @Override
    public void run() {
        // Write thread
        if (this.name.equals("sync1")) {
            for (int i = 0; i < numOperations; i++) {
                addElement();
            }
            // Read thread
        } else if (this.name.equals("sync2")) {
            for (int i = 0; i < numOperations; i++) {
                removeElement();
            }
        }
    }

    public static void main(String[] args){
        SynchroniseAttempt sync1 = new SynchroniseAttempt("sync1");
        SynchroniseAttempt sync2 = new SynchroniseAttempt("sync2");

        sync1.start();
        sync2.start();
        try {
            sync1.join();
            sync2.join();
        } catch (InterruptedException e) { }
//        Semaphore testSemaphore = new Semaphore(10);
//        for (int i = 0; i < 10; i++) {
//            testSemaphore.release(1);
//            System.out.println(
//                    "available permits: " + testSemaphore.availablePermits());
//        }
    }
}
