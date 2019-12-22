package types

type Status string
type SpeakerID string
type CompanyRef string
type DateString string
type Link string

type Tweet struct {
	Speaker Speaker `json:"speaker"`
	Status  Status  `json:"status"`
	Meetup  Meetup  `json:"meetup"`
}

type Meetup struct {
	meetupInternal
}

type meetupInternal struct {
	Name string     `json:"name"`
	URL  string     `json:"url"`
	Date DateString `json:"date"`
}

type Speaker struct {
	speakerInternal
}

type speakerInternal struct {
	ID      SpeakerID  `json:"id"`
	Name    string     `json:"name"`
	Title   string     `json:"title,omitempty"`
	Email   string     `json:"email"`
	Company CompanyRef `json:"company"`
	Github  string     `json:"github"`
	Twitter string     `json:"twitter,omitempty"`
}
