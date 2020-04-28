//very basic implementation based on Jeremy Cerise's BLT tutorial

package main

type GameEntity struct {
	ID     int
	Components map[string]Component
}

func (e *GameEntity) setupComponentsMap() {
	e.Components = make(map[string]Component)
}


func (e *GameEntity) HasComponent(componentName string) bool {
	// Check to see if the entity has the given component
	if _, ok := e.Components[componentName]; ok {
		return true
	} else {
		return false
	}
}

func (e *GameEntity) HasComponents(componentNames []string) bool {
	// Check to see if the entity has the given components
	containsAll := true
	if e != nil {
		for i := 0; i < len(componentNames); i++ {
			if !e.HasComponent(componentNames[i]) {
				containsAll = false
			}
		}
	} else {
		return false
	}
	return containsAll
}

func (e *GameEntity) AddComponent(name string, component Component) {
	// Add a single component to the entity
	e.Components[name] = component
}

func (e *GameEntity) AddComponents(components map[string]Component) {
	// Add several (or one) components to the entity
	for name, component := range components {
		if component != nil {
			//fmt.Printf("Adding component: %s - %v\n", name, component)
			e.Components[name] = component
		}
	}
}

func (e *GameEntity) RemoveComponent(componentName string) {
	// Remove of a component from the entity
	_, ok := e.Components[componentName]

	if ok {
		delete(e.Components, componentName)
	}
}

func (e *GameEntity) RemoveComponents(componentNames []string) {
	for i := 0; i < len(componentNames); i++ {
		_, ok := e.Components[componentNames[i]]

		if ok {
			delete(e.Components, componentNames[i])
		}
	}
}

func (e *GameEntity) GetComponent(componentName string) Component {
	// Return the named component from the entity, if present
	if _, ok := e.Components[componentName]; ok {
		return e.Components[componentName]
	} else {
		return nil
	}
}