package services

import (
	"dev-portfolio-api/models"
	"dev-portfolio-api/pkg/global"
	"errors"
)

// GetSkillGroups 获取所有技能分类（含子项）
func GetSkillGroups() ([]models.SkillGroup, error) {
	groups := make([]models.SkillGroup, 0)
	err := global.DB.Preload("SkillGroupItems").Find(&groups).Error
	return groups, err
}

// CreateSkillGroup 创建技能分类
func CreateSkillGroup(group *models.SkillGroup) error {
	return global.DB.Create(group).Error
}

// UpdateSkillGroup 更新技能分类
func UpdateSkillGroup(id uint, title string) error {
	return global.DB.Model(&models.SkillGroup{}).Where("id = ?", id).Update("skill_group_title", title).Error
}

// DeleteSkillGroup 删除技能分类
func DeleteSkillGroup(id uint) error {
	return global.DB.Delete(&models.SkillGroup{}, id).Error
}

// GetSkillItems 获取指定分类下的技能项
func GetSkillItems(groupId uint) ([]models.SkillItem, error) {
	items := make([]models.SkillItem, 0)
	err := global.DB.Where("skill_group_id = ?", groupId).Find(&items).Error
	return items, err
}

// CreateSkillItem 创建技能项
func CreateSkillItem(item *models.SkillItem) error {
	return global.DB.Create(item).Error
}

// UpdateSkillItem 更新技能项
func UpdateSkillItem(id uint, icon, title string) error {
	updates := map[string]interface{}{
		"skill_icon":  icon,
		"skill_title": title,
	}
	return global.DB.Model(&models.SkillItem{}).Where("id = ?", id).Updates(updates).Error
}

// DeleteSkillItem 删除技能项
func DeleteSkillItem(id uint) error {
	return global.DB.Delete(&models.SkillItem{}, id).Error
}

// UpdateSkillGroupTitle 保留旧接口兼容
func UpdateSkillGroupTitle(id uint, title string) error {
	return UpdateSkillGroup(id, title)
}

// CreateSkillGroupByAdmin 保留旧接口兼容
func CreateSkillGroupByAdmin(group *models.SkillGroup) error {
	return CreateSkillGroup(group)
}

// DeleteSkillGroupAndItems 删除分类及其所有技能项
func DeleteSkillGroupAndItems(id uint) error {
	// 先删除该分类下的所有技能项
	global.DB.Where("skill_group_id = ?", id).Delete(&models.SkillItem{})
	// 再删除分类
	return global.DB.Delete(&models.SkillGroup{}, id).Error
}

// GetSkillGroupByID 根据 ID 获取分类
func GetSkillGroupByID(id uint) (*models.SkillGroup, error) {
	group := &models.SkillGroup{}
	err := global.DB.First(group, id).Error
	if err != nil {
		return nil, err
	}
	return group, nil
}

// GetSkillItemByID 根据 ID 获取技能项
func GetSkillItemByID(id uint) (*models.SkillItem, error) {
	item := &models.SkillItem{}
	err := global.DB.First(item, id).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

// UpdateSkillItemWithGroup 更新技能项（含分类变更）
func UpdateSkillItemWithGroup(id uint, groupId uint64, icon, title string) error {
	updates := map[string]interface{}{
		"skill_group_id": groupId,
		"skill_icon":     icon,
		"skill_title":    title,
	}
	return global.DB.Model(&models.SkillItem{}).Where("id = ?", id).Updates(updates).Error
}

// BatchCreateSkillItems 批量创建技能项
func BatchCreateSkillItems(items []models.SkillItem) error {
	if len(items) == 0 {
		return errors.New("no items to create")
	}
	return global.DB.Create(&items).Error
}

// UpdateSkillItemsOrder 批量更新技能项排序（通过 order_no 字段，如果需要的话）
func UpdateSkillItemsOrder(items []models.SkillItem) error {
	for _, item := range items {
		if err := global.DB.Save(&item).Error; err != nil {
			return err
		}
	}
	return nil
}

// ReorderSkillGroup 重新排序分类
func ReorderSkillGroup(id uint, order int) error {
	// 目前没有 order_no 字段，如果需要可以添加
	return nil
}

// GetSkillItemsByGroupIDs 根据多个分类 ID 获取技能项
func GetSkillItemsByGroupIDs(groupIds []uint) ([]models.SkillItem, error) {
	items := make([]models.SkillItem, 0)
	err := global.DB.Where("skill_group_id IN ?", groupIds).Find(&items).Error
	return items, err
}

// GetSkillGroupCount 获取分类总数
func GetSkillGroupCount() (int64, error) {
	var count int64
	err := global.DB.Model(&models.SkillGroup{}).Count(&count).Error
	return count, err
}

// GetSkillItemCount 获取技能项总数
func GetSkillItemCount(groupId uint) (int64, error) {
	var count int64
	query := global.DB.Model(&models.SkillItem{})
	if groupId > 0 {
		query = query.Where("skill_group_id = ?", groupId)
	}
	err := query.Count(&count).Error
	return count, err
}
