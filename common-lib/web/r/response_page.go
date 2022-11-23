package r

type Page[T any] struct {
	TotalPage int `json:"totalPage"` // 总页数
	List      []T `json:"list"`      // 列表数据
}

// PageResp 分页数据响应体
// T 列表每一项的数据类型
type PageResp[T any] struct {
	PageReq
	Page[T]
}

// PageMetaResp 分页数据响应体 （携带额外数据）
// T 列表每一项的数据类型
// M 非列表数据的数据类型
type PageMetaResp[T any, M any] struct {
	PageReq
	Page[T]
	Meta M `json:"meta"`
}

// NewPage 分页数据组装
//
// currPage 当前数据是第几页
// totalPage 总的页数
// pageSize 每一页多少条数据
func NewPage[T any](list []T, currPage, totalPage, pageSize int) *PageResp[T] {
	return &PageResp[T]{
		PageReq: PageReq{currPage, pageSize},
		Page:    Page[T]{totalPage, list},
	}
}

// NewPageMeta 分页数据组装-携带非列表数据
//
// currPage 当前数据是第几页
// totalPage 总的页数
// pageSize 每一页多少条数据
// meta 非列表数据
func NewPageMeta[T any, M any](list []T, currPage, totalPage, pageSize int, meta M) *PageMetaResp[T, M] {
	return &PageMetaResp[T, M]{
		PageReq: PageReq{currPage, pageSize},
		Page:    Page[T]{totalPage, list},
		Meta:    meta,
	}
}
