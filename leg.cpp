#include <iostream>
#include <vector>

using namespace std;

int main() {
    int n, k;
    cin >> n >> k;
    vector<vector<int>> matrix(n, vector<int>(n, 0));
    matrix[n - 1][k - 1] = 1;
    for(int i = n - 1; i >= 0; i --){
        for(int j = 0; j < n; j ++){
            if (i + 1 < n){
                if (j - 1 >= 0){
                    matrix[i][j] += matrix[i + 1][j - 1];
                }
                if (j + 1 < n){
                    matrix[i][j] += matrix[i + 1][j + 1];
                }
            }
        }
    }
    int cnt = 0;
    for (int i = 0; i < n; i ++){
        cnt += matrix[0][i];
    }
    cout << cnt << endl;
    cout << "Matrix:" << endl;
    for (vector<int> v: matrix){
        for (int el: v){
            cout << el << " ";
        }
        cout << endl;
    }

}