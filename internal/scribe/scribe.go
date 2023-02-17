package scribe

type Scribe struct {
    svcaccJSON []byte
    smtpPass string
}

func NewScribe() *Scribe {
    s := Scribe{}

    return &s
}

