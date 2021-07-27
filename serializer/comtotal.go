package serializer

type TotalNum struct {
	Total int `json:"total"`
	PageNum int `json:"page_num"`
}

func BuildTotalPageNum(total ,pageNum int) TotalNum {
	return TotalNum{
		Total:     total,
		PageNum:   pageNum,
	}
}
