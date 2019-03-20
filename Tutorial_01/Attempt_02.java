class Attempt_02 extends Thread {

    // number of increments per thread
    static int N = 1000;

    // shared data
    static volatile int counter = 0;

    // protocol variables
    // ...
    //static volatile boolean turn = false;
    static volatile boolean p = false;
    static volatile boolean q = false;

    public void run() {
        try {
            // Thread p
            if (this.getName().equals("Thread-0")) {
                int temp;
                for (int i = 0; i < N; i++) {
                    // non-critical section
                    sleep(1);
                    //System.out.println("p non-critical");
                    // pre-protocol section
                    while (true == p) {
                        //System.out.println("p pre-protocol");
                    }

                    q = true;

                    // critical section
                    temp = counter;
                    counter = temp + 1;

                    // post-protocol section
                    q = false;
                }
            // Thread q
            } else if (this.getName().equals("Thread-1")) {
                int temp;
                for (int i = 0; i < N; i++) {
                    // non-critical section
                    //System.out.println("q non-critical");
                    sleep(1);
                    // pre-protocol section
                    while (true == q) {
                        //System.out.println("q pre-protocol");
                    }

                    p = true;

                    // critical section
                    temp = counter;
                    counter = temp + 1;

                    // post-protocol section
                    p = false;
                }
            }
        } catch (
                InterruptedException e) {

        }

    }

    public static void main(String[] args) {
        Attempt_02 p = new Attempt_02();
        Attempt_02 q = new Attempt_02();
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
