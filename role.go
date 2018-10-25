package main

import (
    
)

type Role struct {
    faceIdx int // 所在面
    squareIdx int // 所在面的格子

    batteryOrien int // 炮台朝向, 当前炮台朝向[]RotateRule索引
    batteryHP int // 炮台生命值

    propRotation *Rotation // 道具:旋转
    propMissile *Missile // 道具:导弹
    propDice *Dice // 道具:骰子
    propBlood *Blood // 道具:回血
}

const MAX_HP = 5

func NewRole(face int) *Role {
    if RoleFace(face) {
        role := &Role{
            faceIdx : face,
            squareIdx : Rand.Number(8),

            batteryOrien : Rand.Number(4),
            batteryHP : MAX_HP,
        }
        role.propRotation = &Rotation{Role : role}
        role.propMissile = &Missile{Role : role}
        role.propDice = &Dice{Role : role}
        role.propBlood = &Blood{Role : role}
        return role
    }
    return nil
}

// 0,2,4面才有玩家
func RoleFace(face int) (ret bool) {
    if face%2 == 0 {
        ret = true
    }
    return
}

func (role *Role) Move(cube *Cube, point int) {
    newSquareIdx := (role.squareIdx+point)%8

    newSquarePorpId := cube.face[role.faceIdx][newSquareIdx].propId
    switch newSquarePorpId {
    case PROP_ROTATE:
        role.propRotation.num++
    case PROP_MISSILE:
        role.propMissile.num++
    case PROP_BLOOD:
        role.propBlood.num++
    }

    cube.face[role.faceIdx][role.squareIdx].propId = GenProp()
    role.squareIdx = newSquareIdx
}
