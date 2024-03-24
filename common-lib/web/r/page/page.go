package page

// Query 分页请求的基类
type Query struct {
	// 当前数据是第几页
	// 1. PageNum = 0 -> PageNum = 1
	// 2. PageNum < 0 -> PageNum = -1 （不分页）
	PageNum int `json:"pageNum" form:"pageNum"`
	// 每一页多少条数据
	PageSize int `json:"pageSize" form:"pageSize"`
}

func NewQuery(pageNum, pageSize int) Query {
	return Query{PageNum: pageNum, PageSize: pageSize}
}

// NewNot 分页接口查询所有(不分页)
func NewNot() Query {
	return NewQuery(-1, -1)
}

func (r Query) Offset() int {
	pNum := r.PageNum
	if r.PageNum < 0 {
		return -1
	}
	if r.PageNum == 0 {
		pNum = 1
	}
	return (int(pNum) - 1) * r.PageSize
}

func (r Query) Limit() int {
	if r.PageSize < 0 {
		return -1
	}
	if r.PageSize == 0 {
		return 5
	}
	return r.PageSize
}

// Data 分页数据响应体
// T 列表每一项的数据类型
type Data[T any] struct {
	Total int64 `json:"total"` // 总页数
	List  []T   `json:"list"`  // 列表数据
}

func New[T any](total int64, list ...T) *Data[T] {
	if len(list) == 0 {
		list = make([]T, 0)
	}
	return &Data[T]{Total: total, List: list}
}

// DataMeta 分页数据响应体（携带额外数据）
// T 列表每一项的数据类型
// M 非列表数据的数据类型
type DataMeta[T any, M any] struct {
	Data[T]
	Meta M `json:"meta"`
}
