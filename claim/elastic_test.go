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

func TestQuery(t *testing.T) {
    c := New()
    c.In <- Item{Id:1, Body: "First"}
    c.In <- Item{Id:3, Body: "Second"}
    c.In <- Item{Id:3, Body: "Third"}
    ch := make(chan Item)
    qry := Query{Id: 3, Out: ch}
    c.Query <- qry
    ei := Item{}
    if ei == <- qry.Out {
        t.Fail()
    }
}

func TestQueryCorrectItem(t *testing.T) {
    c := New()
    c.In <- Item{Id:1, Body: "First"}
    c.In <- Item{Id:2, Body: "Second"}
    c.In <- Item{Id:3, Body: "Third"}
    ch := make(chan Item)
    qry := Query{Id: 2, Out: ch}
    c.Query <- qry
    ei := <- qry.Out
    if "Second" != ei.Body {
        t.Fail()
    }
}

func TestQueryNotFound(t *testing.T) {
    c := New()
    c.In <- Item{Id:1, Body: "Hello World"}
    c.In <- Item{Id:3, Body: "Goodbye World"}
    ch := make(chan Item)
    qry := Query{Id: 2, Out: ch}
    c.Query <- qry
    ei := Item{}
    if ei != <- qry.Out {
        t.Fail()
    }
}

func TestQueryDelete(t *testing.T) {
    c := New()
    c.In <- Item{Id:1, Body: "Hello World"}
    c.In <- Item{Id:2, Body: "Goodbye World"}
    ch := make(chan Item)
    qry := Query{Id: 1, Out: ch, Delete: true}
    c.Query <- qry
    <- qry.Out
    qry = Query{Id: 1, Out: ch}
    c.Query <- qry
    ei := Item{}
    if ei != <- qry.Out {
        t.Fail()
    }
}