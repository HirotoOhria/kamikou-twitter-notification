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

func Test_tweet_IsTradeWithMoney(t1 *testing.T) {
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
			name: "金額交換であること",
			fields: fields{
				Text: `【譲渡】プロセカ プロジェクトセカイ
神山高校文化祭 神高文化祭

譲：11/6(日)入場+13時公演チケット×2

求：定価+手数料

・2枚まとめてCloakでの分配
・ゆうちょ銀行への振込可能な方
・取引初心者でも構わない方`,
			},
			want: true,
		},
		{
			name: "金額交換でないこと",
			fields: fields{
				Text: `プロセカ 交換 特典 トレカ ニーゴ 3rd CD

譲)瑞希

求)同種：まふゆ

◎郵送／神高文化祭(11/6)手渡し でのお取引`,
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
			if got := t.IsTradeWithMoney(); got != tt.want {
				t1.Errorf("IsTradeWithMoney() = %v, want %v", got, tt.want)
			}
		})
	}
}
