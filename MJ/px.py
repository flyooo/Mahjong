#!/usr/bin/env python

jiang =0

def showi(y):
    j,k=1,0
    while y:    
        if y&1:
            print k,
        y=y>>1
        if j%4:
            j+=1
        else:
            j=1
            k+=1
    print 'end'


def sorti(lst):
    y=0
    for i in lst:
        x=y<<1
        x=1<<i*4|x
        y|=15<<i*4&x
    return y


def combv():
    j=1
    for k in range(4):
        i =0
        i=(i<<4)+j
        i=(i<<4)+j
        i=(i<<4)+j
        print i,bin(i)
        j=(j<<1)+1 

def hu7xd():
    pass

def hu(dat):
    x=sorti(dat)
    n=len(dat)
    arr=[]
    b,w,aaa=0,0,0
    sig = True
    tag = ''

    if n%3==1:
        print n,'card error'
        return 0

    while x:
        b+=1
        i=x&15
        if i==1:
            if (x&273) ==273:   # 'abc'
                x>>=4
                j=(x&15)>>1
                k=(x>>5)&7
                x=(((x>>8<<4)+k)<<4)+j
                arr+=[b,b+1,b+2]
            else:
                if n%3==2 and aaa :    #rollback if n%3==2 and aaa and (x&17)==17 and lst[-1]==(b-1)
                    sig,tag=False,'rollback'
                    print tag
                    return sig,tag,aaa
                else:
                    sig,tag=False,'ABC false'
                    x>>=4
                    print tag

        elif i== 3:
            if n%3==0:
                if n >=6 and (x&819)==819:     #'aabbcc'             
                    x>>=4
                    j=(x&15)>>2
                    k=(x>>6)&3
                    x=(((x>>8<<4)+k)<<4)+j
                    arr+=[b,b+1,b+2,b,b+1,b+2]
                else:
                    x>>=4
                    sig,tag=False,'aa false'
                    print tag
            else:
                if n >=6 and (x&819)==819:     #'aabbcc' of 'aabbccdd'              
                    x>>=4
                    j=(x&15)>>2
                    k=(x>>6)&3
                    x=(((x>>8<<4)+k)<<4)+j
                    arr+=[b,b+1,b+2,b,b+1,b+2]
                else:
                    x>>=4
                    w+=1
                    continue

        elif i==7:
            if n%3==0:
                x>>=4 
                arr+=[b,b,b]
            else:
                if (x&4375)==279:    #'aaabc_' superior to 'aaabcd'   n>=5
                    x>>=4
                    j=(x&15)>>1
                    k=(x>>5)&7
                    x=(((x>>8<<4)+k)<<4)+j
                    arr+=[b,b+1,b+2]
                elif (x&4375)==4375:
                    x>>=4
                    aaa=b
                    arr+=[b,b,b]          
                else:
                    x>>=4 
                    arr+=[b,b,b]

        elif i==15:
            if n%3==2 and n >=8 and (x&831)==831:  #'aaaabbcc'
                x>>=4
                j=(x&15)>>2
                k=(x>>6)&3
                x=(((x>>8<<4)+k)<<4)+j
                w+=1
                arr+=[b,b+1,b+2,b,b+1,b+2]
                continue
            #n%3==0
            if (x&4095)==4095:   #'aaaabbbbcccc'
                x>>=12
                arr+=[b,b+1,b+2,b,b+1,b+2,b,b+1,b+2,b,b+1,b+2]
            elif (x&287)==287:       #'aaaabc  100010001111'
                x>>=4
                j=(x&15)>>1
                k=(x>>5)&7
                x=(((x>>8<<4)+k)<<4)+j
                arr+=[b,b,b,b,b+1,b+2]
            else:
                print 'aaaa false'
                x>>=4 

        else:
            x>>=4
            print '-'

    print w, lst
    return sig,tag,aaa


if __name__ == '__main__':
    # # lst=[random.randint(1,9) for a in range(13)]
    # lst=[3, 5, 6, 4,5,9, 6,7, 1, 8, 4, 2]     #3,6,9,12
    # print lst
    # i=sorti(lst)
    # showi(i)
    # i=i>>4
    # hu(i,len(lst))

    # print '-'*30

    # lst=[1,1,1,2,2,2,5,5,5,5,6,7]     #3,6,9,12
    # print lst
    # i=sorti(lst)
    # showi(i)
    # i=i>>4
    # hu(i,len(lst))


    # print '-'*30
    # lst=[3, 5, 6, 5,5,9, 7,7, 1, 8, 4, 2, 1, 1]     #2,5,8,11,14
    # print lst
    # i=sorti(lst)
    # showi(i)
    # i=i>>4
    # hu(i,len(lst))

    print '-'*30
    lst=[3, 5, 6, 6,5,9, 7,7, 4, 8, 4, 2, 2, 2]     #2,5,8,11,14
    print lst
    i=sorti(lst)
    showi(i)
    i=i>>4
    sig,tag,aaa = hu(lst)
    print sig,tag,aaa
    if tag=='rollback':
        ix=lst.index(aaa)
        del lst[ix]
        ix=lst.index(aaa)
        del lst[ix]
        print lst
        i=sorti(lst)
        i=i>>4
        hu(lst)
