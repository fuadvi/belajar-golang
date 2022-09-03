
def factorial(num):
    if num ==0:
        return 1
    
    hasil= num * factorial(num - 1)
    return hasil

print(factorial(0))