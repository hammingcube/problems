# include <iostream>
# include <map>

using namespace std;

bool can_string_be_made_palindrome(const string & s) {
	map<char, int> m;
	for(int i = 0; i < s.size(); i++) {
		m[s[i]] += 1;
	}
	int numOdd = 0;
	for (auto& kv : m) {
		if (kv.second % 2 == 1) {
			numOdd += 1;
			if (numOdd > 1) {
				return false;
			}
		}
	}
	return numOdd <= 1;
}

int main() {
  string s;
  while(cin >> s) {
    cout << can_string_be_made_palindrome(s) << endl;
  }
}