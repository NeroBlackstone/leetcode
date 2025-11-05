class RecentCounter {
  mpq: Array<number>;
  last: number;
  constructor() {
    this.mpq = [];
    this.last = 0;
  }

  ping(t: number): number {
    this.mpq.push(t);
    while (this.mpq[this.last]! < t - 3000) {
      this.last++;
    }
    return this.mpq.length - this.last;
  }
}

const rc = new RecentCounter();
rc.ping(1);
rc.ping();
