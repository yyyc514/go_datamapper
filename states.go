package main

import (
	"./dm";
	"reflect";
	"fmt";
)

type State struct {
	Id	int;
	Name	string;
	Abbv	string;
	*dm.Model;
}

// declare models globally
var States dm.Model;

func setup() {
	dm.Init("states.db");	
	States = dm.AddModel("State", "states", reflect.Typeof(State{}));
}

func main() {
	setup();
	state := States.Find(12).(State);
	fmt.Printf("FIND BY ID: %s\n", state.Name);
	
	state = States.First().(State);
	fmt.Printf("FIRST: %s\n", state.Name);

	state = States.Last().(State);
	fmt.Printf("LAST: %s\n", state.Name);

	count := States.Count();
	fmt.Printf("COUNT: %d\n", count);

	println("First 5 states");
	states := States.All(dm.Opts{"limit": 5});
	for s := range states.Results.Iter() {
		state = s.(State);
		fmt.Printf("id: %d  name: %s abbv: %s\n", state.Id, state.Name, state.Abbv);
	}
	println();

	println("States");
	states = States.All();
	for s := range states.Results.Iter() {
		state = s.(State);
		fmt.Printf("id: %d  name: %s abbv: %s\n", state.Id, state.Name, state.Abbv);
	}
}
