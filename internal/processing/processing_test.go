package processing

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProcessMail(t *testing.T) {
	mail_path := "./../../test_data_small/1."
	t.Run("reads files correctly", func(t *testing.T) {
		_, _, err := ProcessMail(mail_path)

		if err != nil {
			t.Errorf("error processing email:  %s", err)
		}
	})

	t.Run("reads sender", func(t *testing.T) {
		want := "heather.dunton@enron.com"
		from, _, err := ProcessMail(mail_path)

		if err != nil {
			t.Errorf("error processing email:  %s", err)
		}

		if from != want {
			t.Errorf("want : %s, got : %s", want, from)
		}
	})

	t.Run("reads 1 receiver", func(t *testing.T) {
		want := []string{"k..allen@enron.com"}
		_, to, err := ProcessMail(mail_path)

		if err != nil {
			t.Errorf("error processing email:  %s", err)
		}

		if len(to) < 1 {
			t.Errorf("didn't get any receiver")
		} else if !reflect.DeepEqual(to, want) {
			t.Errorf("want : %s, got : %s", want, to)
		}
	})

	t.Run("reads multiple receivers", func(t *testing.T) {
		mail_path = "./../../test_data_small/35."
		want := []string{"k..allen@enron.com", "tom.alonso@enron.com",
			"kysa.alport@enron.com", "robert.badeer@enron.com",
			"tim.belden@enron.com"}
		_, to, err := ProcessMail(mail_path)

		if err != nil {
			t.Errorf("error processing email:  %s", err)
		}

		if len(to) < 1 {
			t.Errorf("didn't get any receiver")
		} else if !reflect.DeepEqual(to, want) {
			t.Errorf("want : %s, got : %s", want, to)
		}
	})

	t.Run("empty receiver", func(t *testing.T) {
		mail_path = "./../../test_data_small/2."
		_, to, err := ProcessMail(mail_path)

		if err != nil {
			t.Errorf("error processing email:  %s", err)
		}

		if len(to) != 0 {
			t.Errorf("not empty receivers!")
		}
	})

	t.Run("a lot (87) of receivers", func(t *testing.T) {
		mail_path = "./../../test_data_small/39."
		_, to, err := ProcessMail(mail_path)

		if err != nil {
			t.Errorf("error processing email:  %s", err)
		}
		fmt.Println(to)
		if len(to) != 87 {
			t.Errorf("wanted : 87 receiver, got %d receiver", len(to))
		}
	})
}

func TestRun(t *testing.T) {
	t.Run("process 3 emails", func(t *testing.T) {
		Run("./../../test_data_small")
	})
}

func BenchmarkRun(b *testing.B) {
	for b.Loop() {
		Run("/Users/mehdi/code/go/enron_email/maildir/allen-p")
	}
}
