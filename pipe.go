// Copyright 2015 Liu Dong <ddliuhb@gmail.com>.
// Licensed under the MIT license.

package spider

// The pipe interface
type Pipe func(spider *Spider, task *Task)

// Run pipes in parallel
func Parallel(pipes ...Pipe) Pipe {
    return func(s *Spider, t *Task) {
        // chs := make()
    }
}

// Run pipes in series
func Series(pipes ...Pipe) Pipe {
    return func(s *Spider, t *Task) {
        for i := 0; i < len(pipes); i++ {

            // only process working tasks
            if t.Status != WORKING {
                break
            }

            pipes[i](s, t)
        }
    }
}