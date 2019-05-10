#!/usr/bin/env python

import queue
import threading
import time

# a debug func to print function stack info
import inspect

class WorkItem():
    #print("enter func -> ", inspect.stack()[0])

    def __init__(self, fn, args, kwargs=None):

        self.fn = fn
        self.args = args
        self.kwargs = kwargs
    def run(self):
        print("enter func -> ", inspect.stack()[0][3])
        try:
            result = self.fn(*self.args, **self.kwargs)
        except Exception as exc:
            print(exc)
            self = None
        else:
            return result


def worker(work_queue):
    #print("enter func -> ", inspect.stack()[0])
    try:
        while True:
            work_item = work_queue.get(block=True)
            work_item.run()
            del work_item
            continue
    except Exception as err:
        print(err)


def caller(fn, *args, **kwargs):
    # print("enter func -> ", inspect.stack()[0])
    if fn:
        workitem = WorkItem(fn, *args, **kwargs)
        work_queue = queue.Queue()
        work_queue.put(workitem)
        ## add caller

        t = threading.Thread(name="test thread", target=worker, args=(work_queue,))
        t.deamon = True
        t.start()
        #_worker = worker(work_queue)

def task1(message, *args, **kwargs):
    print("enter func -> ", inspect.stack()[0])

    print("Handling message <-",kwargs['task_id'])
    time.sleep(3)
    print("Task %d ->message %s done " %(kwargs['task_id'], message))
    return message

if __name__ == '__main__':

    messages = [1,2,3,4,5]
    for i in range(3):
        caller(task1, [messages], dict(task_id=i))
