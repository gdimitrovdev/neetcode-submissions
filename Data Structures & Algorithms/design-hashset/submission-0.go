type MyHashSet struct {
	l []int
}

func Constructor() MyHashSet {
    return MyHashSet{l: []int{}}
}

func (this *MyHashSet) Add(key int) {
    if this.Contains(key) {
		return
	}

	this.l = append(this.l, key)
}

func (this *MyHashSet) Remove(key int) {
	for i := 0; i < len(this.l); i++ {
		if this.l[i] == key {
			this.l = append(this.l[:i], this.l[i+1:]...)
			return
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
    for i := 0; i < len(this.l); i++ {
		if this.l[i] == key {
			return true
		}
	}

	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 