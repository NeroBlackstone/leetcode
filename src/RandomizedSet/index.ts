class RandomizedSet {
  // key:arr value,value:arr index
  valMap = new Map<number, number>();
  valArr: number[] = [];
  insert(val: number): boolean {
    if (this.valMap.has(val)) {
      return false;
    } else {
      this.valArr.push(val);
      this.valMap.set(val, this.valArr.length - 1);
      return true;
    }
  }

  remove(val: number): boolean {
    if (this.valMap.has(val)) {
      // element index
      const index = this.valMap.get(val)!;
      const lastElm = this.valArr.at(-1)!;

      // 注意这里要先覆盖再删，否则的话如果要删的是最后一个元素，map里会把这个元素加回来
      // 即，map要删的元素和要更新key的元素是同一个
      // map要做两件事，删指定元素，更新指定元素的value（entry）
      this.valArr[index] = lastElm;
      this.valMap.set(lastElm, index);

      this.valArr.pop();
      this.valMap.delete(val);

      return true;
    } else {
      return false;
    }
  }

  getRandom(): number {
    return this.valArr[Math.floor(Math.random() * this.valArr.length)]!;
  }
}
//https://leetcode.com/problems/insert-delete-getrandom-o1/description/
