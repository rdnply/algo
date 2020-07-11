#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef long double ld;
typedef pair<int, int> pii;
typedef pair<ll, ll> pll;
#define IOS ios_base::sync_with_stdio(false); cin.tie(0);
#define pb push_back
#define PI acos(-1)
#define EPS 1e-5
#define sz(v) ((int)(v).size())
#define all(v) v.begin(),v.end()
#define forn(i, n) for(ll i=0;i<(ll)n;i++)
#define fi first
#define se second
 
void inserting_sort(vector<int>& v) {
	for(int i = 1; i < sz(v); i++) {
		int j = i; 
		while(j > 0 && v[j] < v[j-1]) {
			swap(v[j], v[j-1]);
			j--;
		}
	}
}

int main() {
	int n;
	cin >> n;
	vector<int> a(n);
	forn(i, n) cin >> a[i];
	inserting_sort(a);
	for(int i = 0; i< sz(a); i++) {
		cout << a[i] << " ";
	}
	cout << endl;
	return 0;
}

