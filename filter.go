package tcd

const (
	IDDesc = "id-desc"
	AbsrisDesc = "absris-desc"
	NoPutDesc = "noput-desc"

	IDAsc = "id-asc"
	AbsrisAsc = "absris-asc"
	NoPutAsc = "noput-asc"

	_contain_noput = "noput~contains~"
	_contain_absris = "absris~contains~"
	_and = "~and~"
)

func (q *Q) FilterByNoPut(noput string) {
	if q.Filter != "" {q.Filter = ""}

	noput = "'"+noput+"'"
	q.Filter = _contain_noput + noput

	return
}

func (q *Q) FilterByAbsris(absris string) {
	if q.Filter != "" {q.Filter = ""}

	absris = "'"+absris+"'"
	q.Filter = _contain_absris + absris
	
	return
}

func (q *Q) FilterByNoPutAndAbsris(noput, absris string) {
	if q.Filter != "" {q.Filter = ""}

	noput = "'"+noput+"'"
	absris = "'"+absris+"'"
	q.Filter = _contain_noput + noput + _and + _contain_absris + absris
	
	return
}