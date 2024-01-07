### Task 1 - Matching Behaviour


* What happens if you remove the `go-command` from the `Seek` call in the `main` function?
  * The go-command makes the seek function run concurrently and makes a separate routine for each person. Due to this, there may be different
 senders and receiver each time the program runs, so the printed message varies. Thus, without the go-command, the function becomes sequential and will 
  always complete in the same order. As a result, `Anna sent a message to Bob. 
  Cody sent a message to Dave. No one received Evaâ€™s message.` will always be printed. 


* What happens if you switch the declaration `wg := new(sync.WaitGroup`) to `var wg sync.WaitGroup`
and the parameter `wg *sync.WaitGroup` to `wg sync.WaitGroup`?
  * `wg` will be sent into the function as a copy instead of a pointer. As a result, the wait group inside the seek function will be different 
  to the one in main, so the wait group in main will end up waiting for a seek function that never communicates that it is done. Thus, the programme becomes
  deadlocked and the last message `No one received Evaâ€™s message.` is never printed. 


* What happens if you remove the buffer on the channel match?
  * The buffer makes room for 1 unmatched send. When it is removed, the last person (Eva) will be sent but won't have a receiver in a unbuffered channel,
  so the program becomes deadlocked.


* What happens if you remove the default-case from the case-statement in the `main` function?
  * Default takes care of values that are not handled by any case-statements. For an uneven array, there will always be an unmatched name,
  so the default case is not necessary and the programme will function as intended. However, for an even array, the case-statement will not be fulfilled. 
  So without a default statement, the programme will stay stuck in the select statement and become deadlocked.  

### Task 3 - MapReduce

|Variant       | Runtime (ms) |
| ------------ | ------------:|
| singleworker |          1376 |
| mapreduce    |          803 |

The optimal number of go routines is 15. 
