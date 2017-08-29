package Problem0381

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
type para struct {
	orders []string
	paras  []int
}

// ans 是答案
type ans struct {
	solutions []interface{}
}

func Test_Problem0382(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]string{"RandomizedCollection", "insert", "insert", "insert", "insert", "insert", "remove", "remove", "remove", "insert", "remove", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom"},
				[]int{0, 1, 1, 2, 2, 2, 1, 1, 2, 1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			ans{
				[]interface{}{nil, true, false, true, false, false, true, true, true, true, true, 1, 2, 2, 1, 2, 2, 1, 1, 2, 2},
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		sol, ord, par := q.ans.solutions, q.para.orders, q.para.paras
		rs := Constructor()
		for i := 1; i < len(ord); i++ {
			switch ord[i] {
			case "insert":
				ast.Equal(sol[i], rs.Insert(par[i]), "Insert %d", par[i])
			case "remove":
				ast.Equal(sol[i], rs.Remove(par[i]), "Remove %d", par[i])
			case "getRandom":
				r := rs.GetRandom()
				ast.Contains([]int{1, 2}, r, "GetRandom %d", r)
			default:
				ast.Fail("无法处理的命令", "%s", ord[i])
			}
		}
	}
}
