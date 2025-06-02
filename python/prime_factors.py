# The challenge is to break down a given number into its prime
#  factors. For example, if the input number is 24, we need to find all 
# the prime factors that multiply together to equal 24. In this case,
#  24 can be factored as 2 × 2 × 2 × 3, because 2 and 3 are prime numbers, 
# and their product gives 24.


def prime_factors(n):
    factors = [] 
    while n>1:
        if n%2==0:
            factors.append(2)
            n=n//2
        elif n%3==0:
            factors.append(3)
            n=n//3
        else:
            factors.append(n)
            break
    return factors
    
factors = prime_factors(15)
print(factors)