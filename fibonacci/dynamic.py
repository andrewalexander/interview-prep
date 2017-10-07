seen = {0:0, 1:1}

def fib(n):
    if n == 0:
        return 0
    if n == 1:
        return 1
    if not seen.get(n):
        # print str(n)+': cache miss...'
        seen[n] = fib(n-1) + fib(n-2)

    return seen.get(n)

def main():
    print 'input 30: ' + str(fib(30))
    print 'input 12: ' + str(fib(12))
    print 'input 27: ' + str(fib(27))

if __name__ == '__main__':
    main()

