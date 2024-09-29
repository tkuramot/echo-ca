//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package session

type Repository interface {
	Get() (*Session, error)
	Delete() error
	Save(session *Session) error
	Verify() error
}
