package main

import (
    "os"
)

type Actions struct {
    face int
    orien bool
}

var ActionsMap map[byte]Actions = map[byte]Actions{
    'U' : Actions{0, true},
    'u' : Actions{0, false},

    'D' : Actions{1, true},
    'd' : Actions{1, false},

    'F' : Actions{2, true},
    'f' : Actions{2, false},

    'B' : Actions{3, true},
    'b' : Actions{3, false},

    'L' : Actions{4, true},
    'l' : Actions{4, false},

    'R' : Actions{5, true},
    'r' : Actions{5, false},
}

func main() {
    cube := NewCube()
    cube.Print()
    if len(os.Args) >= 2 {
        rotation := NewRotation()
        rotation.num = len(os.Args[1])

        missile := NewMissile()
        missile.num = 10
        for _, op := range []byte(os.Args[1]) {
            if act, ok := ActionsMap[op]; ok {
                rotation.handleFace = act.face
                rotation.clockWise = act.orien
                rotation.Use(cube)

                missile.selfFace = act.face
                missile.Use(cube)
            }
        }
    }
    cube.Print()
}
