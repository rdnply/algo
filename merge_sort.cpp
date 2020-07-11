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
 
vector<int> merge(vector<int> a, vector<int> b) {
	int i = 0, j = 0;
	int n = sz(a), m = sz(b);
	vector<int> res(n+m);
	while(i+j < n+m) {
		if(j==m || (i < n && a[i] < b[j])){
			res[i+j] = a[i];
			i++;
		} else {
			res[i+j] = b[j];
			j++;
		}
	}

	return res;
}

vector<int> merge_sort(vector<int> v) {
	if(sz(v) <= 1)
		return v;

	int n = sz(v);
	int mid = n/2;
	vector<int> l = vector<int>(v.begin(), v.begin() + mid);
	vector<int> r = vector<int>(v.begin() + mid, v.end());
	l =	merge_sort(l);
	r =	merge_sort(r);
	
	return merge(l, r);
}	

int main() {
	int n;
	cin >> n;
	vector<int> a(n);
	forn(i, n) cin >> a[i];
	vector<int> sorted = merge_sort(a);
	for(int i = 0; i < sz(sorted); i++) {
		cout << sorted[i] << " ";
	}
	
	cout << endl;
	return 0;
}

