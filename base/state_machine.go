package base

type State string

type StateMachine struct {
	CurrentState State
	OnChange     func(prev, current State)
}

func NewStateMachine(current State) *StateMachine {
	return &StateMachine{CurrentState: current}
}

func (sm *StateMachine) ChangeState(newState State) {
	if sm.CurrentState != newState {
		prev := sm.CurrentState
		sm.CurrentState = newState
		if sm.OnChange != nil {
			sm.OnChange(prev, newState)
		}
	}
}

func (sm *StateMachine) Is(state State) bool {
	return sm.CurrentState == state
}

func (sm *StateMachine) IsNot(state State) bool {
	return sm.CurrentState != state
}
