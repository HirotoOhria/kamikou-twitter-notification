package gcf

import "testing"

func Test_tweet_IsRT(t1 *testing.T) {
	type fields struct {
		ID        string
		Text      string
		AuthorID  string
		CreatedAt string
		URL       string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "RTであること",
			fields: fields{
				Text: `RT @xxx_xxx_: 【交換】
ライブ 3rdシングル 
アーティスト
トレーディングカード

よろしく…`,
			},
			want: true,
		},
		{
			name: "RTでないこと",
			fields: fields{
				Text: `【交換】
ライブ 3rdシングル 
アーティスト
トレーディングカード

よろしく…`,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := tweet{
				ID:        tt.fields.ID,
				Text:      tt.fields.Text,
				AuthorID:  tt.fields.AuthorID,
				CreatedAt: tt.fields.CreatedAt,
				URL:       tt.fields.URL,
			}
			if got := t.IsRT(); got != tt.want {
				t1.Errorf("IsRT() = %v, want %v", got, tt.want)
			}
		})
	}
}
