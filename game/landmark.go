package game

type LandmarkList []Landmark

type Landmark struct {
	Name string
	Type string

	Small          Reward
	RemainingSmall float32

	Medium          Reward
	RemainingMedium float32

	Large          Reward
	RemainingLarge float32
}

type Reward struct {
}
