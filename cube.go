package main

import (
    "fmt"    
)

const PROP_TOTAL_NUM = 4
const (
    PROP_BATTERY = iota // 炮台
    PROP_ROTATE // 旋转
    PROP_MISSILE // 导弹
    PROP_BLOOD // 回血
)

// 魔方整体数据结构
type Cube struct {
    face [][]Square
    roles []*Role
}

// 单个格子数据
type Square struct {
    color byte // just for test

    propId int // 道具Id
}

// 道具接口
type Proper interface {
    Use(*Cube) (string, error)
}


var ColorQueue []byte = []byte{'Y','W','B','G','O','R'}

func NewCube() *Cube {
    cube := new(Cube)
    for i:=0; i<6; i++ {
        oneFace := make([]Square, 8)
        for j:=0; j<8; j++ {
            oneFace[j].color = ColorQueue[i]
            oneFace[j].propId = Rand.Number(1, PROP_TOTAL_NUM)
        }
        
        cube.face = append(cube.face, oneFace)
        cube.roles = append(cube.roles, NewRole(i))
    }
    return cube
}

func (cube *Cube) Print() {
    fmt.Printf("|0   1   2   3   4   5   6   7|\n|-   -   -   -   -   -   -   -|\n")
    for i:=0; i<6; i++ {
        for _, square := range cube.face[i] {
            fmt.Printf("[%+v] ", square.propId)
        }
        fmt.Printf("batteryOrien:%+v, batteryHP:%+v, roleIn:%+v, prop:(%+v,%+v)\n", RotateRules[i][cube.roles[i].batteryOrien].faceIdx, cube.roles[i].batteryHP, cube.roles[i].squareIdx, cube.roles[i].propRotation.num, cube.roles[i].propMissile.num)
    }
    fmt.Println("---------------------------------------------")
}

func (cube *Cube) HttpPrint() string {
    var ret string = "|0   1   2   3   4   5   6   7|\n|-   -   -   -   -   -   -   -|\n"
    for i:=0; i<6; i++ {
        for _, square := range cube.face[i] {
            ret += fmt.Sprintf("[%+v] ", square.propId)
        }
        ret += fmt.Sprintf("batteryOrien:%+v, batteryHP:%+v, roleIn:%+v, prop:(%+v,%+v)\n", RotateRules[i][cube.roles[i].batteryOrien].faceIdx, cube.roles[i].batteryHP, cube.roles[i].squareIdx, cube.roles[i].propRotation.num, cube.roles[i].propMissile.num)
    }
    ret += fmt.Sprintf("---------------------------------------------")
    return ret
}
