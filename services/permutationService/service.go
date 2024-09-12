package permutationService

type IServicepermutationService interface {
	PermutationService(req string) []string
}

type permutationService struct {
	conf string // โดยปกติจะกำหนดค่าเป็น interface ของส่วนอื่นๆ เพื่อที่จะเรียกใช้ func อื่นๆต่อไป
}

func NewpermutationServices(conf string) IServicepermutationService {
	return &permutationService{conf: conf}
}
