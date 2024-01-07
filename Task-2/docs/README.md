### Task 1 - Debugging Concurrent Programs
**Bug 1** <br>
The program does not work because the code will result in a deadlock, which is when
go routines wait for each other endlessly. The line `ch <- "Hello world!"`  sends a string to the 
channel, but this action will block the rest of the program because there is nothing receiving this
value from the channel, as there is only one go routine in the program.

To fix this, we must create another go routine that receives the information. 
This will allow the program to continue running and print the statement. 

**Bug 2** <br>
The for loop is in the main thread while the printing occurs in a separate thread. As a result, the main thread may complete and 
close the channel before the other go routine can print out all the numbers in the channel. Therefore, the final number, 11, will not 
be printed. 

The solution is to implement a WaitGroup that waits for the print statements to complete before the main thread terminates.

### Task 2 - Many Senders; Many Receivers
* What happens if you switch the order of the statements
  `wgp.Wait()` and `close(ch)` in the end of the `main` function?
  * The program will fail and give the error `panic: send on closed channel`. This error occurs because the program wants to wait for the go routines to finish. 
  However, routines cannot send information across a closed channel so the routines will never finish.


* What happens if you move the `close(ch)` from the `main` function
  and instead close the channel in the end of the function
  `Produce`?
  * The program will only work for the first go routine that runs Produce because after it finishes, the channel will close. 
  The rest of the routines will try to send information on the closed channel and thus we get the error `panic: send on closed channel`. 


* What happens if you remove the statement `close(ch)` completely?
  * The program works as intended, but the consumers stay active. 


* What happens if you increase the number of consumers from 2 to 4?
  * The number of receiving go routines will become 4 instead of 2 i.e. the number of producers becomes equal to the number of consumers. 
  As a result, the program will run much quicker because more routines are running in parallel. 


* Can you be sure that all strings are printed before the program
  stops?
  * No, because we do not wait for consumers, so the program may terminate before everything is printed. 