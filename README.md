# Go Race Condition with Closed Channel

This repository demonstrates a common race condition in Go involving a closed channel.  The `bug.go` file contains the buggy code, while `bugSolution.go` provides a corrected version.  The race condition arises because multiple goroutines attempt to receive from a channel after it's been closed, leading to unpredictable behavior.