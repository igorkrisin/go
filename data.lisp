(progn
	(define pairlis(lambda(lst1 lst2)
		(if(null lst1)
			(quote())
			(cons(list(car lst1)(car lst2))
	    (pairlis(cdr lst1)(cdr lst2))))))
	(define len (lambda (y)
    	    (if (null y)
    		0
    		(+ (len(cdr y)) 1))))
	(define evlis(lambda(lst dict)
	    (if(null lst)
		(quote())
		(cons(eval(car lst) dict)
	    (evlis(cdr lst) dict)))))
	(define assoc(lambda(lst key)
	    (if(null lst)
		false
		(if (= key (car(car lst)))
		    (car(cdr(car lst)))
		    (assoc(cdr lst) key)))))
	(define evprogn(lambda(lst) 
	    (if (null lst)
		(quote())
		(if(=(len lst) 1)
		    lst
		    (progn(eval(car  lst) dict)
		    (evprogn(cdr lst) dict))))))
	(define evcond(lambda(lst dict)
	    (if (null lst)
		(quote())
		(if (=(eval(car(cdr lst)) dict) false)
		    (evcond(cdr lst) dict)
		    (eval(car(cdr(car lst))) dict)))))
	(define eval(lambda(lst dict)
	   	(cond((numberp lst) lst)
	   	    ((symbolp lst) (assoc dict lst)) 
	   		((= (car lst) (quote +)) (+ (eval(car(cdr lst)) dict) (eval(car(cdr(cdr lst))) dict)))
	   		((= (car lst) (quote -)) (- (eval(car(cdr lst)) dict) (eval(car(cdr(cdr lst))) dict)))
	   		((= (car lst) (quote *)) (* (eval(car(cdr lst)) dict) (eval(car(cdr(cdr lst))) dict)))
	   		((= (car lst) (quote /)) (/ (eval(car(cdr lst)) dict) (eval(car(cdr(cdr lst))) dict)))
	   		((= (car lst) (quote =)) (= (eval(car(cdr lst)) dict) (eval(car(cdr(cdr lst))) dict)))
	   		((= (car lst) (quote quote)) (car(cdr lst))) 
			((= (car lst) (quote car)) (car(eval(car (cdr lst)) dict)))
			((= (car lst) (quote cdr)) (cdr(eval(car (cdr lst)) dict)))
			((= (car lst) (quote list)) (evlis(cdr lst) dict))
			((= (car lst) (quote cond)) (evcond (cdr lst) dict))
			((= (car (car lst)) (quote lambda)) (eval(car(cdr(cdr(car lst))))(pairlis(car(cdr(car lst)))(evlis(cdr lst) dict))))
			((= (car lst) (quote progn))(evprogn(cdr lst) dict))
			((= (car lst) (quote if)) (if(eval(car(cdr lst)) dict) (eval(car(cdr(cdr  lst))) dict)(eval(car(cdr(cdr(cdr lst)))) dict))))))
			
	(eval(quote(progn (+ 1 2) (+ 3 4))) (quote())))
	
	
	
	
	