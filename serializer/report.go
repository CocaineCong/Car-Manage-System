package serializer

import "CarDemo1/model"

type Report struct {
	ID       uint     `json:"id"`
	TypeID   uint     `json:"type_id"`
	TypeName string   `json:"type_name"`
	UserID   uint     `json:"user_id"`
	UserName string   `json:"user_name"`
	Content  string   `json:"content"`
	Picture  string   `json:"picture"`
	Finish   uint     `json:"finish"`
}

//序列化反馈
func BuildReport(item model.Report) Report {
	return Report{
		ID:        item.ID,
		TypeID   :item.TypeID,
		TypeName   : item.TypeName,
		UserID   :item.UserID,
		UserName   : item.UserName,
		Content   :item.Content,
		Picture   :item.Picture,
		Finish   :item.Finish,
	}
}

//序列化反馈列表
func BuildReports(items []model.Report) (reports []Report) {
	for _, item := range items {
		report := BuildReport(item)
		reports = append(reports, report)
	}
	return reports
}
