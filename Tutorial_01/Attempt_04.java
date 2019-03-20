class Attempt_04 extends Thread {

    // number of increments per thread
    static int N = 10000;

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
                    q = true;
                    // pre-protocol section
                    while (true == p) {
                        q = false;
                        q = true;
//                        System.out.println("p pre-protocol");
                    }

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
                    sleep(1);
                    //System.out.println("q non-critical");
                    p = true;
                    // pre-protocol section
                    while (true == q) {
                        p = false;
                        p = true;
                        //                       System.out.println("q pre-protocol");
                    }

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
        Attempt_04 p = new Attempt_04();
        Attempt_04 q = new Attempt_04();
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
