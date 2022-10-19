package main

type trieNode struct {
	Pass  int         //经过该节点的次数
	End   int         //该节点为末尾节点的次数
	Nexts []*trieNode //每个字母的线，如果有a线，thread[1]!= null，有就thread[1] == null
}

var root = &trieNode{Nexts: make([]*trieNode, 32)}

func insert(word string) {
	if word == "" {
		return
	}
	node := &trieNode{Nexts: make([]*trieNode, 32)}
	wordArr := []byte(word)
	node.Pass++
	index := 0
	for i := 0; i < len(wordArr); i++ {
		index = int(wordArr[i] - 'a') //根据字符串,决定走哪条路径
		if node.Nexts[index] == nil {
			node.Nexts[index] = &trieNode{Nexts: make([]*trieNode, 32)}
		}
		node = node.Nexts[index] //当前节点移动到下一个节点
		node.Pass++
	}

}

//查询word之前加过几次
func search(word string) int {
	if word == "" {
		return 0
	}
	node := &trieNode{Nexts: make([]*trieNode, 32)}
	wordArr := []byte(word)
	node.Pass++
	index := 0
	for i := 0; i < len(wordArr); i++ {
		index = int(wordArr[i] - 'a') //根据字符串,决定走哪条路径
		if node.Nexts[index] == nil {
			return 0
		}
		node = node.Nexts[index] //当前节点移动到下一个节点
	}
	return node.End
}

//所有加入的字符串，有几个是以pre作为前缀的
func prefixNum(pre string) int {
	if pre == "" {
		return 0
	}
	node := root
	preArr := []byte(pre)
	node.Pass++
	index := 0
	for i := 0; i < len(preArr); i++ {
		index = int(preArr[i] - 'a') //根据字符串,决定走哪条路径
		if node.Nexts[index] == nil {
			return 0
		}
		node = node.Nexts[index] //当前节点移动到下一个节点
	}
	return node.Pass
}

func deleted(word string) {
	if search(word) > 0 { //确定加入过word这个字符串
		wordArr := []byte(word)
		node := root
		node.Pass--
		index := 0
		for i := 0; i < len(wordArr); i++ {
			index = int(wordArr[i] - 'a')    //根据字符串,决定走哪条路径
			if node.Nexts[index].Pass == 1 { //当前节点的下一个节点的pass值为1时，整个下节点置空
				node.Nexts[index] = nil
				return
			}
			node = node.Nexts[index] //当前节点移动到下一个节点
			node.Pass--              //沿途pass值减1
		}
		node.End--
	}
}

func main() {
	insert("abc")
}
