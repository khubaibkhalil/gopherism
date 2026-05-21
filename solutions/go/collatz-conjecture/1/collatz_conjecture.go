package collatzconjecture
import "errors"

func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
		return 0, errors.New("only positive integers are allowed")
	}
    count:=0
	for n!=1{
        if n%2==0{
            n=n/2
        }else{
            n=n*3
            n=n+1
        }
        count++
    }

    return count,nil
}
