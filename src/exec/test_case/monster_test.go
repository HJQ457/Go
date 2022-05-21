package monster

import "testing"

func TestMonster(t *testing.T) {
	monster := &Monster{
		Name:  "tom",
		Age:   11,
		Skill: "aa",
	}

	//monster.Store()
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store err")
	}
	t.Logf("monster.Store success")
}

func TestRestore(t *testing.T) {
	monster := &Monster{}

	//monster.Store()
	res := monster.Restore()
	if !res {
		t.Fatalf("monster.Restore err")
	}
	t.Logf("monster.Restore success")
}
