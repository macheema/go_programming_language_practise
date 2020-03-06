package section44

import (
	"ch4/section44/embedstructs"
	"fmt"
)

//Test ...
func Test() {

	wheel := embedstructs.Wheel{Spokes: 1}

	//compilation error
	//unknown field 'circle' in struct literal of type embedstructs.Wheel
	//cannot refer to unexported name embedstructs.circle

	// wheel = embedstructs.Wheel{circle: embedstructs.circle{point: embedstructs.point{X: 0, Y: 0}, Radius: 0}, Spokes: 1}

	fmt.Printf("%#v\n", wheel)
}
