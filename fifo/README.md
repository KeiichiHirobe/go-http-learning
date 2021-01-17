# 5 FIFO implementations by Go

* SList
  * Singly linked list. Each element has value and next link. List has both of head and tail link.
  
* PList
  * Persistent singly linked list. Just clone all elements when PushBack and PopFront. 
  
* TwoList
  * [Two list queue by Okasaki](https://www.amazon.co.jp/-/en/Chris-Okasaki/dp/0521631246)
  * Persistent queue.
  * Reverse the list when swapping stages.
  
* NaiveListSlice
  * Most simple implementation by slice
  * Persistent queue.
  
* TwoListSlice
  * [Two list queue by Russ Cox](https://github.com/golang/go/blob/682a1d2176b02337460aeede0ff9e49429525195/src/net/http/transport.go#L1242-L1306)
  * Two list queue without the overhead of reversing the list when swapping stages.
  * NOT a persistent queue as my test shows.
