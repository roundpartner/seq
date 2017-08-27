package claim

import "testing"

func TestNew(t *testing.T) {
    i := Item{}
    c := New()
    c.In <- i
    if <- c.Out != i {
        t.Fail()
    }
}
