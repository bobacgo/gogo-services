package r

// PageReq 分页请求的基类
type PageReq struct {
	CurrPage int `json:"currPage"` // 当前数据是第几页
	PageSize int `json:"pageSize"` // 每一页多少条数据
}
