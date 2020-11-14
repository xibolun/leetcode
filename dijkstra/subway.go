package dijkstra

import (
	"fmt"
	"strings"
)

//1，设计。
//请为北京地铁线路图，设计一下类模型。包括：
//地铁线，例如，1号线，2号线，13号线。要能够按某种顺序列出该线上的地铁站。
//地铁站，例如，西直门，王府井。要能够列出地铁站的名称、所在的地铁线（可能多个），以及该线上的相邻站点。
//​
//2，算法实现：
//设计并且实现一个算法，计算从站点A到站点B的最少换乘的次数。不需要输出具体的路径。

type Line struct {
	// 站点编号
	Num int `json:"num"`
	// 站点别名
	Alias    string  `json:"alias"`
	Stations []*Node `json:"stations"`
}

// Node 车站，即图的节点Node
type Node struct {
	// 站号
	Num int `json:"num"`
	// 所属线路
	LName string `json:"line"`
	// 站名
	Name string `json:"name"`
	Pre  *Node  `json:"-"`
	Next *Node  `json:"-"`
}

// Edges 图的边
type Edge struct {
	Start *Node
	End   *Node
	Cost  int
	LName string
	// 边的方向 -> 后节点 <- 前节点
	Direct string
}

// Graph 图，由节点和边组成
type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

const (
	// 正向
	Forward = "forward"
	// 反射
	Backward = "backward"
)

// AddLineStations 添加每条线路的节点至graph中
func (g *Graph) AddLineStations(line *Line) {
	nodes := line.Stations
	for i := range nodes {
		if i < len(nodes)-1 {
			g.AddEdge(nodes[i], nodes[i+1], nodes[i].LName, Forward)
		}
		g.AddNode(nodes[i])
	}

	// 地铁是双向的，反向添加图
	for j := len(nodes) - 1; j > 0; j-- {
		g.AddEdge(nodes[j], nodes[j-1], nodes[j].LName, Backward)
	}
}

// AddNode 添加节点
func (g *Graph) AddNode(station *Node) {
	noExist := true
	for _, item := range g.Nodes {
		// 认为在不同线路上的节点，不是同样的节点
		if item.Name == station.Name && item.LName == station.LName {
			noExist = false
		}
	}
	if noExist {
		g.Nodes = append(g.Nodes, station)
	}
}

// AddEdge 添加边
func (g *Graph) AddEdge(from, to *Node, lineName, direct string) {
	edge := &Edge{
		Start:  from,
		End:    to,
		LName:  lineName,
		Direct: direct,
	}
	g.Edges = append(g.Edges, edge)
}

// NewCostTable 构建节点Hash表，存储每个节点的距离
func (g *Graph) NewCostTable(startNode *Node) map[*Node][]*Node {
	costTable := make(map[*Node][]*Node)
	costTable[startNode] = []*Node{startNode}

	for _, node := range g.Nodes {
		if node.Name != startNode.Name {
			costTable[node] = nil
		}
	}
	return costTable
}

// GetNodeEdges 获取节点的所有的关联的边
func (g *Graph) GetNodeEdges(node *Node) (edges []*Edge) {
	for _, edge := range g.Edges {
		if edge.Start.Name == node.Name {
			edges = append(edges, edge)
		}
	}
	return edges
}

// GetNode 根据名称获取站点信息，包括所属线路列表，前置、后置节点
func (g *Graph) GetNode(name string) []*Node {
	res := make([]*Node, 0)
	for _, item := range g.Nodes {
		if item.Name != name {
			continue
		}
		node := *item
		edges := g.GetNodeEdges(item)
		for _, edge := range edges {
			if item.LName != edge.LName {
				continue
			}
			// 站点之间可以是双向的，添加一个指向
			if edge.Direct == Forward {
				node.Next = edge.End
			} else {
				node.Pre = edge.End
			}
		}
		res = append(res, &node)
	}
	return res
}

// GetSingleNode 根据名称获取站点信息
func (g *Graph) GetSingleNode(name string) *Node {
	for _, item := range g.Nodes {
		if item.Name != name {
			continue
		}
		return item
	}
	return nil
}

// ShortestPath 求出最短路径
func (g *Graph) ShortestPath(from, to string) ([]*Node, error) {
	startNode := g.GetSingleNode(from)
	if startNode == nil {
		return nil, fmt.Errorf("开始节点信息不存在,%s", from)
	}

	end := g.GetSingleNode(to)
	if end == nil {
		return nil, fmt.Errorf("结束节点信息不存在，%s", to)
	}

	nodeTable := g.NewCostTable(startNode)

	// 已经遍历过的节点
	var visited []*Node
	loopVisited(g, startNode, end, nodeTable, visited)

	// 获取最短路径
	var path []*Node
	for k, v := range nodeTable {
		if k.Name == to {
			if len(path) <= 0 {
				path = v
				continue
			}
			if getTransferTimes(v) <= getTransferTimes(path) {
				path = v
			}
		}
	}
	return path, nil
}

// MinTransferTimes 最小换乘次数
func (g *Graph) MinTransferTimes(from, to string) (int, error) {
	paths, err := g.MinTransfer(from, to)
	return getTransferTimes(paths), err
}

// MinTransfer 最小换乘
func (g *Graph) MinTransfer(from, to string) ([]*Node, error) {
	startNode := g.GetSingleNode(from)
	if startNode == nil {
		return nil, fmt.Errorf("开始节点信息不存在,%s", from)
	}

	end := g.GetSingleNode(to)
	if end == nil {
		return nil, fmt.Errorf("结束节点信息不存在，%s", to)
	}

	nodeTable := g.NewCostTable(startNode)

	// 已经遍历过的节点
	var visited []*Node
	transferLoopVisited(g, startNode, end, nodeTable, visited)

	// 获取最小换乘，此处可能会存在多个，因为节点在不同线路上会认为是不同的节点
	var path []*Node
	for k, v := range nodeTable {
		if k.Name == to {
			if len(path) <= 0 {
				path = v
				continue
			}
			if getTransferTimes(v) <= getTransferTimes(path) {
				path = v
			}
		}
	}
	return path, nil
}

// transferLoopVisited 循环访问各个节点，计算每个节点的路径
func transferLoopVisited(g *Graph, node, end *Node, nodeTable map[*Node][]*Node, visited []*Node) {
	if node == nil || len(visited) == len(g.Nodes) {
		return
	}
	visited = append(visited, node)
	nodeEdges := g.GetNodeEdges(node)

	// 加入线路权重
	if edges := addLineWeight(nodeEdges, end.LName); len(edges) > 0 {
		for _, edge := range edges {
			// child path
			paths := append(nodeTable[node], edge.End)
			// prevent cycle
			if checkIsCycle(paths) {
				continue
			}
			// 最小换乘
			if getTransferTimes(paths) < getTransferTimes(nodeTable[edge.End]) || nodeTable[edge.End] == nil {
				nodeTable[edge.End] = paths
				if paths[len(paths)-1].Name == end.Name {
					return
				}
				transferLoopVisited(g, edge.End, end, nodeTable, visited)
			}
		}
	} else {
		for _, edge := range nodeEdges {
			// child path
			paths := append(nodeTable[node], edge.End)
			// prevent cycle
			if checkIsCycle(paths) {
				continue
			}
			// 最小换乘，如果换乘次数小于现有的换乘次数，添加进来，并更新掉
			if getTransferTimes(paths) < getTransferTimes(nodeTable[edge.End]) || nodeTable[edge.End] == nil {
				nodeTable[edge.End] = paths
				if paths[len(paths)-1].Name == end.Name {
					return
				}
				transferLoopVisited(g, edge.End, end, nodeTable, visited)
			}
		}
	}
}

// checkIsCycle  校验方向，在同一线上的节点，不允许出现反向
func checkIsCycle(paths []*Node) bool {
	var isCycle bool
	for i := range paths {
		for j := range paths {
			if i == j {
				continue
			}
			if paths[i].Name != paths[j].Name {
				continue
			}

			if i < len(paths)-1 && j > 1 && paths[i+1].Name == paths[j-1].Name {
				isCycle = true
			}

			if i > 1 && j < len(paths)-1 && paths[i-1].Name == paths[j+1].Name {
				isCycle = true
			}
		}
	}
	return isCycle
}

// getTransferTimes 获取换乘次数，用于比较
func getTransferTimes(paths []*Node) int {
	if len(paths) <= 0 {
		return 0
	}

	var sum int
	tmpL := paths[0].LName
	for i := range paths {
		if paths[i].LName == tmpL {
			continue
		}
		tmpL = paths[i].LName
		sum++
	}
	return sum
}

// loopVisited 循环访问各个节点，计算每个节点的路径
func loopVisited(g *Graph, node, end *Node, nodeTable map[*Node][]*Node, visited []*Node) {
	if node == nil || len(visited) == len(g.Nodes) {
		return
	}
	visited = append(visited, node)
	nodeEdges := g.GetNodeEdges(node)

	// 加入线路权重
	if edges := addLineWeight(nodeEdges, end.LName); len(edges) > 0 {
		for _, edge := range edges {
			// child path
			paths := append(nodeTable[node], edge.End)
			if checkIsCycle(paths) {
				continue
			}
			// 最少的站
			if len(paths) < len(nodeTable[edge.End]) || nodeTable[edge.End] == nil {
				nodeTable[edge.End] = paths
				if paths[len(paths)-1].Name == end.Name {
					return
				}

				loopVisited(g, edge.End, end, nodeTable, visited)
			}
		}
	} else {
		for _, edge := range nodeEdges {
			// child path
			paths := append(nodeTable[node], edge.End)
			if checkIsCycle(paths) {
				continue
			}
			// 最少的站
			if len(paths) < len(nodeTable[edge.End]) || nodeTable[edge.End] == nil {
				nodeTable[edge.End] = paths
				if paths[len(paths)-1].Name == end.Name {
					return
				}
				loopVisited(g, edge.End, end, nodeTable, visited)
			}
		}
	}

}

// addLineWeight 添加线路权重
// edge end node 和 目的节点在同一线路上时，会有两种方向
// 添加一个权重处理，同一线路上的节点优先做进行寻路
func addLineWeight(edges []*Edge, lName string) []*Edge {
	w := make([]*Edge, 0)
	for i := range edges {
		if edges[i].LName == lName {
			w = append(w, edges[i])
		}
	}
	return w
}

// printPath 打印路径链路
func printPath(paths []*Node) string {
	sb := strings.Builder{}
	for i := range paths {
		sb.WriteString(fmt.Sprintf("%d:%s(%s) -> ", i, paths[i].Name, paths[i].LName))
	}
	return fmt.Sprintf("%s end\n", sb.String())
}
