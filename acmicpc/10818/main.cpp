#include <iostream>

using namespace std;

int main() {

    int min = 1000000;
    int max = -1000000;
    int count = 0;

    cin >> count;

    int temp = 0;
    for (int i = 0; i < count; i++) {
        cin >> temp;
        if (temp > max) {
            max = temp;
        }
        if (temp < min) {
            min = temp;
        }
    }
    cout << min << " " << max;

    return 0;
}
