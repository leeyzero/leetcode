package MyCircularQueue

import (
    "testing"
)

func TestMyCircularQueue(t *testing.T) {
    circularQueue := Constructor(3); // 设置长度为 3
    if !circularQueue.EnQueue(1) {
        t.Errorf("enQueue(1) fail")
    }
    if !circularQueue.EnQueue(2) {
        t.Errorf("enQueue(2) fail")
    }
    if !circularQueue.EnQueue(3) {
        t.Errorf("enQueue(3) fail")
    }
    if circularQueue.EnQueue(4) {
        t.Errorf("enQueue(4) fail")
    }
    if r := circularQueue.Rear(); r != 3 {
        t.Errorf("Rear fail.got:%d want:%d", r, 3)
    }
    if !circularQueue.IsFull() {
        t.Errorf("test isFull fail")
    }
    if !circularQueue.DeQueue() {
        t.Errorf("deQueue fail")
    }
    if !circularQueue.EnQueue(4) {
        t.Errorf("enQueue(4) fail")
    }
    if r := circularQueue.Rear(); r != 4 {
        t.Errorf("Rear fail.got:%d want:%d", r, 4)
    }
}