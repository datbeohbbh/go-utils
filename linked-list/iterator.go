package linkedlist

func (li *linkedList[T]) Iterator() func() *T {
	curHead := li.head
	return func() *T {
		if curHead == nil {
			return nil
		} else {
			ret := curHead.value
			curHead = curHead.next
			return &ret
		}
	}
}
