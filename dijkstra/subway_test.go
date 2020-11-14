package dijkstra

import (
	"encoding/json"
	"sort"
	"testing"
)

const (
	// 北京地铁数据信息
	BJSubwayData = `[{"num":10,"alias":"1号线","stations":[{"num":0,"line":"1号线","name":"苹果园"},{"num":1,"line":"1号线","name":"古城"},{"num":2,"line":"1号线","name":"八角游乐园"},{"num":3,"line":"1号线","name":"八宝山"},{"num":4,"line":"1号线","name":"玉泉路"},{"num":5,"line":"1号线","name":"五棵松"},{"num":6,"line":"1号线","name":"万寿路"},{"num":7,"line":"1号线","name":"公主坟"},{"num":8,"line":"1号线","name":"军事博物馆"},{"num":9,"line":"1号线","name":"木樨路"},{"num":10,"line":"1号线","name":"南礼士路"},{"num":11,"line":"1号线","name":"复兴门"},{"num":12,"line":"1号线","name":"西单"},{"num":13,"line":"1号线","name":"天安门西"},{"num":14,"line":"1号线","name":"天安门东"},{"num":15,"line":"1号线","name":"王府井"},{"num":16,"line":"1号线","name":"东单"},{"num":17,"line":"1号线","name":"建国门"},{"num":18,"line":"1号线","name":"永安里"},{"num":19,"line":"1号线","name":"国贸"},{"num":20,"line":"1号线","name":"大望路"},{"num":21,"line":"1号线","name":"四惠"},{"num":22,"line":"1号线","name":"四惠东"}]},{"num":20,"alias":"2号线","stations":[{"num":0,"line":"2号线","name":"西直门"},{"num":1,"line":"2号线","name":"积水潭"},{"num":2,"line":"2号线","name":"鼓楼大街"},{"num":3,"line":"2号线","name":"安定门"},{"num":4,"line":"2号线","name":"雍和宫"},{"num":5,"line":"2号线","name":"东直门"},{"num":6,"line":"2号线","name":"东四十条"},{"num":7,"line":"2号线","name":"朝阳门"},{"num":8,"line":"2号线","name":"建国门"},{"num":9,"line":"2号线","name":"北京站"},{"num":10,"line":"2号线","name":"崇文门"},{"num":11,"line":"2号线","name":"前门"},{"num":12,"line":"2号线","name":"和平门"},{"num":13,"line":"2号线","name":"宣武门"},{"num":14,"line":"2号线","name":"长椿街"},{"num":15,"line":"2号线","name":"复兴门"},{"num":16,"line":"2号线","name":"阜成门"},{"num":17,"line":"2号线","name":"车公庄"}]},{"num":40,"alias":"4号线","stations":[{"num":0,"line":"4号线","name":"安河桥北"},{"num":1,"line":"4号线","name":"北宫门"},{"num":2,"line":"4号线","name":"西苑"},{"num":3,"line":"4号线","name":"圆明园"},{"num":4,"line":"4号线","name":"北京大学东门"},{"num":5,"line":"4号线","name":"中关村"},{"num":6,"line":"4号线","name":"海淀黄庄"},{"num":7,"line":"4号线","name":"人民大学"},{"num":8,"line":"4号线","name":"魏公村"},{"num":9,"line":"4号线","name":"国家图书馆"},{"num":10,"line":"4号线","name":"动物园"},{"num":11,"line":"4号线","name":"西直门"},{"num":12,"line":"4号线","name":"新街口"},{"num":13,"line":"4号线","name":"平安里"},{"num":14,"line":"4号线","name":"西四"},{"num":15,"line":"4号线","name":"灵境胡同"},{"num":16,"line":"4号线","name":"西单"},{"num":17,"line":"4号线","name":"宣武门"},{"num":18,"line":"4号线","name":"菜市口"},{"num":19,"line":"4号线","name":"陶然亭"},{"num":20,"line":"4号线","name":"北京南站"},{"num":21,"line":"4号线","name":"马家堡"},{"num":22,"line":"4号线","name":"角门西"},{"num":23,"line":"4号线","name":"公益西桥"},{"num":24,"line":"4号线","name":"新宫"},{"num":25,"line":"4号线","name":"西红门"},{"num":26,"line":"4号线","name":"高米店北"},{"num":27,"line":"4号线","name":"高米店南"},{"num":28,"line":"4号线","name":"枣园"},{"num":29,"line":"4号线","name":"清源路"},{"num":30,"line":"4号线","name":"黄村西大街"},{"num":31,"line":"4号线","name":"黄村火车站"},{"num":32,"line":"4号线","name":"义和庄"},{"num":33,"line":"4号线","name":"生物医药基地"},{"num":34,"line":"4号线","name":"天宫院"}]},{"num":50,"alias":"5号线","stations":[{"num":0,"line":"5号线","name":"天通苑北"},{"num":1,"line":"5号线","name":"天通苑"},{"num":2,"line":"5号线","name":"天通苑南"},{"num":3,"line":"5号线","name":"立水桥"},{"num":4,"line":"5号线","name":"立水桥南"},{"num":5,"line":"5号线","name":"北苑路北"},{"num":6,"line":"5号线","name":"大屯东路"},{"num":7,"line":"5号线","name":"惠新四街北口"},{"num":8,"line":"5号线","name":"惠新四街南口"},{"num":9,"line":"5号线","name":"和平西桥"},{"num":10,"line":"5号线","name":"和平里北街"},{"num":11,"line":"5号线","name":"雍和宫"},{"num":12,"line":"5号线","name":"北新桥"},{"num":13,"line":"5号线","name":"张自忠路"},{"num":14,"line":"5号线","name":"东四"},{"num":15,"line":"5号线","name":"灯市口"},{"num":16,"line":"5号线","name":"东单"},{"num":17,"line":"5号线","name":"崇文门"},{"num":18,"line":"5号线","name":"磁器口"},{"num":19,"line":"5号线","name":"天坛东门"},{"num":20,"line":"5号线","name":"蒲黄榆"},{"num":21,"line":"5号线","name":"刘家窑"},{"num":22,"line":"5号线","name":"宋家庄"}]},{"num":60,"alias":"6号线","stations":[{"num":0,"line":"6号线","name":"海淀五路居"},{"num":1,"line":"6号线","name":"慈寿寺"},{"num":2,"line":"6号线","name":"花园桥"},{"num":3,"line":"6号线","name":"白石桥南"},{"num":4,"line":"6号线","name":"车公庄西"},{"num":5,"line":"6号线","name":"车公庄"},{"num":6,"line":"6号线","name":"平安里"},{"num":7,"line":"6号线","name":"北海北"},{"num":8,"line":"6号线","name":"南锣鼓巷"},{"num":9,"line":"6号线","name":"东四"},{"num":10,"line":"6号线","name":"朝阳门"},{"num":11,"line":"6号线","name":"东大桥"},{"num":12,"line":"6号线","name":"呼家楼"},{"num":13,"line":"6号线","name":"金台路"},{"num":14,"line":"6号线","name":"十里堡"},{"num":15,"line":"6号线","name":"青年路"},{"num":16,"line":"6号线","name":"褡裢坡"},{"num":17,"line":"6号线","name":"黄渠"},{"num":18,"line":"6号线","name":"常营"},{"num":19,"line":"6号线","name":"草房"},{"num":20,"line":"6号线","name":"物资学院路"},{"num":21,"line":"6号线","name":"通州北关"},{"num":22,"line":"6号线","name":"北运河西"},{"num":23,"line":"6号线","name":"郝家府"},{"num":24,"line":"6号线","name":"东夏园"},{"num":25,"line":"6号线","name":"潞城"}]},{"num":70,"alias":"7号线","stations":[{"num":0,"line":"7号线","name":"北京西站"},{"num":1,"line":"7号线","name":"湾子"},{"num":2,"line":"7号线","name":"达官营"},{"num":3,"line":"7号线","name":"广安门内"},{"num":4,"line":"7号线","name":"菜市口"},{"num":5,"line":"7号线","name":"虎坊桥"},{"num":6,"line":"7号线","name":"珠市口"},{"num":7,"line":"7号线","name":"桥湾"},{"num":8,"line":"7号线","name":"磁器口"},{"num":9,"line":"7号线","name":"广渠门内"},{"num":10,"line":"7号线","name":"广渠门外"},{"num":11,"line":"7号线","name":"九龙山"},{"num":12,"line":"7号线","name":"大郊亭"},{"num":13,"line":"7号线","name":"百子湾"},{"num":14,"line":"7号线","name":"化工"},{"num":15,"line":"7号线","name":"南楼梓庄"},{"num":16,"line":"7号线","name":"欢乐谷景区"},{"num":17,"line":"7号线","name":"双合"},{"num":18,"line":"7号线","name":"焦化厂"}]},{"num":80,"alias":"8号线","stations":[{"num":0,"line":"8号线","name":"朱辛庄"},{"num":1,"line":"8号线","name":"育知路"},{"num":2,"line":"8号线","name":"平西府"},{"num":3,"line":"8号线","name":"回龙观东大街"},{"num":4,"line":"8号线","name":"霍营"},{"num":5,"line":"8号线","name":"育新"},{"num":6,"line":"8号线","name":"西小口"},{"num":7,"line":"8号线","name":"永泰庄"},{"num":8,"line":"8号线","name":"林萃桥"},{"num":9,"line":"8号线","name":"森林公园南门"},{"num":10,"line":"8号线","name":"奥林匹克公园"},{"num":11,"line":"8号线","name":"奥林中心"},{"num":12,"line":"8号线","name":"北土城"},{"num":13,"line":"8号线","name":"安华桥"},{"num":14,"line":"8号线","name":"安德里北街"},{"num":15,"line":"8号线","name":"鼓楼大街"},{"num":16,"line":"8号线","name":"什刹海"},{"num":17,"line":"8号线","name":"南锣鼓巷"}]},{"num":90,"alias":"9号线","stations":[{"num":0,"line":"9号线","name":"国家图书馆"},{"num":1,"line":"9号线","name":"白石桥南"},{"num":2,"line":"9号线","name":"白堆子"},{"num":3,"line":"9号线","name":"军事博物馆"},{"num":4,"line":"9号线","name":"北京西站"},{"num":5,"line":"9号线","name":"六里桥东"},{"num":6,"line":"9号线","name":"六里桥"},{"num":7,"line":"9号线","name":"七里庄"},{"num":8,"line":"9号线","name":"丰台东大街"},{"num":9,"line":"9号线","name":"丰台南路"},{"num":10,"line":"9号线","name":"科怡路"},{"num":11,"line":"9号线","name":"丰台科技园"},{"num":12,"line":"9号线","name":"郭公庄"}]},{"num":100,"alias":"10号线","stations":[{"num":0,"line":"10号线","name":"巴沟"},{"num":1,"line":"10号线","name":"苏州街"},{"num":2,"line":"10号线","name":"海淀黄庄"},{"num":3,"line":"10号线","name":"知春里"},{"num":4,"line":"10号线","name":"知春路"},{"num":5,"line":"10号线","name":"西土城"},{"num":6,"line":"10号线","name":"牡丹园"},{"num":7,"line":"10号线","name":"健德门"},{"num":8,"line":"10号线","name":"北土城"},{"num":9,"line":"10号线","name":"安贞门"},{"num":10,"line":"10号线","name":"惠新西街南口"},{"num":11,"line":"10号线","name":"芍药居"},{"num":12,"line":"10号线","name":"太阳宫"},{"num":13,"line":"10号线","name":"三元桥"},{"num":14,"line":"10号线","name":"亮马桥"},{"num":15,"line":"10号线","name":"农业展览馆"},{"num":16,"line":"10号线","name":"团结湖"},{"num":17,"line":"10号线","name":"呼家楼"},{"num":18,"line":"10号线","name":"金台夕照"},{"num":19,"line":"10号线","name":"国贸"},{"num":20,"line":"10号线","name":"双井"},{"num":21,"line":"10号线","name":"劲松"},{"num":22,"line":"10号线","name":"潘家园"},{"num":23,"line":"10号线","name":"十里河"},{"num":24,"line":"10号线","name":"分钟寺"},{"num":25,"line":"10号线","name":"成寿寺"},{"num":26,"line":"10号线","name":"宋家庄"},{"num":27,"line":"10号线","name":"石榴庄"},{"num":28,"line":"10号线","name":"大红门"},{"num":29,"line":"10号线","name":"角门东"},{"num":30,"line":"10号线","name":"角门西"},{"num":31,"line":"10号线","name":"草桥"},{"num":32,"line":"10号线","name":"季家庙"},{"num":33,"line":"10号线","name":"首经贸"},{"num":34,"line":"10号线","name":"丰台站"},{"num":35,"line":"10号线","name":"泥洼"},{"num":36,"line":"10号线","name":"西局"},{"num":37,"line":"10号线","name":"六里桥"},{"num":38,"line":"10号线","name":"莲花桥"},{"num":39,"line":"10号线","name":"公主坟"},{"num":40,"line":"10号线","name":"西钓鱼台"},{"num":41,"line":"10号线","name":"慈寿寺"},{"num":42,"line":"10号线","name":"车道沟"},{"num":43,"line":"10号线","name":"长春桥"},{"num":44,"line":"10号线","name":"火磁营"}]},{"num":130,"alias":"13号线","stations":[{"num":0,"line":"13号线","name":"西直门"},{"num":1,"line":"13号线","name":"大钟寺"},{"num":2,"line":"13号线","name":"知春路"},{"num":3,"line":"13号线","name":"五道口"},{"num":4,"line":"13号线","name":"上地"},{"num":5,"line":"13号线","name":"西二旗"},{"num":6,"line":"13号线","name":"龙泽"},{"num":7,"line":"13号线","name":"回龙观"},{"num":8,"line":"13号线","name":"霍营"},{"num":9,"line":"13号线","name":"立水桥"},{"num":10,"line":"13号线","name":"北苑"},{"num":11,"line":"13号线","name":"望京西"},{"num":12,"line":"13号线","name":"芍药居"},{"num":13,"line":"13号线","name":"光熙门"},{"num":14,"line":"13号线","name":"柳芳"},{"num":15,"line":"13号线","name":"东直门"}]},{"num":140,"alias":"14号线东段","stations":[{"num":0,"line":"14号线东段","name":"善各庄"},{"num":1,"line":"14号线东段","name":"来广营"},{"num":2,"line":"14号线东段","name":"东湖渠"},{"num":3,"line":"14号线东段","name":"望京"},{"num":4,"line":"14号线东段","name":"阜通"},{"num":5,"line":"14号线东段","name":"望京南"},{"num":6,"line":"14号线东段","name":"将台"},{"num":7,"line":"14号线东段","name":"东风北桥"},{"num":8,"line":"14号线东段","name":"枣营"},{"num":9,"line":"14号线东段","name":"金台路"},{"num":10,"line":"14号线东段","name":"大望路"},{"num":11,"line":"14号线东段","name":"九龙山"},{"num":12,"line":"14号线东段","name":"北工大西门"},{"num":13,"line":"14号线东段","name":"十里河"},{"num":14,"line":"14号线东段","name":"方庄"},{"num":15,"line":"14号线东段","name":"蒲黄榆"},{"num":16,"line":"14号线东段","name":"景泰"},{"num":17,"line":"14号线东段","name":"永定门外"},{"num":18,"line":"14号线东段","name":"北京南站"}]},{"num":141,"alias":"14号线西段","stations":[{"num":0,"line":"14号线西段","name":"张郭庄"},{"num":1,"line":"14号线西段","name":"园博园"},{"num":2,"line":"14号线西段","name":"大瓦窑"},{"num":3,"line":"14号线西段","name":"郭庄子"},{"num":4,"line":"14号线西段","name":"大井"},{"num":5,"line":"14号线西段","name":"七里庄"},{"num":6,"line":"14号线西段","name":"西局"}]},{"num":150,"alias":"15号线","stations":[{"num":0,"line":"15号线","name":"清华东路西口"},{"num":1,"line":"15号线","name":"六道口"},{"num":2,"line":"15号线","name":"北沙滩"},{"num":3,"line":"15号线","name":"奥林匹克公园"},{"num":4,"line":"15号线","name":"安立路"},{"num":5,"line":"15号线","name":"大屯路东"},{"num":6,"line":"15号线","name":"关庄"},{"num":7,"line":"15号线","name":"望京西"},{"num":8,"line":"15号线","name":"望京"},{"num":9,"line":"15号线","name":"望京东"},{"num":10,"line":"15号线","name":"崔各庄"},{"num":11,"line":"15号线","name":"马泉营"},{"num":12,"line":"15号线","name":"孙河"},{"num":13,"line":"15号线","name":"国展"},{"num":14,"line":"15号线","name":"花梨坎"},{"num":15,"line":"15号线","name":"后沙峪"},{"num":16,"line":"15号线","name":"南法信"},{"num":17,"line":"15号线","name":"石门"},{"num":18,"line":"15号线","name":"顺义"},{"num":19,"line":"15号线","name":"俸伯"}]},{"num":160,"alias":"八通线","stations":[{"num":0,"line":"八通线","name":"四惠"},{"num":1,"line":"八通线","name":"四惠东"},{"num":2,"line":"八通线","name":"高碑店"},{"num":3,"line":"八通线","name":"传媒大学"},{"num":4,"line":"八通线","name":"双桥"},{"num":5,"line":"八通线","name":"管庄"},{"num":6,"line":"八通线","name":"八里桥"},{"num":7,"line":"八通线","name":"通州北苑"},{"num":8,"line":"八通线","name":"果园"},{"num":9,"line":"八通线","name":"九棵树"},{"num":10,"line":"八通线","name":"梨园"},{"num":11,"line":"八通线","name":"临河里"},{"num":12,"line":"八通线","name":"土桥"}]},{"num":170,"alias":"昌平线","stations":[{"num":0,"line":"昌平线","name":"昌平西山口"},{"num":1,"line":"昌平线","name":"十三陵景区"},{"num":2,"line":"昌平线","name":"昌平"},{"num":3,"line":"昌平线","name":"昌平东关"},{"num":4,"line":"昌平线","name":"北邵洼"},{"num":5,"line":"昌平线","name":"南邵"},{"num":6,"line":"昌平线","name":"沙河高教城"},{"num":7,"line":"昌平线","name":"沙河"},{"num":8,"line":"昌平线","name":"巩华城"},{"num":9,"line":"昌平线","name":"朱辛庄"},{"num":10,"line":"昌平线","name":"生命科学院"},{"num":11,"line":"昌平线","name":"西二旗"}]},{"num":180,"alias":"亦庄线","stations":[{"num":0,"line":"亦庄线","name":"宋家庄"},{"num":1,"line":"亦庄线","name":"肖村"},{"num":2,"line":"亦庄线","name":"小红门"},{"num":3,"line":"亦庄线","name":"旧宫"},{"num":4,"line":"亦庄线","name":"亦庄桥"},{"num":5,"line":"亦庄线","name":"亦庄文化园"},{"num":6,"line":"亦庄线","name":"万源街"},{"num":7,"line":"亦庄线","name":"荣京东街"},{"num":8,"line":"亦庄线","name":"荣昌东街"},{"num":9,"line":"亦庄线","name":"同济南路"},{"num":10,"line":"亦庄线","name":"经海路"},{"num":11,"line":"亦庄线","name":"次渠南"},{"num":12,"line":"亦庄线","name":"次渠"}]},{"num":190,"alias":"房山线","stations":[{"num":0,"line":"房山线","name":"郭公庄"},{"num":1,"line":"房山线","name":"大葆台"},{"num":2,"line":"房山线","name":"稻田"},{"num":3,"line":"房山线","name":"长阳"},{"num":4,"line":"房山线","name":"篱笆房"},{"num":5,"line":"房山线","name":"广阳城"},{"num":6,"line":"房山线","name":"良乡大学城北"},{"num":7,"line":"房山线","name":"良乡大学城"},{"num":8,"line":"房山线","name":"良乡大学城西"},{"num":9,"line":"房山线","name":"良乡南关"},{"num":10,"line":"房山线","name":"苏庄"}]},{"num":200,"alias":"机场线","stations":[{"num":0,"line":"机场线","name":"东直门"},{"num":1,"line":"机场线","name":"三元桥"},{"num":2,"line":"机场线","name":"3号航站楼"},{"num":3,"line":"机场线","name":"2号航站楼"}]}]`
)

// 全局的图
var graph *Graph
// 所有线路
var allLine []*Line

// init 初始化数据信息
func init() {
	allLine = make([]*Line, 0)

	data := BJSubwayData
	_ = json.Unmarshal([]byte(data), &allLine)

	// 初始化图数据
	graph = &Graph{}
	for _, item := range allLine {
		graph.AddLineStations(item)
	}
}

// TestSortNodes 测试排序
func TestSortNodes(t *testing.T) {
	for _, line := range allLine {
		stations := line.Stations
		sort.Slice(stations, func(i, j int) bool {
			return stations[i].Num < stations[j].Num
		})
	}
}

// TestGetNodes 测试获取节点详情信息
func TestGetNodes(t *testing.T) {
	var source = []struct {
		name      string   // 站点名称
		lines     []string // 线路名称
		neighbors []string // 邻近节点
	}{
		{"天宫院", []string{"4号线", "7号线"}, []string{"生物医药基地", "AA"}},
		{"天宫院", []string{"4号线"}, []string{"生物医药基地"}},
	}

	for _, item := range source {
		name, lines, neighbors := item.name, item.lines, item.neighbors

		var actualLines []string
		var actualNeighbors []string
		for _, node := range graph.GetNode(name) {

			// 添加线路
			lexist := false
			for i := range actualLines {
				if node.LName == actualLines[i] {
					lexist = true
				}
			}
			if !lexist {
				actualLines = append(actualLines, node.LName)
			}

			// 添加邻近节点
			preExits := false
			nextExits := false
			for i := range actualNeighbors {
				if node.Next != nil {
					if node.Next.Name == actualNeighbors[i] {
						nextExits = true
					}
				}

				if node.Pre != nil {
					if node.Pre.Name == actualNeighbors[i] {
						preExits = true
					}
				}
			}

			if !preExits && node.Pre != nil {
				actualNeighbors = append(actualNeighbors, node.Pre.Name)
			}
			if !nextExits && node.Next != nil {
				actualNeighbors = append(actualNeighbors, node.Next.Name)
			}
		}

		//if len(actualLines) != len(lines) || len(neighbors) != len(actualNeighbors) {
		//	t.Errorf("(%s) actual line num is %d, expected: %d, neighbor length is %d, expected: %d", name, len(actualLines), len(lines), len(actualNeighbors), len(neighbors))
		//}

		// check line
		isSame := true
		for i := range lines {
			var exist bool
			for j := range actualLines {
				if lines[i] == actualLines[j] {
					exist = true
				}
			}
			if !exist {
				isSame = false
			}
		}
		if !isSame {
			t.Errorf("(%s) actual lines  is %v, expected: %v", name, actualLines, lines)
		}

		isSame = true
		// check neighbor
		for i := range neighbors {
			var exist bool
			for j := range actualNeighbors {
				if neighbors[i] == actualNeighbors[j] {
					exist = true
				}
			}
			if !exist {
				isSame = false
			}
		}
		if !isSame {
			t.Errorf("(%s) actual neighbors  is %v, expected: %v", name, actualNeighbors, neighbors)
		}

	}
}

// TestMinTransfer 测试最小换乘
func TestMinTransfer(t *testing.T) {
	var source = []struct {
		from  string //开始
		to    string //结束
		times int    // 换乘次数，线路数-1
	}{
		{"公益西桥", "牡丹园", 1},
		{"西红门", "天安门西", 1},
		{"枣园", "高米店北", 0},
		{"湾子", "角门东", 2},
	}

	for _, item := range source {
		from, to, times := item.from, item.to, item.times

		actual, err := graph.MinTransferTimes(from, to)
		if err != nil {
			t.Fatal(err)
			return
		}

		if actual != times {
			t.Errorf("(%s -> %s) min transfer time is %d, expected: %d", from, to, actual, times)
		}

	}
}

// TestMinTransfer 测试最短路径
func TestShortestPath(t *testing.T) {
	var source = []struct {
		from   string // 开始
		to     string // 结束
		length int    // 路径长度 = 节点个数
	}{
		{"公益西桥", "新宫", 2},
		{"公益西桥", "角门东", 3},
		{"公益西桥", "刘家窑", 7},
	}
	for _, item := range source {
		from, to, length := item.from, item.to, item.length

		paths, err := graph.ShortestPath(from, to)
		if err != nil {
			t.Fatal(err)
			return
		}

		if len(paths) != length {
			t.Logf("path info is %s", printPath(paths))
			t.Errorf("(%s -> %s) min path length is %d, expected: %d", from, to, len(paths), length)
		}

	}
}
