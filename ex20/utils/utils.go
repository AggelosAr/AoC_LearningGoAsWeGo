package utils

import (
	"errors"
	"fmt"
)

type Modules struct {
	Modules map[string]Module
}

func (m Modules) Get(moduleName string) (Module, error) {
	module, exists := m.Modules[moduleName]
	if !exists {
		return Module{}, errors.New("MODULE DOESNT EXISTS : " + moduleName)
	}
	return module, nil
}

func GetNewModulesMap(sources []string, destinations [][]string) Modules {
	newModulesMap := Modules{Modules: map[string]Module{}}

	for idx, source := range sources {

		newModule := getNewModule(source, destinations[idx])
		newModulesMap.Modules[newModule.Name] = newModule
	}
	
	
	// 2nd pass to find all receivers
	for moduleName := range newModulesMap.Modules {

		sourceModule := newModulesMap.Modules[moduleName]
		for _, dest := range sourceModule.Destinations {
			
			destinationModule, err := newModulesMap.Get(dest)
			if err != nil {
				continue
			}
			if destinationModule.GetType() == "Conjunction" {
				destinationModule.UpdateMemory(sourceModule.Name)
				}
			}
		}

	return newModulesMap
}

func (m Modules) PrintModules() {
	for name, module := range m.Modules {
		fmt.Println("MODULE : ", name)
		fmt.Println("DATA : ", module)
		fmt.Println("_____________________________")
	}
}

type ModuleInterface interface {
	// return the new signal 
	// errors if you shouldnt send the new signal to the destinations
	Send(pulse Signal, args ...interface{}) (Signal, error) 
	GetType() string
	UpdateMemory(sourceName string) // THIS IS BAD //  
}

type Module struct {
	Name         string
	Destinations []string
	ModuleInterface
}

func (m Module) String() string {
	return fmt.Sprintf("<%s -> %v>", m.Name, m.Destinations)
}

func getNewModule(name string, destinations []string) Module {

	module := Module{Destinations: destinations}
	moduleType := string(name[0])

	switch moduleType {
		case "&":
			module.Name = name[1:]
			newMemory := map[string]Signal {}
			module.ModuleInterface = &Conjunction{Memory: newMemory}
		case "%":
			module.Name = name[1:]
			module.ModuleInterface = &FlipFlop{}
		default:
			module.Name = name
			module.ModuleInterface = Broadcaster{}
	}
	
	return module
}


// Conjunction modules (prefix &)
type Conjunction struct {
	Memory map[string]Signal
}

//!!!!!!!!!!!!!!!
// We 1st find all Inputs to all Conjunction To init this correclty !!!!!!!!
//!!!!!!!!!!!!!!!
// We must find all RECEIVERS that are Conjunction XXX -> YYY
// To update their memory (if YYY is Conjunction update it)
func (c *Conjunction) UpdateMemory(inputModule string) {
	c.Memory[inputModule] = Signal{Status: Low}
}


func (c *Conjunction) Send(pulse Signal, args ...interface{}) (Signal, error) {
	
	inputName := args[0].(string) // !

	newSignal := Signal{Status: Low}
	// if all are Hight SEND low
	c.Memory[inputName] = pulse
	for _, signal := range c.Memory {
		if signal.Status == Low {
			newSignal.Flip()
			break
		}
	}
	
	return newSignal, nil
}

func (c Conjunction) GetType() string {
	return "Conjunction"
}


// Flip-flop modules (prefix %)
type FlipFlop struct {
	Status bool // on -> true // off -> false
}

func (f *FlipFlop) Send(pulse Signal, args ...interface{}) (Signal, error) {

	err := errors.New("Rejected")
	newSignal := NewSignal("low")

	if pulse.Status == High {
		// Ignore
	} else if pulse.Status == Low {
		// on -> true
		if f.Status == true {
			//newSignal.Status = Low // already set to low
		// off -> false
		} else if f.Status == false {
			f.Status = !f.Status
			return NewSignal("high"), nil
		}
		f.Status = !f.Status
		err = nil
	}
	
	return newSignal, err
}

func (f FlipFlop) GetType() string {
	return "FlipFlop"
}

func (d FlipFlop) UpdateMemory(inputModule string) {
}

type Broadcaster struct {
	
}

func (b Broadcaster) Send(pulse Signal, args ...interface{}) (Signal, error) {
	return pulse, nil
}

func (b Broadcaster) GetType() string {
	return "Broadcaster"
}
func (b Broadcaster) UpdateMemory(inputModule string) {
}

type SignalType string

const (
    Low  SignalType = "low"
    High SignalType = "high"
)

type Signal struct {
    Status SignalType
}

func (s Signal) String() string {
	return fmt.Sprintf("<STATUS - %s>", s.Status)
}
func NewSignal(signalStr string) Signal {
	newSignal := Signal{}
	switch signalStr {
		case "low":
			newSignal.Status = Low
		case "high":
			newSignal.Status = High
	}
	return newSignal
}

func (s *Signal) Flip() {
	if s.Status == Low {
		s.Status = High
	} else {
		s.Status = Low
	}
}


func (s Signal) FlipNewSig() Signal {
	newSignal := NewSignal("low")
	if s.Status == Low {
		newSignal.Status = High
		return newSignal
	}
	return newSignal
}

