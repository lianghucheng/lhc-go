#define _CRT_SECURE_NO_WARNINGS
#include <iostream>

using namespace std;
int a[1000000] = {0};
int main()
{
    int n;
    cin >> n;
    int sum = n * (n + 1) / 2;
    cout << "sum" <<sum << endl;

   if (sum % 2 == 1)
        return 0;
   int x = sum /2 ;
   a[0] = 1;
   for (int i = 1; i <= n; ++i)
   {
        for (int j = i * (i + 1) / 2; j >= i; --j)
            a[j] += a[j - i];
        for (int j = 1; j <= i * (i + 1) / 2 && j <= x; ++j)
            cout << a[j] << "   ";
        cout << endl;
   }
   cout << a[x] << endl;

   return -1;
}