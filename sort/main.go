package main

import (
	"sort"
	"fmt"
)

var s []string
var per []Person

type Person struct{
	 Name string
	 Age int
}


// ByPerson implements the sort interface for the Person array
type ByPerson []Person
func (arr ByPerson) Len() int {return len(per)}
func (arr ByPerson) Swap(i, j int) {per[i], per[j] = per[j],per[i]}
func (arr ByPerson) Less(i, j int) bool { return per[i].Name < per[j].Name }

// ByName implements the sort interface for the name array
type ByName []string
func (arr ByName) Len() int {return len(s)}
func (arr ByName) Swap(i, j int) {s[i], s[j] = s[j],s[i]}
func (arr ByName) Less(i, j int) bool { return s[i] < s[j] }



func main() {
	s = []string {"Zeno","John","Al","Jenny"}
	per = []Person{{"Beatrice",20},{"Zack",30},{"Al",23},{"Rick",59},{"Aaron",90}}

	fmt.Println(s)
	sort.Sort(ByName(s))
	fmt.Println("Sorted: ",s)


	fmt.Println(per)
	sort.Sort(ByPerson(per))
	fmt.Println("Sorted: ",per)

}
