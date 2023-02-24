package utils

import (
	"errors"
	"sync"
)

type IntegerStack struct {
	elements []int
	size     int
	lock     sync.RWMutex
}

func NewIntStack() *IntegerStack {
	return &IntegerStack{
		elements: []int{},
		size:     0,
	}
}

func (inetegerStack *IntegerStack) Push(n int) {
	inetegerStack.lock.Lock()
	defer inetegerStack.lock.Unlock()
	inetegerStack.elements = append(inetegerStack.elements, n)
	inetegerStack.size++
}

func (inetegerStack *IntegerStack) Pop() (int, error) {
	if inetegerStack.size == 0 {
		return 0, errors.New("stack is empty")
	} else {
		inetegerStack.lock.Lock()
		defer inetegerStack.lock.Unlock()
		popedInteger := inetegerStack.elements[inetegerStack.size-1]
		inetegerStack.elements = inetegerStack.elements[:inetegerStack.size-1]
		inetegerStack.size--
		return popedInteger, nil
	}
}

func (inetegerStack *IntegerStack) Back() (int, error) {
	if inetegerStack.size == 0 {
		return 0, errors.New("stack is empty")
	} else {
		inetegerStack.lock.Lock()
		defer inetegerStack.lock.Unlock()
		popedInteger := inetegerStack.elements[inetegerStack.size-1]
		return popedInteger, nil
	}
}

func (inetegerStack *IntegerStack) Size() int {
	return inetegerStack.size
}

func (inetegerStack *IntegerStack) Clear() {
	inetegerStack.lock.Lock()
	defer inetegerStack.lock.Unlock()
	inetegerStack.elements = []int{}
	inetegerStack.size = 0
}

type SymbolStack struct {
	elements []string
	size     int64
	lock     sync.RWMutex
}

func NewSymbolStack() *SymbolStack {
	return &SymbolStack{
		elements: []string{},
		size:     0,
	}
}

func (symbolStack *SymbolStack) Push(n string) {
	symbolStack.lock.Lock()
	defer symbolStack.lock.Unlock()
	symbolStack.elements = append(symbolStack.elements, n)
	symbolStack.size++
}

func (symbolStack *SymbolStack) Pop() (string, error) {
	if symbolStack.size == 0 {
		return "", errors.New("stack is empty")
	} else {
		symbolStack.lock.Lock()
		defer symbolStack.lock.Unlock()
		popedstring := symbolStack.elements[symbolStack.size-1]
		symbolStack.elements = symbolStack.elements[:symbolStack.size-1]
		symbolStack.size--
		return popedstring, nil
	}
}

func (symbolStack *SymbolStack) Back() (string, error) {
	if symbolStack.size == 0 {
		return "", errors.New("stack is empty")
	} else {
		symbolStack.lock.Lock()
		defer symbolStack.lock.Unlock()
		backstring := symbolStack.elements[symbolStack.size-1]
		return backstring, nil
	}
}
func (symbolStack *SymbolStack) PreBack() (string, error) {
	if symbolStack.size <= 1 {
		return "", errors.New("stack is to short")
	} else {
		symbolStack.lock.Lock()
		defer symbolStack.lock.Unlock()
		backstring := symbolStack.elements[symbolStack.size-2]
		return backstring, nil
	}
}
func (symbolStack *SymbolStack) Size() int64 {
	return symbolStack.size
}

func (symbolStack *SymbolStack) Clear() {
	symbolStack.lock.Lock()
	defer symbolStack.lock.Unlock()
	symbolStack.elements = []string{}
	symbolStack.size = 0
}
