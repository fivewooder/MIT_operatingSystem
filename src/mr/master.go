package mr

import "log"
import "net"
import "os"
import "net/rpc"
import "net/http"


/* Define the identity of the worker
*/
const(
	mapper = 1
	reducer = 2
)

/* Define the status of the worker machine
*/
const(
	idle = 1
	in_progress = 2
	completed = 3
)

type Master struct {
	// Your definitions here.
	resource []string
}

// Your code here -- RPC handlers for the worker to call.

//
// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
//
func (m *Master) TaskProvider(args *msgWorker2Master, reply *msgMaster2Worker) error {
	if args.status == idle{
		reply.filepos = "test.txt"
		reply.logmsg = "success"
	}
	return nil
}

//
// start a thread that listens for RPCs from worker.go
//
func (m *Master) server() {
	rpc.Register(m)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := masterSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

//
// main/mrmaster.go calls Done() periodically to find out
// if the entire job has finished.
//
func (m *Master) Done() bool {
	ret := false
	// Your code here.


	return ret
}

//
// create a Master.
// main/mrmaster.go calls this function.
// nReduce is the number of reduce tasks to use.
//
func MakeMaster(files []string, nReduce int) *Master {
	m := Master{}
	m.resource = files

	// Your code here.


	m.server()
	return &m
}
