def rotate_naive(arr, n):
    new = []
    wrap = []
    counter = 0
    # n digits wrap around every time, so lets grab those from the end
    for i in range(len(arr)-n, len(arr)):
        counter = counter + 1
        wrap.append(arr[i])
    for i in range(len(arr)):
        counter = counter + 1
        if i < n:
            new.append(wrap[i])
        else:
            new.append(arr[i-n])
    print 'hit counter: '+ str(counter)
    return new

def rotate(arr, n):
    # swap one pair at a time, separated by our wrap offset. When our right
    # bound reaches the end of the array, we just swap adjacent pairs since we
    # know things have been sorted up until there
    offset = len(arr)-n
    counter = 0
    for i in range(len(arr)-1):
        counter = counter + 1
        if i + offset < len(arr):
            tmp = arr[i+offset]
            arr[i+offset] = arr[i]
            arr[i] = tmp
        else:
            tmp = arr[i+1]
            arr[i+1] = arr[i]
            arr[i] = tmp
    print 'hit counter: '+ str(counter)
    return arr

def main():
    inputs = [
        {'arr': [1,2,3],'n':1},
        {'arr': [1,2,3,4,5],'n':3},
        {'arr': [1,2,3,4,5,6,7,8,9,10],'n':5}
    ]
    print 'naive:'
    for i in inputs:
        print str(i['arr']) + ' -> ' + str(rotate_naive(**i))

    print '-------'
    print 'better:'
    for i in inputs:
        print str(i['arr']) + ' -> ' + str(rotate(**i))

if __name__ == '__main__':
    main()
