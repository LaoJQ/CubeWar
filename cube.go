package main

import (
    "fmt"    
)

const PROP_TOTAL_NUM = 3
const (
    PROP_BATTERY = iota // 炮台
    PROP_ROTATE // 旋转
    PROP_MISSILE // 导弹
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
    batteryOrien int // 炮台朝向, 当前炮台朝向[]RotateRule索引
    batteryHP int // 炮台生命值
}

// 道具接口
type Proper interface {
    Use(*Cube) error
}


var ColorQueue []byte = []byte{'Y','W','B','G','O','R'}

func NewCube() *Cube {
    cube := new(Cube)
    for i:=0; i<6; i++ {
        oneFace := make([]Square, 9)
        for j:=0; j<9; j++ {
            oneFace[j].color = ColorQueue[i]
            if j == 8{
                oneFace[j].propId = PROP_BATTERY
            } else {
                oneFace[j].propId = Rand.Number(1, PROP_TOTAL_NUM)
                oneFace[8].batteryOrien = Rand.Number(4)
                oneFace[8].batteryHP = 5
            }
        }
        
        cube.face = append(cube.face, oneFace)
        cube.roles = append(cube.roles, NewRole(i))
    }
    return cube
}

// func (cube *Cube) Print() {
//     for i:=0; i<6; i++ {
//         squares := make([]byte, 9)
//         for _, square := range cube.face[i] {
//             squares = append(squares, square.color)
//         }
//         fmt.Println(string(squares), RotateRules[i][cube.face[i][8].batteryOrien].faceIdx, cube.face[i][8])
//     }
//     fmt.Println("---------------")
// }

func (cube *Cube) Print() {
    for i:=0; i<6; i++ {
        for _, square := range cube.face[i] {
            fmt.Printf("[%+v] ", square.propId)
        }
        fmt.Printf("batteryOrien:%+v, batteryHP:%+v, roleIn:%+v, prop:(%+v,%+v)\n", RotateRules[i][cube.face[i][8].batteryOrien].faceIdx, cube.face[i][8].batteryHP, cube.roles[i].squareIdx, cube.roles[i].propRotation.num, cube.roles[i].propMissile.num)
    }
    fmt.Println("---------------")
}
