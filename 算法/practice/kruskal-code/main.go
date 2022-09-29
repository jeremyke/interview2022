package main

//思路：
//1. 把所有的边按照权从小到大排序生成数组
//2. 依次便历数组，把边串起来，如果形成了环，就跳过，不形成环就串起来。
//
//怎么判断环？
//1. 刚开始所有节点都各自是一个集合
//2. 判断第一条边的两个节点是否在一个集合，不在一个集合就串起来这条边，且把2好节点合并为同一个集合，在一个集合就是环了
//

type Graph struct {
	Nodes map[int]*Node
	Edges []*Edge
}
type Node struct {
	Value string
	In    int //入度
	Out   int //出度
	Nexts map[int]*Node
	Edges []*Edge
}
type Edge struct {
	Weight int   //权值：一般表示距离
	From   *Node //起始点
	To     *Node //指向点
}

func kruskal() {
	nodeHashMap := make(map[*Node]map[*Node]int, 1) //每个节点存放该节点的集合

}

func isSameSet(from, to *Node) bool {
	fromSet := nodeHashMap[from]
	toSet := nodeHashMap[to]
	return fromSet == toSet

}
func union(from, to *Node) {

}
