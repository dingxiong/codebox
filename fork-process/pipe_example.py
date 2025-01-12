import os
import sys
import select
import time
import fcntl

def main():
    [rf, wf] = os.pipe()
    print(f"fd: {rf} {wf}")
    for fd in [rf, wf]:
        flags = fcntl.fcntl(fd, fcntl.F_GETFL) | os.O_NONBLOCK
        fcntl.fcntl(fd, fcntl.F_SETFL, flags)
        flags = fcntl.fcntl(fd, fcntl.F_GETFD)
        flags |= fcntl.FD_CLOEXEC
        fcntl.fcntl(fd, fcntl.F_SETFD, flags)

    pid = os.fork()
    if pid:
        # parent
        while True:
            read_ready, write_ready, err_ready = select.select([rf], [], [], 1.0)
            if not read_ready:
                continue
            msg = b''
            while True:
                c = os.read(rf, 1)
                msg += c
                if c == b'\n':
                    break
            print(msg)
    else:
        # child
        os.write(wf, b'abc\n')
        os.write(wf, b'1234\n')


if __name__ == "__main__":
    main()
