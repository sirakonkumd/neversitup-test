package permutationService

import "fmt"

func (s *permutationService) PermutationService(req string) []string {
	fmt.Println("=> ", s.conf)
	results := recursivePermutations([]rune(req), make(map[string]bool), 0)

	var perm []string
	for p := range results {
		perm = append(perm, p)
	}
	return perm
}

func Permutations(s string) []string {
	results := recursivePermutations([]rune(s), make(map[string]bool), 0)

	var perm []string
	for p := range results {
		perm = append(perm, p)
	}
	return perm
}

func recursivePermutations(ru []rune, results map[string]bool, index int) map[string]bool {
	if index == len(ru)-1 {
		results[string(ru)] = true
		return results
	}

	for i := index; i < len(ru); i++ {
		ru[index], ru[i] = ru[i], ru[index]
		recursivePermutations(ru, results, index+1)
		ru[index], ru[i] = ru[i], ru[index]
	}
	return results
}
