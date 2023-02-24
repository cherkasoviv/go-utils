package utils

import (
	"errors"
)

const STARTING_BLOCK_NUMBER = 16000

type IntDeque struct {
	head   int
	tail   int
	count  int
	blocks map[int][100]int
}

func InitIntDeque() *IntDeque {
	newBlocks := map[int][100]int{}
	newBlocks[STARTING_BLOCK_NUMBER] = [100]int{}
	return &IntDeque{
		head:   STARTING_BLOCK_NUMBER * 100,
		tail:   STARTING_BLOCK_NUMBER * 100,
		count:  0,
		blocks: newBlocks,
	}
}

func (deque *IntDeque) PushFront(pushValue int) {
	deque.head--
	deque.count++
	headBlockNumber := deque.head / 100
	headInBlockIndex := deque.head % 100
	if block, inBlocks := deque.blocks[headBlockNumber]; inBlocks {
		block[headInBlockIndex] = pushValue
		deque.blocks[headBlockNumber] = block
	} else {
		deque.blocks[headBlockNumber] = [100]int{}
		block, _ := deque.blocks[headBlockNumber]
		block[headInBlockIndex] = pushValue
		deque.blocks[headBlockNumber] = block

	}
}

func (deque *IntDeque) PushBack(pushValue int) {

	deque.count++
	tailBlockNumber := deque.tail / 100
	tailInBlockIndex := deque.tail % 100
	if block, inBlocks := deque.blocks[tailBlockNumber]; inBlocks {
		block[tailInBlockIndex] = pushValue
		deque.blocks[tailBlockNumber] = block
	} else {
		deque.blocks[tailBlockNumber] = [100]int{}
		block, _ := deque.blocks[tailBlockNumber]
		block[tailInBlockIndex] = pushValue
		deque.blocks[tailBlockNumber] = block

	}
	deque.tail++
}

func (deque *IntDeque) PopFront() (int, error) {
	if deque.count == 0 {
		err := errors.New("stack is empty")
		return 0, err
	} else {
		deque.count--
		popedValue := deque.blocks[deque.head/100][deque.head%100]
		if deque.head%100 == 99 {
			delete(deque.blocks, deque.head/100)
		}
		deque.head++
		return popedValue, nil
	}
}

func (deque *IntDeque) PopBack() (int, error) {
	if deque.count == 0 {
		err := errors.New("stack is empty")
		return 0, err
	} else {
		deque.count--
		deque.tail--
		popedValue := deque.blocks[deque.tail/100][deque.tail%100]
		if deque.tail%100 == 99 {
			delete(deque.blocks, deque.tail/100+1)
		}
		block, _ := deque.blocks[deque.tail/100]
		block[deque.tail%100] = 0
		deque.blocks[deque.tail/100] = block
		return popedValue, nil
	}
}

func (deque *IntDeque) Front() (int, error) {
	if deque.count == 0 {
		err := errors.New("stack is empty")
		return 0, err
	} else {
		return deque.blocks[deque.head/100][deque.head%100], nil
	}

}

func (deque *IntDeque) Back() (int, error) {
	if deque.count == 0 {
		err := errors.New("stack is empty")
		return 0, err
	}
	return deque.blocks[(deque.tail-1)/100][(deque.tail-1)%100], nil
}

func (deque *IntDeque) Size() int {
	return deque.count
}

func (deque *IntDeque) Clear() {
	newBlocks := map[int][100]int{}
	newBlocks[STARTING_BLOCK_NUMBER] = [100]int{}
	*deque = IntDeque{
		head:   STARTING_BLOCK_NUMBER * 100,
		tail:   STARTING_BLOCK_NUMBER * 100,
		count:  0,
		blocks: newBlocks,
	}

}
