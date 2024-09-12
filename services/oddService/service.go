package oddService

import repositories "neversitup-test/repositories/odd"

type IServiceOddService interface {
	FindOddService(arr []int) (OddServiceResponse, error)
}

type oddService struct {
	conf repositories.OddRepository // โดยปกติจะกำหนดค่าเป็น interface ของส่วนอื่นๆ เพื่อที่จะเรียกใช้ func อื่นๆต่อไป
}

func NewOddServices(conf repositories.OddRepository) IServiceOddService {
	return &oddService{conf: conf}
}
