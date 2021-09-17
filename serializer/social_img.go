package serializer

import "CarDemo1/model"

// ProductImg 商品图片序列化器
type SocialImg struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"social_id"`
	ImgPath   string `json:"img_path"`
	CreatedAt int64  `json:"created_at"`
}

// BuildImg 序列化帖子图片
func BuildImg(item model.SocialImg) SocialImg {
	return SocialImg{
		ID:        item.ID,
		ProductID: item.SocialID,
		ImgPath:   item.ImgPath,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildImgs 序列化帖子列表
func BuildImgs(items []model.SocialImg) (imgs []SocialImg) {
	for _, item := range items {
		img := BuildImg(item)
		imgs = append(imgs, img)
	}
	return imgs
}
