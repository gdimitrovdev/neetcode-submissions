type MyHashMap struct {
	data [][2]int
}

func Constructor() MyHashMap {
    return MyHashMap{data: [][2]int{}}
}

func (this *MyHashMap) Put(key int, value int) {
    for i, duo := range this.data {
		if duo[0] == key {
			this.data[i][1] = value
			return
		}
	}

	this.data = append(this.data, [2]int{key, value})
}

func (this *MyHashMap) Get(key int) int {
    for _, duo := range this.data {
		if duo[0] == key {
			return duo[1]
		}
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
    for i, duo := range this.data {
		if duo[0] == key {
			this.data = append(this.data[:i], this.data[i+1:]...)
			return
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */