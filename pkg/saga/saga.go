package saga

import (
	"context"
	"fmt"
	"log"
)

type Saga struct {
	operations    []Operation
	compensations []Compensation
}

func NewSaga() *Saga {
	return &Saga{
		operations:    make([]Operation, 0),
		compensations: make([]Compensation, 0),
	}
}

type Operation struct {
	Func func(context.Context, ...interface{}) error
	Args []interface{}
}

type Compensation struct {
	Func func(context.Context, ...interface{}) error
	Args []interface{}
}

func (s *Saga) AddOperation(op func(context.Context, ...interface{}) error, args ...interface{}) {
	s.operations = append(s.operations, Operation{Func: op, Args: args})
}

func (s *Saga) AddCompensation(comp func(context.Context, ...interface{}) error, args ...interface{}) {
	s.compensations = append(s.compensations, Compensation{Func: comp, Args: args})
}

func (s *Saga) Execute(ctx context.Context) error {
	for i, op := range s.operations {
		if err := op.Func(ctx, op.Args...); err != nil {
			// En caso de error, ejecuta las compensaciones en orden inverso
			for j := len(s.compensations) - 1; j >= 0; j-- {
				if err := s.compensations[j].Func(ctx, s.compensations[j].Args...); err != nil {
					log.Printf("Error al ejecutar compensación %d: %v", j, err)
				}
			}
			return fmt.Errorf("error en operación %d: %v", i, err)
		}
	}
	return nil
}
