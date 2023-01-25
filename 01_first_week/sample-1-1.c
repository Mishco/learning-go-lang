#include <stdio.h>

#define max 1300000
int main()
{
    int i = 0, k = 1, j = 0, n = 1; 

    int arrayNumber[100050];
    char arrayPrimeNumber[max]; 
    
    // fill array with 1
    for (i = 1; i < max; i++)
    { 
        arrayPrimeNumber[i] = '1';
    }

    for (i = 2; i < max; i++)
    {
        // check if number is prime number
        if (arrayPrimeNumber[i] == '1')
        {                       
            // put this prime number in array
            arrayNumber[n] = i; 
            n++;
            // if that nummer is not prime nummer, fill 0 
            for (j = 2 * i; j < max; j += i)
                arrayPrimeNumber[j] = '0'; 
        }
    }

    while (scanf("%d", &k) > 0)
    {
        printf("%d\n", arrayNumber[k]);
    }
    return 0;
}
