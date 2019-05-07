# Release 0.6.1 (2019-05-07)

 - vendor: update golang.org/x/sys.
 - fifo: fix tests in windows.
 - sync: fix compiation issue on windows with latest golang.org/x/sys.

# Release 0.6.0 (2018-10-29)

 - vendor: gomod support.
 - ipc: remove Destroyer interface.

# Release 0.5.0 (2016-12-29)

 - sync: condvar on darwin now uses semaphore-based events.
 - sync: events optimization with lightweight event.
 - sync: timed mutex for darwin.
 - sync: timed semaphore for freebsd/darwin.
 - minimum go version is 1.4.

# Release 0.4.0 (2016-12-18)

 - sync: RWMutex for all platforms.
 - sync: Semaphore for all platforms.
 - mq: FastMq has become about 30% faster.
 - all: Added examples.
 - sync: Event for all platforms.

# Release 0.3.0 (2016-10-12)

 - sync: Limited condvar support on darwin and windows via spin waiters.
 - mq: FastMq blocking mode for darwin, windows.
 - sync: Default mutex implementation on freebsd uses futex via umtx syscall.

# Release 0.2.0 (2016-06-22)

- Added FastMq implementation. FastMq is a priority queue based on shared memory. See mq package docs for details.
- Huge windows native shm refactor. It has new API now.
- Improved errors reporting: use github.com/pkg/errors to wrap errors.

# Release 0.1.0 (2016-05-24)

- Initial release.