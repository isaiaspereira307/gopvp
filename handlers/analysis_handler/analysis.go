package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isaiaspereira307/gopvp/configs"
	"golang.org/x/exp/rand"
)

type APIResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

type Message struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type UserData struct {
	UserID         int32
	Category       string
	Subcategory    string
	HowAmI         string
	HowDoIWantToBe string
	WhatWillIDo    string
	WhenWillIDoIt  string
}

// @BasePath /api/v1
// @Summary Show Analysis
// @Description Show a Analysis
// @Tags analysis
// @Accept json
// @Produce json
// @Param id path string true "Show Analysis Request"
// @Success 200 {object} ShowAnalysisResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /analysis/{id} [get]
func GetAnalysisByUserId(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt64, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		sendErr(ctx, err, http.StatusBadRequest)
	}
	idInt32 := int32(idInt64)
	userData, err := getDataByUserId(ctx, idInt32)
	if err != nil {
		sendErr(ctx, err, http.StatusInternalServerError)
	}

	aiConfig := configs.GetAIConfig()

	request := formatString(userData)

	requestBody, err := json.Marshal(map[string]interface{}{
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": request,
			},
		},
		"model": aiConfig.Model,
	})
	if err != nil {
		sendErr(ctx, err, http.StatusInternalServerError)
	}

	req, err := http.NewRequest("POST", aiConfig.Endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		sendErr(ctx, err, http.StatusInternalServerError)
	}

	req.Header.Set("Authorization", "Bearer "+aiConfig.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		sendErr(ctx, err, http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		sendErr(ctx, fmt.Errorf("erro na resposta: %v", resp.Status), http.StatusInternalServerError)
	}

	var apiResponse APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		sendErr(ctx, err, http.StatusInternalServerError)
	}

	content := apiResponse.Choices[0].Message.Content
	if content == "" {
		sendErr(ctx, fmt.Errorf("erro: não foi possível obter uma resposta"), http.StatusInternalServerError)
	}
	sendSuccess(ctx, "Show Analysis", content)
}

func getDataByUserId(ctx *gin.Context, id int32) ([]UserData, error) {
	objectives, err := queries.GetObjectivesByUserId(ctx, id)
	if err != nil {
		return nil, err
	}
	userData := make([]UserData, len(objectives))

	for i, objective := range objectives {
		userData[i].UserID = id
		category, err := queries.GetCategoryById(ctx, objective.CategoryID)
		if err != nil {
			return nil, err
		}
		userData[i].Category = category.Name
		subcategory, err := queries.GetSubcategoryById(ctx, objective.SubcategoryID)
		if err != nil {
			return nil, err
		}
		userData[i].Subcategory = subcategory.Name
		userData[i].HowAmI = objective.HowAmI
		userData[i].HowDoIWantToBe = objective.HowDoIWantToBe
		userData[i].WhatWillIDo = objective.WhatWillIDo
		userData[i].WhenWillIDoIt = objective.WhenWillIDoIt
	}
	return userData, nil
}

func formatUserData(data UserData) string {
	return fmt.Sprintf("Eu sou %s, quero ser %s, farei %s, quando %s", data.HowAmI, data.HowDoIWantToBe, data.WhatWillIDo, data.WhenWillIDoIt)
}

func formatString(userData []UserData) string {
	personality := "Você é uma IA dedicada a ajudar pessoas a se tornarem melhores, orientando-as na busca de seus objetivos. Sua abordagem deve ser simples e direta, fornecendo conselhos objetivos e concisos com no máximo 20 palavras, baseados nos ensinamentos dos santos, na Bíblia Católica e na Doutrina da Igreja Católica. Cada resposta deve consistir em uma única frase ou citação, sem explicações adicionais."
	currentDate := time.Now().Format("02/01/2006")
	randomNumber := rand.Intn(len(userData))
	userDataFormatted := formatUserData(userData[randomNumber])
	system := fmt.Sprintf("Personality: %s\nCurrentDate: %s\nData: %s\n", personality, currentDate, userDataFormatted)
	return system
}
