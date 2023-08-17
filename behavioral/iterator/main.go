package main

import "fmt"

type UserProfile struct {
	ID   int
	Name string
}

// Aggregate interface
type UserProfileCollection interface {
	CreateIterator() Iterator
}

// Iterator interface
type Iterator interface {
	HasNext() bool
	Next() *UserProfile
}

// Concrete Aggregate: Slice
type UserProfileSlice struct {
	profiles []UserProfile
}

func (us *UserProfileSlice) CreateIterator() Iterator {
	return &SliceIterator{profiles: us.profiles, index: 0}
}

type SliceIterator struct {
	profiles []UserProfile
	index    int
}

func (si *SliceIterator) HasNext() bool {
	return si.index < len(si.profiles)
}

func (si *SliceIterator) Next() *UserProfile {
	if si.HasNext() {
		profile := &si.profiles[si.index]
		si.index++
		return profile
	}
	return nil
}

// Concrete Aggregate: Map
type UserProfileMap struct {
	profiles map[int]UserProfile
	keys     []int
}

func (um *UserProfileMap) CreateIterator() Iterator {
	keys := make([]int, 0, len(um.profiles))
	for k := range um.profiles {
		keys = append(keys, k)
	}
	return &MapIterator{profiles: um.profiles, keys: keys, index: 0}
}

type MapIterator struct {
	profiles map[int]UserProfile
	keys     []int
	index    int
}

func (mi *MapIterator) HasNext() bool {
	return mi.index < len(mi.keys)
}

func (mi *MapIterator) Next() *UserProfile {
	if mi.HasNext() {
		key := mi.keys[mi.index]
		profile := mi.profiles[key]
		mi.index++
		return &profile
	}
	return nil
}

func main() {
	sliceData := &UserProfileSlice{
		profiles: []UserProfile{
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bob"},
		},
	}

	mapData := &UserProfileMap{
		profiles: map[int]UserProfile{
			3: {ID: 3, Name: "Charlie"},
			4: {ID: 4, Name: "Dave"},
		},
	}

	// Iterating over slice
	iterator := sliceData.CreateIterator()
	fmt.Println("Iterating over slice:")
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}

	// Iterating over map
	iterator = mapData.CreateIterator()
	fmt.Println("\nIterating over map:")
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}
