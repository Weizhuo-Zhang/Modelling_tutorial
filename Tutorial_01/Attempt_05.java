class Attempt_05 extends Thread {

    // number of increments per thread
    static int N = 100000;

    // shared data
    static volatile int counter = 0;

    // protocol variables
    // ...
    static volatile boolean turn = false;
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
                        if (true == turn) {
                            q = false;
                            while (true == turn) {}
                            q = true;
                        }
//                        System.out.println("p pre-protocol");
                    }

                    // critical section
                    temp = counter;
                    counter = temp + 1;

                    // post-protocol section
                    turn = true;
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
                        if (false == turn) {
                            p = false;
                            while (false == turn) {}
                            p = true;
                        }
 //                       System.out.println("q pre-protocol");
                    }

                    // critical section
                    temp = counter;
                    counter = temp + 1;

                    // post-protocol section
                    turn = false;
                    p = false;
                }
            }
        } catch (
                InterruptedException e) {

        }

    }

    public static void main(String[] args) {
        Attempt_05 p = new Attempt_05();
        Attempt_05 q = new Attempt_05();
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
