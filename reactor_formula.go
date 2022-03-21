package basistheory_go

import (
	"fmt"
)

type PaginationObject struct {
	TotalItems int `json:"total_items"`
	PageNumber int `json:"page_number"`
	PageSize   int `json:"page_size"`
	TotalPages int `json:"total_pages"`
}

type PaginatedReactorFormulas struct {
	Pagination PaginationObject `json:"pagination"`
	Data       []ReactorFormula `json:"data"`
}

type ReactorFormula struct {
	Id                string                           `json:"id,omitempty"`
	Name              string                           `json:"name"`
	Description       string                           `json:"description"`
	Type              string                           `json:"type"`
	Code              string                           `json:"code"`
	Icon              string                           `json:"icon"`
	Configuration     []ReactorFormulaConfiguration    `json:"configuration"`
	RequestParameters []ReactorFormulaRequestParameter `json:"request_parameters"`
	CreatedAt         string                           `json:"created_at"`
	ModifiedAt        string                           `json:"modified_at"`
}

type MutableReactorFormula struct {
	Name              string                           `json:"name"`
	Description       string                           `json:"description"`
	Type              string                           `json:"type"`
	Code              string                           `json:"code"`
	Icon              string                           `json:"icon"`
	Configuration     []ReactorFormulaConfiguration    `json:"configuration"`
	RequestParameters []ReactorFormulaRequestParameter `json:"request_parameters"`
}

type ReactorFormulaQuery struct {
	Page            string
	Size            string
	Name            string
}

type ReactorFormulaConfiguration struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type ReactorFormulaRequestParameter struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Optional    bool   `json:"optional"`
}

func (client *BasisTheoryClient) CreateReactorFormula(mutatedReactorFormula MutableReactorFormula) (*ReactorFormula, error) {
	result := &ReactorFormula{}

	err := client.post(result, "/reactor-formulas", mutatedReactorFormula)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *BasisTheoryClient) GetReactorFormula(id string) (*ReactorFormula, error) {
	result := &ReactorFormula{}

	err := client.get(result, fmt.Sprintf("/reactor-formulas/%s", id), nil)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *BasisTheoryClient) GetReactorFormulas() (*PaginatedReactorFormulas, error) {
	return client.GetReactorFormulasWithQuery(ReactorFormulaQuery{})
}

func (client *BasisTheoryClient) GetReactorFormulasWithQuery(reactorFormulaQuery ReactorFormulaQuery) (*PaginatedReactorFormulas, error) {
	result := &PaginatedReactorFormulas{}

	params := map[string]string{
		"name":              reactorFormulaQuery.Name,
		"page":              reactorFormulaQuery.Page,
		"size":              reactorFormulaQuery.Size,
	}

	for key, value := range params {
		if value == "" {
			delete(params, key)
		}
	}

	err := client.get(result, "/reactor-formulas", params)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *BasisTheoryClient) DeleteReactorFormula(id string) error {
	return client.delete(fmt.Sprintf("/reactor-formulas/%s", id))
}

func (client *BasisTheoryClient) UpdateReactorFormula(id string, mutatedReactorFormula MutableReactorFormula) (*ReactorFormula, error) {
	result := &ReactorFormula{}

	err := client.put(result, fmt.Sprintf("/reactor-formulas/%s", id), mutatedReactorFormula)

	if err != nil {
		return nil, err
	}

	return result, nil
}
