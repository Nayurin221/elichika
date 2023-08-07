package model

// CardAwakeningReq ...
type CardAwakeningReq struct {
	CardMasterID     int  `json:"card_master_id"`
	IsAwakeningImage bool `json:"is_awakening_image"`
}

// CardInfo ...
type CardInfo struct {
	CardMasterID               int   `json:"card_master_id"`
	Level                      int   `json:"level"`
	Exp                        int   `json:"exp"`
	LovePoint                  int   `json:"love_point"`
	IsFavorite                 bool  `json:"is_favorite"`
	IsAwakening                bool  `json:"is_awakening"`
	IsAwakeningImage           bool  `json:"is_awakening_image"`
	IsAllTrainingActivated     bool  `json:"is_all_training_activated"`
	TrainingActivatedCellCount int   `json:"training_activated_cell_count"`
	MaxFreePassiveSkill        int   `json:"max_free_passive_skill"`
	Grade                      int   `json:"grade"`
	TrainingLife               int   `json:"training_life"`
	TrainingAttack             int   `json:"training_attack"`
	TrainingDexterity          int   `json:"training_dexterity"`
	ActiveSkillLevel           int   `json:"active_skill_level"`
	PassiveSkillALevel         int   `json:"passive_skill_a_level"`
	PassiveSkillBLevel         int   `json:"passive_skill_b_level"`
	PassiveSkillCLevel         int   `json:"passive_skill_c_level"`
	AdditionalPassiveSkill1ID  int   `json:"additional_passive_skill_1_id"`
	AdditionalPassiveSkill2ID  int   `json:"additional_passive_skill_2_id"`
	AdditionalPassiveSkill3ID  int   `json:"additional_passive_skill_3_id"`
	AdditionalPassiveSkill4ID  int   `json:"additional_passive_skill_4_id"`
	AcquiredAt                 int64 `json:"acquired_at"`
	IsNew                      bool  `json:"is_new"`
}

// CardFavoriteReq ...
type CardFavoriteReq struct {
	CardMasterID int  `json:"card_master_id"`
	IsFavorite   bool `json:"is_favorite"`
}

// UserCardReq ...
type UserCardReq struct {
	UserID       int64 `json:"user_id"`
	CardMasterID int64 `json:"card_master_id"`
}

// NewCardInfo ...
type NewCardInfo struct {
	CardMasterID              int   `json:"card_master_id"`
	Level                     int   `json:"level"`
	Grade                     int   `json:"grade"`
	LoveLevel                 int   `json:"love_level"`
	IsAwakening               bool  `json:"is_awakening"`
	IsAwakeningImage          bool  `json:"is_awakening_image"`
	IsAllTrainingActivated    bool  `json:"is_all_training_activated"`
	ActiveSkillLevel          int   `json:"active_skill_level"`
	PassiveSkillLevels        []int `json:"passive_skill_levels"`
	AdditionalPassiveSkillIds []int `json:"additional_passive_skill_ids"`
	MaxFreePassiveSkill       int   `json:"max_free_passive_skill"`
	TrainingStamina           int   `json:"training_stamina"`
	TrainingAppeal            int   `json:"training_appeal"`
	TrainingTechnique         int   `json:"training_technique"`
	MemberLovePanels          []any `json:"member_love_panels"`
}

type GradeUpCardReq struct {
	CardMasterID int64 `json:"card_master_id"`
	ContentID    int   `json:"content_id"`
}

type CardGradeUpTriggerInfo struct {
	TriggerID            int64 `json:"trigger_id"`
	CardMasterID         int64 `json:"card_master_id"`
	BeforeLoveLevelLimit int   `json:"before_love_level_limit"`
	AfterLoveLevelLimit  int   `json:"after_love_level_limit"`
}
