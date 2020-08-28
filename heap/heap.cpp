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
 
class MinHeap {
	int *harr;
	int capacity;
	int heap_size;
	public:
	MinHeap(int capacity);
	int getMin() { return harr[0]; }
	void removeMin();
	void insert(int value);
};

MinHeap::MinHeap(int cap) {
	heap_size = 0;
	capacity = cap;
	harr = new int[cap];
}

void MinHeap::insert(int v) {
	if(heap_size == capacity) {
		cout << "Overflow\n" << endl;
		return;
	}

	harr[heap_size++] = v;
	int i = heap_size-1;
	while(i > 0 && harr[i] < harr[(i-1)/2]) {
		swap(harr[i], harr[(i-1)/2]);
		i = (i-1)/2;
	}
}

void MinHeap::removeMin() {
	swap(harr[0], harr[heap_size-1]);
	heap_size--;

	int i = 0;
	while(true) {
		int j = i;
		if(2*i+1 < heap_size && harr[2*i+1] < harr[j])
			j = 2*i+1;

		if(2*i+2 < heap_size && harr[2*i+2] < harr[j])
			j = 2*i+2;

		if(i == j) break;
		swap(harr[i], harr[j]);
		i = j;
	}
}

int main() {
	int n;
	cin >> n;
	vector<int> a(n);
	forn(i, n) cin >> a[i];

	MinHeap heap(n);
	forn(i, n) {
		heap.insert(a[i]);
	}
	// heap sort
	forn(i, n) {
		cout << heap.getMin() << " ";
		heap.removeMin();
	}
	cout << endl;
	return 0;
}

