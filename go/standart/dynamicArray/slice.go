package dynamicArray

import (
	"errors"
	"fmt"
)

type DynamicArray struct {
	items []int
	size  int
}

func NewDynamicArray() *DynamicArray {
	return &DynamicArray{
		items: []int{},
	}
}

func (d *DynamicArray) Print() string {
	result := ""
	for i := 0; i < d.size; i++ {
		result += fmt.Sprintf("%d ", d.items[i])
	}
	return result
}

func (d *DynamicArray) GetSize() int {
	return d.size
}

func (d *DynamicArray) Get(index int) (int, error) {
	if index > d.size || index < 0 {
		return 0, errors.New("index out of range")
	}
	return d.items[index], nil
}

func (d *DynamicArray) Find(key int) int {
	for i := 0; i < d.size; i++ {
		if d.items[i] == key {
			return i
		}
	}
	return -1
}

func (d *DynamicArray) FindLast(key int) int {
	for i := d.size - 1; i > 0; i-- {
		if d.items[i] == key {
			return i
		}
	}
	return -1
}

func (d *DynamicArray) IncreaseArray() {
	newCount := d.size * 2
	if d.size == 0 {
		newCount = 4
	}
	newArray := make([]int, newCount)

	for i := 0; i < d.size; i++ {
		newArray[i] = d.items[i]
	}
	d.items = newArray
}

func (d *DynamicArray) PushBack(item int) {
	if d.size == len(d.items) {
		d.IncreaseArray()
	}
	d.items[d.size] = item
	d.size++
}

func (d *DynamicArray) Insert(index int, items int) {
	if index < 0 || index > d.size {
		panic("index вышел за пределы массива")
	}
	if d.size == len(d.items) {
		d.IncreaseArray()
	}

	for i := d.size - 1; i >= index; i-- {
		d.items[i+1] = d.items[i]
	}
	d.items[index] = items
	d.size++
}

func (d *DynamicArray) PushFront(item int) {
	d.Insert(0, item)
}

func (d *DynamicArray) PushBackRange(arr DynamicArray) {
	for i := 0; i < arr.size; i++ {
		d.PushBack(arr.items[i])
	}
}

func (d *DynamicArray) InsertRange(index int, arr DynamicArray) {
	if index < 0 || index > d.size {
		panic("вышли за пределы")
	}

	if d.size+arr.size > len(d.items) {
		d.IncreaseArray()
	}

	for i := 0; i < arr.size; i++ {
		d.Insert(index+i+1, arr.items[i])
	}
}

func (d *DynamicArray) PopBack() error {
	if d.GetSize() == 0 {
		return errors.New("array is empty")
	}
	d.size--
	return nil
}

func (d *DynamicArray) RemoveByIndex(index int) error {
	if d.GetSize() == 0 {
		return errors.New("array is empty")
	}

	if index < 0 || index >= d.GetSize() {
		return errors.New("index out of range")
	}

	for i := index; i < d.GetSize(); i++ {
		d.items[i-1] = d.items[i]
	}

	d.size--
	return nil
}

func (d *DynamicArray) PopFront() error {
	return d.RemoveByIndex(0)
}

func (d *DynamicArray) Remove(item int) bool {
	ind := d.Find(item)
	if err := d.RemoveByIndex(ind); err != nil {
		return false
	}
	return true
}

func (d *DynamicArray) RemoveAll(item int) int {
	var count int
	var ind int

	for ind != -1 {
		ind = d.Find(item)
		if err := d.RemoveByIndex(ind); err != nil {
			return count
		}
		count++
	}
	return count
}

func main() {
	var arr = NewDynamicArray()
	arr.items = []int{1, 2, 3, 4, 5, 5, 5}
	arr.size = len(arr.items)
	fmt.Println(arr.size)

	fmt.Println(arr.RemoveAll(5))

}
