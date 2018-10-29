package main

import (
    "fmt"    
)

const PROP_TOTAL_NUM = 4
const (
    PROP_RUINED = iota - 1 // 被毁坏
    PROP_NULL // 空白格子
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
    propId int // 道具Id
}

// 道具接口
type Proper interface {
    Use(*Cube) (string, error)
}

func NewCube() *Cube {
    cube := new(Cube)
    for i:=0; i<6; i++ {
        newRole := NewRole(i)
        oneFace := make([]Square, 8)
        for j:=0; j<8; j++ {
            if newRole == nil || newRole.squareIdx == j { // 初始角色所在格子没有道具
                continue
            }
            oneFace[j].propId = GenProp()
        }
        cube.face = append(cube.face, oneFace)
        cube.roles = append(cube.roles, newRole)
    }
    return cube
}

func GenProp() int {
    return Rand.Number(1, PROP_TOTAL_NUM)
}

func (cube *Cube) Print() {
    fmt.Printf("|0   1   2   3   4   5   6   7|\n|-   -   -   -   -   -   -   -|\n")
    for i:=0; i<6; i+=2 {
        for _, square := range cube.face[i] {
            if square.propId == PROP_RUINED {
                fmt.Print("[x] ")
                continue
            }
            fmt.Printf("[%+v] ", square.propId)
        }
        if RoleFace(i) {
            fmt.Printf("batteryOrien:%+v, batteryHP:%+v, roleIn:%+v, prop:(%+v,%+v,%+v)\n", RotateRules[i][cube.roles[i].batteryOrien].faceIdx, cube.roles[i].batteryHP, cube.roles[i].squareIdx, cube.roles[i].propRotation.num, cube.roles[i].propMissile.num, cube.roles[i].propBlood.num)
        } else {
            fmt.Println("")
        }
    }
    fmt.Println("---------------------------------------------")
}

func (cube *Cube) HttpPrint() string {
    var ret string = "|0   1   2   3   4   5   6   7|\n|-   -   -   -   -   -   -   -|\n"
    for i:=0; i<6; i+=2 {
        for _, square := range cube.face[i] {
            if square.propId == PROP_RUINED {
                ret += fmt.Sprint("[x] ")
                continue
            }
            ret += fmt.Sprintf("[%+v] ", square.propId)
        }
        if RoleFace(i) {
            ret += fmt.Sprintf("batteryOrien:%+v, batteryHP:%+v, roleIn:%+v, prop:(%+v,%+v,%+v)\n", RotateRules[i][cube.roles[i].batteryOrien].faceIdx, cube.roles[i].batteryHP, cube.roles[i].squareIdx, cube.roles[i].propRotation.num, cube.roles[i].propMissile.num, cube.roles[i].propBlood.num)
        } else {
            ret += "\n"
        }
    }
    ret += fmt.Sprintf("---------------------------------------------")
    return ret
}
