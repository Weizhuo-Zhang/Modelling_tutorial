class Attempt_01 extends Thread {

    // number of increments per thread
    static int N = 100000;

    // shared data
    static volatile int counter = 0;

    // protocol variables
    // ...
    static volatile boolean turn = false;

    public void run() {
        try {
            if (this.getName().equals("Thread-0")) {
                int temp;
                for (int i = 0; i < N; i++) {
                    // non-critical section
                    sleep(1);
                    // pre-protocol section
                    while (true != turn) {
                    }
                    // critical section
                    temp = counter;
                    counter = temp + 1;

                    // post-protocol section
                    turn = false;
                }
            } else if (this.getName().equals("Thread-1")) {
                int temp;
                for (int i = 0; i < N; i++) {
                    // non-critical section
                    sleep(1);
                    // pre-protocol section
                    while (false != turn) {
                    }

                    // critical section
                    temp = counter;
                    counter = temp + 1;

                    // post-protocol section
                    turn = true;
                }
            }
        } catch (
                InterruptedException e) {

        }

    }

    public static void main(String[] args) {
        Attempt_01 p = new Attempt_01();
        Attempt_01 q = new Attempt_01();
        p.start();
        q.start();
        try {
            p.join();
            q.join();
        } catch (InterruptedException e) {
        }
        System.out.println("The final value of the counter is " + counter);
    }
}
