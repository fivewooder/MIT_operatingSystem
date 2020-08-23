package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import "os"
import "strconv"


//
// example to show how to declare the arguments
// and reply for an RPC.
//

type msgWorker2Master struct {
	status int
	identity int
}

type msgMaster2Worker struct {
	filepos string
	logmsg string
}

// Add your RPC definitions here.


// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the master.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func masterSock() string {
	s := "/var/tmp/824-mr-"
	/*
		Return the current process’s user id.
	*/
	s += strconv.Itoa(os.Getuid())
	return s
}
