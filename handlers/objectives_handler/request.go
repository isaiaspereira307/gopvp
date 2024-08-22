package handlers

import (
	"fmt"
)

func errParamIsRequired(param string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

type CreateObjectiveRequest struct {
	UserID         int32  `json:"user_id" binding:"required"`
	CategoryID     int32  `json:"category_id" binding:"required"`
	SubcategoryID  int32  `json:"subcategory_id" binding:"required"`
	HowAmI         string `json:"how_am_i" binding:"required"`
	HowDoIWantToBe string `json:"how_do_i_want_to_be" binding:"required"`
	WhatWillIDo    string `json:"what_will_i_do" binding:"required"`
	WhenWillIDoIt  string `json:"when_will_i_do_it" binding:"required"`
}

func (r *CreateObjectiveRequest) Validate() error {
	fields := map[string]interface{}{
		"user_id":             r.UserID,
		"category_id":         r.CategoryID,
		"subcategory_id":      r.SubcategoryID,
		"how_am_i":            r.HowAmI,
		"how_do_i_want_to_be": r.HowDoIWantToBe,
		"what_will_i_do":      r.WhatWillIDo,
		"when_will_i_do_it":   r.WhenWillIDoIt,
	}

	types := map[string]string{
		"user_id":             "int32",
		"category_id":         "int32",
		"subcategory_id":      "int32",
		"how_am_i":            "string",
		"how_do_i_want_to_be": "string",
		"what_will_i_do":      "string",
		"when_will_i_do_it":   "string",
	}

	for field, value := range fields {
		switch v := value.(type) {
		case string:
			if v == "" {
				return errParamIsRequired(field, types[field])
			}
		case int32:
			if v == 0 {
				return errParamIsRequired(field, types[field])
			}
		}
	}

	return nil
}

type UpdateObjectiveRequest struct {
	ID             int32  `json:"id" binding:"required"`
	UserID         int32  `json:"user_id" binding:"required"`
	CategoryID     int32  `json:"category_id" binding:"required"`
	SubcategoryID  int32  `json:"subcategory_id" binding:"required"`
	HowAmI         string `json:"how_am_i" binding:"required"`
	HowDoIWantToBe string `json:"how_do_i_want_to_be" binding:"required"`
	WhatWillIDo    string `json:"what_will_i_do" binding:"required"`
	WhenWillIDoIt  string `json:"when_will_i_do_it" binding:"required"`
}

func (r *UpdateObjectiveRequest) Validate() error {
	if r.ID == 0 || r.UserID == 0 || r.CategoryID == 0 || r.SubcategoryID == 0 || r.HowAmI == "" || r.HowDoIWantToBe == "" || r.WhatWillIDo == "" || r.WhenWillIDoIt == "" {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
