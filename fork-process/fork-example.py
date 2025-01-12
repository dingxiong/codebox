import os
import sys
import signal
import errno
import time

class ForkExample1():
    """
    Test SIGCHLD handler can successfully capture child exit event,
    and how to extract the exit status of the child process
    """
    CHILDREN = {}

    def __init__(self) -> None:
        self.init_signals()

    def spawn_child(self, child_name: str):
        pid = os.fork()
        if pid != 0:
            # in parent
            self.CHILDREN[pid] = child_name
            print(self.CHILDREN)
            return

        # in child
        pid = os.getpid()
        time.sleep(5)
        sys.exit(0)

    def init_signals(self):
        signal.signal(signal.SIGCHLD, self.handle_sigchld)

    def handle_sigchld(self, sig, frame):
        try:
            while True:
                wpid, status = os.waitpid(-1, os.WNOHANG)
                if not wpid:
                    break
                print(f"child process id {wpid}. exit status: {status}")
                if os.WIFEXITED(status):
                    exit_code = os.WEXITSTATUS(status)
                    print(f"exit code is {exit_code}")
                self.CHILDREN.pop(wpid, None)
        except OSError as e:
            if e.errno != errno.ECHILD:
                raise

    def run(self):
        time.sleep(60)

if __name__ == "__main__":
    c = ForkExample1()
    for i in range(3):
        c.spawn_child(f"consumer-{i}")
    c.run()
