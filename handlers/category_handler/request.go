package handlers

import "fmt"

func errParamIsRequired(param string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

func (r *CreateCategoryRequest) Validate() error {
	fields := map[string]interface{}{
		"name": r.Name,
	}

	types := map[string]string{
		"name": "string",
	}

	for field, value := range fields {
		switch v := value.(type) {
		case string:
			if v == "" {
				return errParamIsRequired(field, types[field])
			}
		}
	}

	return nil
}

type UpdateCategoryRequest struct {
	ID   int32  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func (r *UpdateCategoryRequest) Validate() error {
	if r.Name != "" {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
