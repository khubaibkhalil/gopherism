package hamming
import "errors"
func Distance(a, b string) (int, error) {
    count:=0
	if len(a) == len(b){
        for i:=range a {
            if a[i]!=b[i]{
                count++
            }
        }
        return count,nil
        
    }else{
        return 0 ,errors.New( "Cant calculate hamming distance ")
    }
}
