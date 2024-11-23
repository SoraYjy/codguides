package ttkdescriber

import "testing"

func TestAdd(t *testing.T) {
	damage := NewDamage()
	CalCombinationTTK(damage)
}

func NewDamage() Damage {
	damage := Damage{}
	damage.firerate = 800
	damage.health = 300

	damage.headRate = 15
	damage.neckRate = 5
	damage.upperTorsoRate = 30
	damage.lowerTorsoRate = 20
	damage.upperArmRate = 10
	damage.lowerArmRate = 5
	damage.upperLegRate = 50
	damage.lowerLegRate = 50

	damage.head = 38
	damage.neck = 35
	damage.upperTorso = 35
	damage.lowerTorso = 35
	damage.upperArm = 35
	damage.lowerArm = 35
	damage.upperLeg = 30
	damage.lowerLeg = 30

	return damage
}
