(progn
	(define pairlis(lambda(lst1 lst2)
		(if(null lst1)
			(quote())
			(cons(list(car lst1)(car lst2))
	(pairlis(cdr lst1)(cdr lst2))))))
	(define evlis(lambda(lst)
	    (if(null lst)
		(quote())
		(cons(eval(car lst))
	(evlis(cdr lst))))))
	(define eval(lambda(lst)
	   	(cond((numberp lst) lst)
	   		((= (car lst) (quote +)) (+ (eval(car(cdr lst))) (eval(car(cdr(cdr lst))))))
	   		((= (car lst) (quote -)) (- (eval(car(cdr lst))) (eval(car(cdr(cdr lst))))))
	   		((= (car lst) (quote *)) (* (eval(car(cdr lst))) (eval(car(cdr(cdr lst))))))
	   		((= (car lst) (quote /)) (/ (eval(car(cdr lst))) (eval(car(cdr(cdr lst))))))
	   		((= (car lst) (quote quote)) (car(cdr lst))) 
			((= (car lst) (quote car)) (car(eval(car (cdr lst)))))
			((= (car lst) (quote cdr)) (cdr(eval(car (cdr lst)))))
			((= (car lst) (quote list)) (evlis(cdr lst)))
			((= (car lst) (quote if)) (if(eval(car(cdr lst))) (car(cdr(cdr lst)))(cdr(cdr(cdr lst))))))))
			
	(eval(quote(if(= 1 2) 3  6))))
	
	
	
	
	