
// go test github.com/CobaltSato/mock/unit_test/is_even -parallel 2

package is_even_test

import ("testing")

func isEven(number int) bool {
  return (number % 2 == 0);
}

func TestIsEven(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {in int; want bool}{
		"+odd":  {5, false},
		"+even": {6, true},
		"-odd":  {-5, false},
		"-even": {-6, true},
		"zero":  {0, true},
	}
	for name, tt := range cases {
		tt := tt // ローカル変数
		t.Run(name, func(t *testing.T) {
			t.Parallel() // 並列実行
			if got := isEven(tt.in); tt.want != got {
				t.Errorf( "want IsEven(%d) = %v, got %v", tt.in, tt.want, got)
      }
    })
	}
}

