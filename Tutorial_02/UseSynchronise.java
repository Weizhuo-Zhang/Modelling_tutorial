class P extends Thread
{
  //the shared synch instance
  protected Synchronise s;

  public P(Synchronise s)
  {
    this.s = s;
  }

  public void run()
  {
    while (true) {
        synchronized(s){
            task1p();
            s.synch();
            task2p();
        }
    }
  }

  private void task1p()
  {
    System.out.println("1p");
  }

  private void task2p()
  {
    System.out.println("2p");
  }
}

class Q extends Thread
{
  //the shared synch instance
  protected Synchronise s;

  public Q(Synchronise s)
  {
    this.s = s;
  }

  public void run()
  {
    while (true) {
        synchronized(s){
            task1q();
            s.synch();
            task2q();
        }
    }
  }

  private void task1q()
  {
    System.out.println("1q");
  }

  private void task2q()
  {
    System.out.println("2q");
  }
}

class Synchronise
{
  // any useful variables go here

  public synchronized void synch()
  {
    // the code to synchronise goes here
  }
}

class UseSynchronise
{
  public static void main(String [] args)
  {
    Synchronise s = new Synchronise();
    Thread p = new P(s);
    Thread q = new Q(s);
    p.start();
    q.start();
  }
}
