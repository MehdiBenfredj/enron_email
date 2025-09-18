package processing

import (
	"os"
	"reflect"
	"testing"
)

func TestProcessMail(t *testing.T) {
	mail_path := "./../../test_data_small/1."
	t.Run("reads files correctly", func(t *testing.T) {
		r := ProcessMail(mail_path)

		if r.err != nil {
			t.Errorf("error processing email:  %s", r.err)
		}
	})

	t.Run("reads sender", func(t *testing.T) {
		want := "heather.dunton@enron.com"
		r := ProcessMail(mail_path)

		if r.err != nil {
			t.Errorf("error processing email:  %s", r.err)
		}

		if r.from != want {
			t.Errorf("want : %s, got : %s", want, r.from)
		}
	})

	t.Run("reads 1 receiver", func(t *testing.T) {
		want := []string{"k..allen@enron.com"}
		r := ProcessMail(mail_path)

		if r.err != nil {
			t.Errorf("error processing email:  %s", r.err)
		}

		if len(r.to) < 1 {
			t.Errorf("didn't get any receiver")
		} else if !reflect.DeepEqual(r.to, want) {
			t.Errorf("want : %s, got : %s", want, r.to)
		}
	})

	t.Run("reads multiple receivers", func(t *testing.T) {
		mail_path = "./../../test_data_small/35."
		want := []string{"k..allen@enron.com", "tom.alonso@enron.com",
			"kysa.alport@enron.com", "robert.badeer@enron.com",
			"tim.belden@enron.com"}
		r := ProcessMail(mail_path)

		if r.err != nil {
			t.Errorf("error processing email:  %s", r.err)
		}

		if len(r.to) < 1 {
			t.Errorf("didn't get any receiver")
		} else if !reflect.DeepEqual(r.to, want) {
			t.Errorf("want : %s, got : %s", want, r.to)
		}
	})

	t.Run("empty receiver", func(t *testing.T) {
		mail_path = "./../../test_data_small/2."
		r := ProcessMail(mail_path)

		if r.err != nil {
			t.Errorf("error processing email:  %s", r.err)
		}

		if len(r.to) != 0 {
			t.Errorf("not empty receivers!")
		}
	})

	t.Run("a lot (87) of receivers", func(t *testing.T) {
		mail_path = "./../../test_data_small/39."
		r := ProcessMail(mail_path)

		if r.err != nil {
			t.Errorf("error processing email:  %s", r.err)
		}
		if len(r.to) != 87 {
			t.Errorf("wanted : 87 receiver, got %d receiver", len(r.to))
		}
	})
}

func TestRun(t *testing.T) {
	t.Run("process 3 emails", func(t *testing.T) {
		want := make(map[string]map[string]int)
		want["heather.dunton@enron.com"] = map[string]int{
			"k..allen@enron.com": 2,
		}
		want["james.bruce@enron.com"] = map[string]int{
			"k..allen@enron.com":      1,
			"tom.alonso@enron.com":    1,
			"kysa.alport@enron.com":   1,
			"robert.badeer@enron.com": 1,
			"tim.belden@enron.com":    1,
		}
		got, err := Run("./../../test_data_small/multiple_emails_test_data/", "./../../test_data_small/multiple_emails_test_data/result.txt")
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want : %v \n\n got %v", want, got)
		}

	})
}

func TestRunParallel(t *testing.T) {
	t.Run("process 3 emails", func(t *testing.T) {
		want := make(map[string]map[string]int)
		want["heather.dunton@enron.com"] = map[string]int{
			"k..allen@enron.com": 2,
		}
		want["james.bruce@enron.com"] = map[string]int{
			"k..allen@enron.com":      1,
			"tom.alonso@enron.com":    1,
			"kysa.alport@enron.com":   1,
			"robert.badeer@enron.com": 1,
			"tim.belden@enron.com":    1,
		}
		got, err := RunParallel("./../../test_data_small/multiple_emails_test_data/", "./../../test_data_small/multiple_emails_test_data")
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want : %v \n\n got %v", want, got)
		}
	})
}
func BenchmarkRun(b *testing.B) {
	for b.Loop() {
		Run("/Users/mehdi/code/go/enron_email/maildir/allen-p", "./result.txt")
	}
	os.Remove("./result.txt")
}

func BenchmarkRunParallel(b *testing.B) {
	for b.Loop() {
		Run("/Users/mehdi/code/go/enron_email/maildir/allen-p", "./result.txt")
	}
	os.Remove("./result.txt")
}
