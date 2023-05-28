package api

import "github.com/aman-singh7/go-hex-arch/internal/ports"

type APIAdapter struct {
	db    ports.DbPort
	arith ports.ArithmeticPort
}

func NewAdapter(db ports.DbPort, arith ports.ArithmeticPort) *APIAdapter {
	return &APIAdapter{
		db:    db,
		arith: arith,
	}
}

func (a *APIAdapter) GetAddition(x, y int32) (int32, error) {
	answer, err := a.arith.Addition(x, y)
	if err != nil {
		return 0, err
	}

	err = a.db.AddToHistory(answer, "addition")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (a *APIAdapter) GetSubtraction(x, y int32) (int32, error) {
	answer, err := a.arith.Subtraction(x, y)
	if err != nil {
		return 0, err
	}

	err = a.db.AddToHistory(answer, "subtraction")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (a *APIAdapter) GetMultiplication(x, y int32) (int32, error) {
	answer, err := a.arith.Multiplication(x, y)
	if err != nil {
		return 0, err
	}

	err = a.db.AddToHistory(answer, "multiplication")
	if err != nil {
		return 0, err
	}

	return answer, nil
}
func (a *APIAdapter) GetDivision(x, y int32) (int32, error) {
	answer, err := a.arith.Division(x, y)
	if err != nil {
		return 0, err
	}

	err = a.db.AddToHistory(answer, "division")
	if err != nil {
		return 0, err
	}

	return answer, nil
}
